// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {IReceiverTemplate} from "./IReceiverTemplate.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title AdEscrowV2
 * @notice Trustless escrow for Twitter ad campaigns with platform fees,
 *         accept/cancel lifecycle, dispute resolution, and CRE verification.
 * @dev Extends IReceiverTemplate for Chainlink CRE (Keystone) forwarder integration.
 */
contract AdEscrowV2 is IReceiverTemplate, ReentrancyGuard {
    using SafeERC20 for IERC20;

    // ──────────────────────────── Constants ────────────────────────────
    address public constant NATIVE_ETH_ADDRESS = address(0);
    uint256 public constant MAX_FEE_BPS = 1000; // 10% cap

    // ──────────────────────────── State ────────────────────────────────
    enum CampaignState {
        Funded,     // 0 — advertiser deposited, waiting for influencer to accept
        Accepted,   // 1 — influencer accepted, can now post
        Completed,  // 2 — verification passed, funds released to influencer
        Refunded,   // 3 — expired or failed verification, funds returned
        Disputed,   // 4 — dispute raised by either party
        Cancelled   // 5 — advertiser cancelled before influencer accepted
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
        uint256 platformFee;       // fee amount (calculated at deposit)
        bool    influencerAccepted;
        string  channelName;       // influencer's X handle — set at deposit time
        string  tweetUrl;          // posted tweet URL — set when influencer accepts
    }

    struct DepositParams {
        uint256 dealId;
        address token;
        address influencer;
        uint256 amount;
        bytes32 contentHash;
        uint256 minViews;
        uint256 expiryDeadline;
        uint64  campaignDuration;
        string  channelName;
    }

    // CRE report actions
    enum CreActions {
        Refund, // 0 — refund funds to advertiser
        Release // 1 — release funds to influencer
    }

    struct CreReport {
        CreActions action;
        bytes data;
    }

    uint256 public platformFeeBps;  // e.g. 250 = 2.5%
    address public feeRecipient;

    // dealId is provided by the backend (e.g. backend deal.id) so it matches off-chain
    mapping(uint256 => Campaign) public deals;
    mapping(address => bool) public whitelistedTokens;

    // ──────────────────────────── Events ───────────────────────────────
    event DealCreated(
        uint256 indexed dealId,
        address indexed advertiser,
        address indexed influencer,
        address token,
        uint256 amount,
        bytes32 contentHash,
        string  channelName
    );
    event DealAccepted(uint256 indexed dealId, address indexed influencer, string tweetUrl);
    event DealCancelled(uint256 indexed dealId, address indexed advertiser);
    event FundsReleased(
        uint256 indexed dealId,
        address indexed influencer,
        address token,
        uint256 influencerAmount,
        uint256 feeAmount
    );
    event FundsRefunded(
        uint256 indexed dealId,
        address indexed advertiser,
        address token,
        uint256 amount
    );
    event DisputeRaised(uint256 indexed dealId, address indexed raisedBy);
    event DisputeResolved(uint256 indexed dealId, bool releasedToInfluencer);
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
    error AlreadyAccepted();
    error FeeTooHigh();
    error DealAlreadyExists();

    // ──────────────────────────── Constructor ──────────────────────────
    constructor(
        address _forwarder,
        address[] memory _whitelistedTokens,
        uint256 _platformFeeBps,
        address _feeRecipient
    ) {
        if (_platformFeeBps > MAX_FEE_BPS) revert FeeTooHigh();
        forwarderAddress = _forwarder;
        platformFeeBps = _platformFeeBps;
        feeRecipient = _feeRecipient == address(0) ? msg.sender : _feeRecipient;

        // Native ETH always whitelisted
        whitelistedTokens[NATIVE_ETH_ADDRESS] = true;

        for (uint256 i = 0; i < _whitelistedTokens.length; i++) {
            whitelistedTokens[_whitelistedTokens[i]] = true;
            emit TokenWhitelisted(_whitelistedTokens[i]);
        }
    }

    // ──────────────────────────── Campaign Lifecycle ───────────────────

    /// @notice Advertiser deposits funds for a deal (dealId must match backend deal.id)
    /// @dev For ERC20: caller must approve this contract first. For ETH: send msg.value.
    function deposit(DepositParams calldata p) external payable nonReentrant {
        if (!whitelistedTokens[p.token]) revert TokenNotWhitelisted();
        if (p.amount == 0) revert InvalidAmount();
        if (p.expiryDeadline <= block.timestamp) revert CampaignExpired();
        if (deals[p.dealId].advertiser != address(0)) revert DealAlreadyExists();

        uint256 fee = (p.amount * platformFeeBps) / 10000;

        if (p.token == NATIVE_ETH_ADDRESS) {
            if (msg.value != p.amount) revert InvalidAmount();
        } else {
            require(msg.value == 0, "Cannot send ETH with ERC20 deposit");
            IERC20(p.token).safeTransferFrom(msg.sender, address(this), p.amount);
        }

        deals[p.dealId] = Campaign({
            advertiser: msg.sender,
            influencer: p.influencer,
            token: p.token,
            amount: p.amount,
            contentHash: p.contentHash,
            minViews: p.minViews,
            campaignDuration: p.campaignDuration,
            deadline: p.expiryDeadline,
            state: CampaignState.Funded,
            platformFee: fee,
            influencerAccepted: false,
            channelName: p.channelName,
            tweetUrl: ""
        });

        emit DealCreated(p.dealId, msg.sender, p.influencer, p.token, p.amount, p.contentHash, p.channelName);
    }

    /// @notice Influencer accepts a funded deal and provides the posted tweet URL
    function acceptDeal(uint256 dealId, string calldata tweetUrl) external {
        Campaign storage c = deals[dealId];
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);
        if (msg.sender != c.influencer) revert NotInfluencer();
        if (block.timestamp >= c.deadline) revert CampaignExpired();

        c.state = CampaignState.Accepted;
        c.influencerAccepted = true;
        c.tweetUrl = tweetUrl;
        emit DealAccepted(dealId, msg.sender, tweetUrl);
    }

    /// @notice Advertiser cancels before influencer accepts
    function cancelDeal(uint256 dealId) external nonReentrant {
        Campaign storage c = deals[dealId];
        if (msg.sender != c.advertiser) revert NotAdvertiser();
        if (c.state != CampaignState.Funded) revert InvalidState(c.state, CampaignState.Funded);
        if (c.influencerAccepted) revert AlreadyAccepted();

        c.state = CampaignState.Cancelled;
        _transferOut(c.token, c.advertiser, c.amount);

        emit DealCancelled(dealId, msg.sender);
    }

    /// @notice Either party raises a dispute (only after influencer accepted)
    function raiseDispute(uint256 dealId) external {
        Campaign storage c = deals[dealId];
        if (msg.sender != c.advertiser && msg.sender != c.influencer) revert NotParticipant();
        if (!c.influencerAccepted) revert InvalidState(c.state, CampaignState.Funded);
        if (c.state == CampaignState.Completed || c.state == CampaignState.Refunded
            || c.state == CampaignState.Cancelled || c.state == CampaignState.Disputed) {
            revert InvalidState(c.state, CampaignState.Accepted);
        }

        c.state = CampaignState.Disputed;
        emit DisputeRaised(dealId, msg.sender);
    }

    /// @notice Owner resolves dispute — sends funds to influencer or refunds advertiser
    function resolveDispute(uint256 dealId, bool releaseToInfluencer) external onlyOwner nonReentrant {
        Campaign storage c = deals[dealId];
        if (c.state != CampaignState.Disputed) revert InvalidState(c.state, CampaignState.Disputed);

        if (releaseToInfluencer) {
            c.state = CampaignState.Completed;
            uint256 influencerAmount = c.amount - c.platformFee;
            if (c.platformFee > 0) {
                _transferOut(c.token, feeRecipient, c.platformFee);
            }
            _transferOut(c.token, c.influencer, influencerAmount);
            emit FundsReleased(dealId, c.influencer, c.token, influencerAmount, c.platformFee);
        } else {
            c.state = CampaignState.Refunded;
            _transferOut(c.token, c.advertiser, c.amount);
            emit FundsRefunded(dealId, c.advertiser, c.token, c.amount);
        }

        emit DisputeResolved(dealId, releaseToInfluencer);
    }

    /// @notice Advertiser reclaims funds after deadline if not completed
    function claimExpired(uint256 dealId) external nonReentrant {
        Campaign storage c = deals[dealId];
        if (msg.sender != c.advertiser) revert NotAdvertiser();
        if (block.timestamp < c.deadline) revert CampaignNotExpired();
        if (c.state == CampaignState.Completed || c.state == CampaignState.Refunded
            || c.state == CampaignState.Cancelled) {
            revert InvalidState(c.state, CampaignState.Funded);
        }

        c.state = CampaignState.Refunded;
        _transferOut(c.token, c.advertiser, c.amount);

        emit FundsRefunded(dealId, c.advertiser, c.token, c.amount);
    }

    // ──────────────────────────── View Functions ───────────────────────

    function getDeal(uint256 dealId) external view returns (Campaign memory) {
        require(deals[dealId].advertiser != address(0), "Deal does not exist");
        return deals[dealId];
    }

    function isExpired(uint256 dealId) external view returns (bool) {
        return block.timestamp >= deals[dealId].deadline;
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

    // ──────────────────────────── CRE Report Processing ──────────────

    /// @inheritdoc IReceiverTemplate
    function _processReport(bytes calldata report) internal override {
        CreReport memory creReport = abi.decode(report, (CreReport));

        if (creReport.action == CreActions.Release) {
            uint256 dealId = abi.decode(creReport.data, (uint256));
            _releaseFunds(dealId);
        } else if (creReport.action == CreActions.Refund) {
            uint256 dealId = abi.decode(creReport.data, (uint256));
            _refundFunds(dealId);
        } else {
            revert("Invalid CRE action");
        }
    }

    // ──────────────────────────── Internal ─────────────────────────────

    function _releaseFunds(uint256 dealId) internal {
        Campaign storage c = deals[dealId];
        // Allow release from Accepted state (influencer accepted + CRE verified)
        if (c.state != CampaignState.Accepted && c.state != CampaignState.Funded) {
            revert InvalidState(c.state, CampaignState.Accepted);
        }

        c.state = CampaignState.Completed;
        uint256 influencerAmount = c.amount - c.platformFee;

        if (c.platformFee > 0) {
            _transferOut(c.token, feeRecipient, c.platformFee);
        }
        _transferOut(c.token, c.influencer, influencerAmount);

        emit FundsReleased(dealId, c.influencer, c.token, influencerAmount, c.platformFee);
    }

    function _refundFunds(uint256 dealId) internal {
        Campaign storage c = deals[dealId];
        if (c.state == CampaignState.Completed || c.state == CampaignState.Refunded
            || c.state == CampaignState.Cancelled) {
            revert InvalidState(c.state, CampaignState.Funded);
        }

        c.state = CampaignState.Refunded;
        _transferOut(c.token, c.advertiser, c.amount);

        emit FundsRefunded(dealId, c.advertiser, c.token, c.amount);
    }

    function _transferOut(address token, address to, uint256 amount) internal {
        if (token == NATIVE_ETH_ADDRESS) {
            (bool success, ) = to.call{value: amount}("");
            require(success, "ETH transfer failed");
        } else {
            IERC20(token).safeTransfer(to, amount);
        }
    }

    receive() external payable {}
}
