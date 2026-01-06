// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IReceiverTemplate} from "./IReceiverTemplate.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title AdEscrow
 * @notice Simple escrow contract for trustless advertisement protocol
 * @dev Advertisers deposit funds, CRE workflow verifies criteria and withdraws to influencers
 */
contract AdEscrow is IReceiverTemplate {
    using SafeERC20 for IERC20;

    // Immutable address for native ETH
    address public immutable NATIVE_ETH_ADDRESS = address(0);

    // Campaign state enum
    enum CampaignState {
        Active,    // Campaign is active, waiting for criteria to be met
        Withdrawn, // Criteria met and funds successfully withdrawn
        Refunded    // Campaign expired and funds refunded to the advertiser
    }

    struct CreReport {
        CreActions action;
        bytes data; // Data for the action
    }

    enum CreActions {
        Refund, // Refund the funds to the advertiser
        Release // Release the funds to the influencer
    }

    // Campaign structure matching the ABI
    struct Campaign {
        address advertiser; // funds depositor
        address influencer; // influencer who will post the content and receive the funds
        address token; // Token address (address(0) for native ETH)
        uint256 amount; // Amount deposited
        string contentText; // Text of the content to validate
        uint256 minViews; // Minimum views required to fulfill the campaign
        uint64 campaignDuration; // Duration the content is expected to be posted for
        uint256 deadline; // Campaign deadline (timestamp)
        CampaignState state; // Current state of the campaign
    }

    // Mapping from campaign ID to campaign data
    mapping(uint256 => Campaign) public campaigns;

    // Mapping for whitelisted tokens (address(0) is always whitelisted for native ETH)
    mapping(address => bool) public whitelistedTokens;

    // Counter for campaign IDs
    uint256 public campaignCounter;

    // Events matching the ABI
    event CampaignDeposited(
        uint256 indexed campaignId,
        address indexed advertiser,
        address indexed influencer,
        address token,
        uint256 amount
    );

    event FundsWithdrawn(
        uint256 indexed campaignId,
        address indexed influencer,
        address token,
        uint256 amount
    );

    event CampaignExpired(
        uint256 indexed campaignId,
        address indexed advertiser,
        address token,
        uint256 amount
    );

    event DeliveryActionCalled(uint256 indexed campaignId);

    event TokenWhitelisted(address indexed token);
    event TokenRemovedFromWhitelist(address indexed token);

    /**
     * @param _forwarder Address of the Chainlink forwarder
     * @param _whitelistedTokens Array of token addresses to whitelist (address(0) is always whitelisted for native ETH)
     */
    constructor(address _forwarder, address[] memory _whitelistedTokens) {
        forwarderAddress = _forwarder;
        campaignCounter = 1; // Start from 1
        
        // Native ETH is always whitelisted
        whitelistedTokens[NATIVE_ETH_ADDRESS] = true;
        
        // Whitelist provided tokens
        for (uint256 i = 0; i < _whitelistedTokens.length; i++) {
            whitelistedTokens[_whitelistedTokens[i]] = true;
            emit TokenWhitelisted(_whitelistedTokens[i]);
        }
    }

    /**
     * @notice Deposit funds (native ETH or ERC20 tokens) for a new campaign
     * @param token Token address (NATIVE_ETH_ADDRESS for native ETH, or ERC20 token address)
     * @param influencer Address that will receive funds when criteria are met
     * @param amount Amount to deposit (msg.value for native ETH, or token amount for ERC20)
     * @param contentText Expected text/content body (temporary stand-in for URL)
     * @param minViews Minimum number of views required
     * @param expiryDeadline Unix timestamp deadline for the campaign
     * @param campaignDuration Duration the content is expected to be posted form if zero the campaing duration will not be checked. 
     * @return campaignId The ID of the newly created campaign
     */
    function deposit(
        address token,
        address influencer,
        uint256 amount,
        string memory contentText,
        uint256 minViews,
        uint256 expiryDeadline,
        uint64 campaignDuration
    ) external payable returns (uint256 campaignId) {
        require(whitelistedTokens[token], "Token not whitelisted");
        require(influencer != address(0), "Invalid influencer address");
        require(expiryDeadline > block.timestamp, "Deadline must be in future");
        
        if (token == NATIVE_ETH_ADDRESS) {
            // Native ETH deposit
            require(msg.value > 0, "Must deposit funds");
            require(msg.value == amount, "Amount mismatch with msg.value");
        } else {
            // ERC20 token deposit
            require(amount > 0, "Must deposit funds");
            require(msg.value == 0, "Cannot send ETH with ERC20 deposit");
            
            // Transfer tokens from advertiser to this contract
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        }

        campaignId = campaignCounter;
        campaignCounter++;

        campaigns[campaignId] = Campaign({
            advertiser: msg.sender,
            influencer: influencer,
            token: token,
            amount: amount,
            contentText: contentText,
            minViews: minViews,
            campaignDuration: campaignDuration,
            deadline: expiryDeadline,
            state: CampaignState.Active
        });

        emit CampaignDeposited(campaignId, msg.sender, influencer, token, amount);
        
        return campaignId;
    }

    /**
     * @notice Get campaign details
     * @param campaignId The campaign ID to query
     * @return Campaign struct with all campaign details
     */
    function getCampaign(uint256 campaignId) external view returns (Campaign memory) {
        require(campaigns[campaignId].advertiser != address(0), "Campaign does not exist");
        return campaigns[campaignId];
    }

    /**
     * @inheritdoc IReceiverTemplate
     * @dev Decodes the report payload as CreReport with action (Release or Refund) and associated data
     */
    function _processReport(bytes calldata report) internal override {
        CreReport memory creReport = abi.decode(report, (CreReport));
        if (creReport.action == CreActions.Release) {
            (uint256 campaignId, uint256 actualViews, uint64 postedDuration) = abi.decode(creReport.data, (uint256, uint256, uint64));  
            _fulfillCampaign(campaignId, actualViews, postedDuration);
        } else if (creReport.action == CreActions.Refund) {
            (uint256 campaignId) = abi.decode(creReport.data, (uint256));
            markExpired(campaignId);
        } else {
            revert("Invalid CRE action");
        }
    }

    /**
     * @notice Manual trigger for delivery action check
     * @dev Can be called by anyone to trigger a check (emits event for CRE workflow)
     * @param campaignId The campaign ID to check
     */
    function deliveryAction(uint256 campaignId) external {
        require(campaigns[campaignId].advertiser != address(0), "Campaign does not exist");
        emit DeliveryActionCalled(campaignId);
    }


    /**
     * @notice Check if a campaign has expired (deadline passed and still active)
     * @param campaignId The campaign ID to check
     * @return True if campaign is expired
     */
    function isExpired(uint256 campaignId) external view returns (bool) {
        Campaign memory campaign = campaigns[campaignId];
        require(campaign.advertiser != address(0), "Campaign does not exist");
        return campaign.state == CampaignState.Active && block.timestamp > campaign.deadline;
    }

    /**
     * @notice Mark a campaign as refunded and refund advertiser
     * @dev Internal function called by CRE workflow when refund conditions are met
     * @param campaignId The campaign ID to refund
     */
    function markExpired(uint256 campaignId) internal {
        Campaign storage campaign = campaigns[campaignId];
        require(campaign.advertiser != address(0), "Campaign does not exist");
        require(campaign.state == CampaignState.Active, "Campaign not active");
        
        campaign.state = CampaignState.Refunded;
        
        // Refund the advertiser
        _transferFunds(campaign.token, campaign.advertiser, campaign.amount);
        
        emit CampaignExpired(campaignId, campaign.advertiser, campaign.token, campaign.amount);
    }

    function _fulfillCampaign(uint256 campaignId, uint256 actualViews, uint64 postedDuration) internal {
        Campaign storage campaign = campaigns[campaignId];
        require(campaign.advertiser != address(0), "Campaign does not exist");
        require(campaign.state == CampaignState.Active, "Campaign not active");
        require(block.timestamp <= campaign.deadline, "Campaign deadline passed");
        require(actualViews >= campaign.minViews, "Views criteria not met");
        if (campaign.campaignDuration > 0) {
            require(postedDuration >= campaign.campaignDuration, "Campaign duration not met");
        }

        campaign.state = CampaignState.Withdrawn;

        // Transfer funds to influencer
        _transferFunds(campaign.token, campaign.influencer, campaign.amount);

        emit FundsWithdrawn(campaignId, campaign.influencer, campaign.token, campaign.amount);
    }

    /**
     * @notice Internal function to transfer funds (native ETH or ERC20)
     * @param token Token address (NATIVE_ETH_ADDRESS for native ETH)
     * @param to Recipient address
     * @param amount Amount to transfer
     */
    function _transferFunds(address token, address to, uint256 amount) internal {
        if (token == NATIVE_ETH_ADDRESS) {
            // Native ETH transfer
            (bool success,) = to.call{value: amount}("");
            require(success, "ETH transfer failed");
        } else {
            // ERC20 transfer
            IERC20(token).safeTransfer(to, amount);
        }
    }

    /**
     * @notice Add a token to the whitelist (owner only)
     * @param token Token address to whitelist
     */
    function whitelistToken(address token) external onlyOwner {
        require(token != NATIVE_ETH_ADDRESS, "Cannot whitelist native ETH explicitly");
        require(!whitelistedTokens[token], "Token already whitelisted");
        whitelistedTokens[token] = true;
        emit TokenWhitelisted(token);
    }

    /**
     * @notice Remove a token from the whitelist (owner only)
     * @param token Token address to remove from whitelist
     */
    function removeTokenFromWhitelist(address token) external onlyOwner {
        require(token != NATIVE_ETH_ADDRESS, "Cannot remove native ETH from whitelist");
        require(whitelistedTokens[token], "Token not whitelisted");
        whitelistedTokens[token] = false;
        emit TokenRemovedFromWhitelist(token);
    }
}
