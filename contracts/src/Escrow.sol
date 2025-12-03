// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IReceiverTemplate} from "./IReceiverTemplate.sol";

/**
 * @title AdEscrow
 * @notice Simple escrow contract for trustless advertisement protocol
 * @dev Advertisers deposit funds, CRE workflow verifies criteria and withdraws to influencers
 */
contract AdEscrow is IReceiverTemplate {
    // Campaign structure matching the ABI
    struct Campaign {
        address advertiser; // Who deposited the funds
        address influencer; // Who will receive the funds
        uint256 amount; // Amount deposited
        string contentText; // Text of the content to validate
        uint256 minViews; // Minimum views required
        uint256 deadline; // Campaign deadline (timestamp)
        bool fulfilled; // Whether criteria have been met
        bool withdrawn; // Whether funds have been withdrawn
    }

    // Mapping from campaign ID to campaign data
    mapping(uint256 => Campaign) public campaigns;

    // Counter for campaign IDs
    uint256 private campaignCounter;

    // Events matching the ABI
    event CampaignDeposited(
        uint256 indexed campaignId, address indexed advertiser, address indexed influencer, uint256 amount
    );

    event FundsWithdrawn(uint256 indexed campaignId, address indexed influencer, uint256 amount);

    event DeliveryActionCalled(uint256 indexed campaignId);

    constructor(address _forwarder) {
        forwarderAddress = _forwarder;
        campaignCounter = 1; // Start from 1
    }

    /**
     * @notice Deposit funds for a new campaign
     * @param campaignId The campaign ID (should be unique, caller is responsible)
     * @param influencer Address that will receive funds when criteria are met
     * @param contentText Expected text/content body (temporary stand-in for URL)
     * @param minViews Minimum number of views required
     * @param deadline Unix timestamp deadline for the campaign
     */
    function deposit(
        uint256 campaignId,
        address influencer,
        string memory contentText,
        uint256 minViews,
        uint256 deadline
    ) external payable {
        require(msg.value > 0, "Must deposit funds");
        require(influencer != address(0), "Invalid influencer address");
        require(deadline > block.timestamp, "Deadline must be in future");
        require(campaigns[campaignId].advertiser == address(0), "Campaign ID already exists");

        campaigns[campaignId] = Campaign({
            advertiser: msg.sender,
            influencer: influencer,
            amount: msg.value,
            contentText: contentText,
            minViews: minViews,
            deadline: deadline,
            fulfilled: false,
            withdrawn: false
        });

        emit CampaignDeposited(campaignId, msg.sender, influencer, msg.value);
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
     * @dev Decodes the report payload as (campaignId, actualViews)
     */
    function _processReport(bytes calldata report) internal override {
        (uint256 campaignId, uint256 actualViews) = abi.decode(report, (uint256, uint256));
        _fulfillCampaign(campaignId, actualViews);
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
     * @notice Get the next available campaign ID
     * @return The next campaign ID
     */
    function getNextCampaignId() external view returns (uint256) {
        return campaignCounter;
    }

    /**
     * @notice Increment and return next campaign ID
     * @dev Useful for getting a unique ID before depositing
     * @return The new campaign ID
     */
    function createCampaignId() external returns (uint256) {
        uint256 newId = campaignCounter;
        campaignCounter++;
        return newId;
    }

    function _fulfillCampaign(uint256 campaignId, uint256 actualViews) internal {
        Campaign storage campaign = campaigns[campaignId];
        require(campaign.advertiser != address(0), "Campaign does not exist");
        require(!campaign.withdrawn, "Already withdrawn");
        require(block.timestamp <= campaign.deadline, "Campaign deadline passed");
        require(actualViews >= campaign.minViews, "Views criteria not met");

        campaign.fulfilled = true;
        campaign.withdrawn = true;

        (bool success,) = campaign.influencer.call{value: campaign.amount}("");
        require(success, "Transfer failed");

        emit FundsWithdrawn(campaignId, campaign.influencer, campaign.amount);
    }
}
