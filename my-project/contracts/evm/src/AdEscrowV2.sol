// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC165} from "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract AdEscrowV2 is Ownable, ReentrancyGuard, IERC165 {
    using SafeERC20 for IERC20;

    // ──────────────────────────── Constants ────────────────────────────
    address public constant NATIVE_ETH_ADDRESS = 0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE;
    uint256 public constant MAX_FEE_BPS = 1000; // 10% cap

    // ──────────────────────────── State ────────────────────────────────
    enum CampaignState {
        Proposed,   // 0 — created but influencer hasn't accepted
        Accepted,   // 1 — influencer accepted, waiting for funding
        Funded,     // 2 — advertiser deposited funds
        Posted,     // 3 — influencer submitted tweet URL (off-chain)
        Verifying,  // 4 — CRE verification triggered (off-chain)
        Completed,  // 5 — verification passed, funds released
        Refunded,   // 6 — expired or cancelled, funds returned
        Disputed,   // 7 — dispute raised
        Cancelled   // 8 — cancelled before acceptance
    }

    struct Campaign {
        address advertiser;
        address influencer;
        address token;
        uint256 amount;
        bytes32 contentHash;       // keccak256(normalizedText) — saves gas vs string
        uint256 minViews;
        uint64  campaignDuration;  // seconds the tweet must stay up
        uint256 deadline;          // unix timestamp — campaign expiry
        CampaignState state;
        uint256 platformFee;       // fee taken on release (calculated at deposit)
        bool    influencerAccepted;
    }

    uint256 public campaignCounter;
    uint256 public platformFeeBps;  // e.g. 250 = 2.5%
    address public feeRecipient;

    mapping(uint256 => Campaign) public campaigns;
    mapping(address => bool) public whitelistedTokens;

    // Keystone/CRE forwarder config
    address public forwarderAddress;
    address public expectedAuthor;
    bytes32 public expectedWorkflowId;
    bytes10 public expectedWorkflowName;

    // ──────────────────────────── Events ───────────────────────────────
    event CampaignCreated(
        uint256 indexed campaignId,
        address indexed advertiser,
        address indexed influencer,
        address token,
        uint256 amount,
        bytes32 contentHash
    );
    event CampaignAccepted(uint256 indexed campaignId, address indexed influencer);
    event CampaignFunded(
        uint256 indexed campaignId,
        address indexed advertiser,
        address token,
        uint256 amount
    );
    event CampaignCancelled(uint256 indexed campaignId, address indexed advertiser);
    event FundsReleased(
        uint256 indexed campaignId,
        address indexed influencer,
        address token,
        uint256 influencerAmount,
        uint256 feeAmount
    );
    event FundsRefunded(
        uint256 indexed campaignId,
        address indexed advertiser,
        address token,
        uint256 amount
    );
    event DisputeRaised(uint256 indexed campaignId, address indexed raisedBy);
    event DisputeResolved(uint256 indexed campaignId, bool releasedToInfluencer);
    event DeliveryActionCalled(uint256 indexed campaignId);
    event TokenWhitelisted(address indexed token);
    event TokenRemovedFromWhitelist(address indexed token);

    // ──────────────────────────── Errors ───────────────────────────────
    error InvalidState(CampaignState current, CampaignState expected);
    error NotAdvertiser();
    error NotInfluencer();
    error NotParticipant();
    error InvalidAmount();
    error TokenNotWhitelisted();
    error CampaignExpired();
    error CampaignNotExpired();
    error InvalidAuthor(address received, address expected);
    error InvalidSender(address sender, address expected);
    error InvalidWorkflowId(bytes32 received, bytes32 expected);
    error InvalidWorkflowName(bytes10 received, bytes10 expected);
    error FeeTooHigh();

    // ──────────────────────────── Constructor ──────────────────────────
    constructor(
        address _forwarder,
        address[] memory _whitelistedTokens,
        uint256 _platformFeeBps,
        address _feeRecipient
    ) Ownable(msg.sender) {
        if (_platformFeeBps > MAX_FEE_BPS) revert FeeTooHigh();
        forwarderAddress = _forwarder;
        platformFeeBps = _platformFeeBps;
        feeRecipient = _feeRecipient == address(0) ? msg.sender : _feeRecipient;

        for (uint256 i = 0; i < _whitelistedTokens.length; i++) {
            whitelistedTokens[_whitelistedTokens[i]] = true;
            emit TokenWhitelisted(_whitelistedTokens[i]);
        }
    }

    // ──────────────────────────── Campaign Lifecycle ───────────────────

    /// @notice Advertiser creates a campaign and deposits funds in one tx
    /// @dev For ERC20: caller must approve this contract first. For ETH: send msg.value.
    function deposit(
        address token,
        address influencer,
        uint256 amount,
        bytes32 contentHash,
        uint256 minViews,
        uint256 expiryDeadline,
        uint64  campaignDuration
    ) external payable nonReentrant returns (uint256 campaignId) {
        if (!whitelistedTokens[token]) revert TokenNotWhitelisted();
        if (amount == 0) revert InvalidAmount();
        if (expiryDeadline <= block.timestamp) revert CampaignExpired();

        // Calculate fee
        uint256 fee = (amount * platformFeeBps) / 10000;

        // Transfer funds to contract
        if (token == NATIVE_ETH_ADDRESS) {
            if (msg.value != amount) revert InvalidAmount();
        } else {
            IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        }

        campaignId = campaignCounter++;
        campaigns[campaignId] = Campaign({
            advertiser: msg.sender,
            influencer: influencer,
            token: token,
            amount: amount,
            contentHash: contentHash,
            minViews: minViews,
            campaignDuration: campaignDuration,
            deadline: expiryDeadline,
            state: CampaignState.Funded,
            platformFee: fee,
            influencerAccepted: false
        });

        emit CampaignCreated(campaignId, msg.sender, influencer, token, amount, contentHash);
        emit CampaignFunded(campaignId, msg.sender, token, amount);
    }

    /// @notice Influencer accepts a funded campaign
    function acceptCampaign(uint256 campaignId) external {
        Campaign storage c = campaigns[campaignId];
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);
        if (msg.sender != c.influencer) revert NotInfluencer();
        if (block.timestamp >= c.deadline) revert CampaignExpired();

        c.influencerAccepted = true;
        emit CampaignAccepted(campaignId, msg.sender);
    }

    /// @notice Advertiser cancels before influencer accepts
    function cancelCampaign(uint256 campaignId) external nonReentrant {
        Campaign storage c = campaigns[campaignId];
        if (msg.sender != c.advertiser) revert NotAdvertiser();
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);
        if (c.influencerAccepted) revert InvalidState(c.state, CampaignState.Funded);

        c.state = CampaignState.Cancelled;
        _transferOut(c.token, c.advertiser, c.amount);

        emit CampaignCancelled(campaignId, msg.sender);
    }

    /// @notice Either party raises a dispute (only after influencer accepted)
    function raiseDispute(uint256 campaignId) external {
        Campaign storage c = campaigns[campaignId];
        if (msg.sender != c.advertiser && msg.sender != c.influencer) revert NotParticipant();
        if (!c.influencerAccepted) revert InvalidState(c.state, CampaignState.Funded);
        // Allow dispute from Funded (accepted) state only
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);

        c.state = CampaignState.Disputed;
        emit DisputeRaised(campaignId, msg.sender);
    }

    /// @notice Owner resolves dispute — sends funds to influencer or refunds advertiser
    function resolveDispute(uint256 campaignId, bool releaseToInfluencer) external onlyOwner nonReentrant {
        Campaign storage c = campaigns[campaignId];
        if (c.state != CampaignState.Disputed) revert InvalidState(c.state, CampaignState.Disputed);

        if (releaseToInfluencer) {
            c.state = CampaignState.Completed;
            uint256 influencerAmount = c.amount - c.platformFee;
            if (c.platformFee > 0) {
                _transferOut(c.token, feeRecipient, c.platformFee);
            }
            _transferOut(c.token, c.influencer, influencerAmount);
            emit FundsReleased(campaignId, c.influencer, c.token, influencerAmount, c.platformFee);
        } else {
            c.state = CampaignState.Refunded;
            _transferOut(c.token, c.advertiser, c.amount);
            emit FundsRefunded(campaignId, c.advertiser, c.token, c.amount);
        }

        emit DisputeResolved(campaignId, releaseToInfluencer);
    }

    /// @notice Called by CRE via Keystone forwarder — handles release or refund
    function onReport(bytes calldata metadata, bytes calldata report) external {
        if (msg.sender != forwarderAddress) revert InvalidSender(msg.sender, forwarderAddress);

        // Decode metadata: (bytes32 workflowId, bytes10 workflowName, address author)
        (bytes32 workflowId, bytes10 workflowName, address author) =
            abi.decode(metadata, (bytes32, bytes10, address));

        if (workflowId != expectedWorkflowId) revert InvalidWorkflowId(workflowId, expectedWorkflowId);
        if (workflowName != expectedWorkflowName) revert InvalidWorkflowName(workflowName, expectedWorkflowName);
        if (author != expectedAuthor) revert InvalidAuthor(author, expectedAuthor);

        // Decode report: (uint8 action, bytes data)
        // action: 1 = release, 2 = refund
        (uint8 action, bytes memory data) = abi.decode(report, (uint8, bytes));

        if (action == 1) {
            // Release: data = abi.encode(campaignId)
            uint256 campaignId = abi.decode(data, (uint256));
            _releaseFunds(campaignId);
        } else if (action == 2) {
            // Refund: data = abi.encode(campaignId)
            uint256 campaignId = abi.decode(data, (uint256));
            _refundFunds(campaignId);
        }
    }

    /// @notice Advertiser can reclaim funds after deadline if not completed
    function claimExpired(uint256 campaignId) external nonReentrant {
        Campaign storage c = campaigns[campaignId];
        if (msg.sender != c.advertiser) revert NotAdvertiser();
        if (block.timestamp < c.deadline) revert CampaignNotExpired();
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);

        c.state = CampaignState.Refunded;
        _transferOut(c.token, c.advertiser, c.amount);

        emit FundsRefunded(campaignId, c.advertiser, c.token, c.amount);
    }

    // ──────────────────────────── View Functions ───────────────────────

    function getCampaign(uint256 campaignId) external view returns (Campaign memory) {
        return campaigns[campaignId];
    }

    function isExpired(uint256 campaignId) external view returns (bool) {
        return block.timestamp >= campaigns[campaignId].deadline;
    }

    // ──────────────────────────── Admin ────────────────────────────────

    function setPlatformFeeBps(uint256 _feeBps) external onlyOwner {
        if (_feeBps > MAX_FEE_BPS) revert FeeTooHigh();
        platformFeeBps = _feeBps;
    }

    function setFeeRecipient(address _recipient) external onlyOwner {
        feeRecipient = _recipient;
    }

    function whitelistToken(address token) external onlyOwner {
        whitelistedTokens[token] = true;
        emit TokenWhitelisted(token);
    }

    function removeTokenFromWhitelist(address token) external onlyOwner {
        whitelistedTokens[token] = false;
        emit TokenRemovedFromWhitelist(token);
    }

    function setForwarderAddress(address _forwarder) external onlyOwner {
        forwarderAddress = _forwarder;
    }

    function setExpectedAuthor(address _author) external onlyOwner {
        expectedAuthor = _author;
    }

    function setExpectedWorkflowId(bytes32 _id) external onlyOwner {
        expectedWorkflowId = _id;
    }

    function setExpectedWorkflowName(string calldata _name) external onlyOwner {
        expectedWorkflowName = bytes10(bytes(_name));
    }

    // ──────────────────────────── Internal ─────────────────────────────

    function _releaseFunds(uint256 campaignId) internal nonReentrant {
        Campaign storage c = campaigns[campaignId];
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);

        c.state = CampaignState.Completed;
        uint256 influencerAmount = c.amount - c.platformFee;

        if (c.platformFee > 0) {
            _transferOut(c.token, feeRecipient, c.platformFee);
        }
        _transferOut(c.token, c.influencer, influencerAmount);

        emit FundsReleased(campaignId, c.influencer, c.token, influencerAmount, c.platformFee);
    }

    function _refundFunds(uint256 campaignId) internal nonReentrant {
        Campaign storage c = campaigns[campaignId];
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);

        c.state = CampaignState.Refunded;
        _transferOut(c.token, c.advertiser, c.amount);

        emit FundsRefunded(campaignId, c.advertiser, c.token, c.amount);
    }

    function _transferOut(address token, address to, uint256 amount) internal {
        if (token == NATIVE_ETH_ADDRESS) {
            (bool success, ) = to.call{value: amount}("");
            require(success, "ETH transfer failed");
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    // ──────────────────────────── ERC165 ───────────────────────────────

    function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
        return interfaceId == type(IERC165).interfaceId;
    }

    receive() external payable {}
}
