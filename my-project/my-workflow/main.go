//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/url"
	"strings"
	"time"

	"bytes"

	"my-project/contracts/evm/src/generated/escrow_v2"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

// EvmConfig defines the configuration for a single EVM chain.
type EvmConfig struct {
	EscrowAddress string `json:"escrowAddress"`
	ChainName     string `json:"chainName"`
	ChainSelector uint64 `json:"chainSelector,omitempty"`
}

// Config contains workflow configuration.
type Config struct {
	// X API v2 Bearer Token (Official Twitter API)
	XApiBearerToken string `json:"xApiBearerToken"`
	// EVM chain configurations
	Evms []EvmConfig `json:"evms"`
	// Authorized keys for HTTP trigger (EVM addresses that can trigger the workflow)
	AuthorizedKeys []string `json:"authorizedKeys"`
}

// Campaign represents a deal from the escrow V2 contract (same struct shape for deal data).
type Campaign = escrow_v2.AdEscrowV2Campaign

// CampaignState enum values matching AdEscrowV2 contract (6 on-chain states)
const (
	CampaignStateFunded    uint8 = 0 // Advertiser deposited, waiting for influencer to accept
	CampaignStateAccepted  uint8 = 1 // Influencer accepted, can post tweet
	CampaignStateCompleted uint8 = 2 // Verification passed, funds released
	CampaignStateRefunded  uint8 = 3 // Expired or failed, funds returned
	CampaignStateDisputed  uint8 = 4 // Dispute raised by either party
	CampaignStateCancelled uint8 = 5 // Advertiser cancelled before acceptance
)

// CreActions enum values matching contract
const (
	CreActionRefund  uint8 = 0 // Refund the funds to the advertiser
	CreActionRelease uint8 = 1 // Release the funds to the influencer
)

// ActionType represents the action taken by the workflow
type ActionType string

const (
	ActionNone    ActionType = "none"
	ActionRelease ActionType = "release"
	ActionRefund  ActionType = "refund"
)

// XAPIv2Response represents the response from X API v2 for tweet data.
type XAPIv2Response struct {
	Data struct {
		ID            string `json:"id"`
		Text          string `json:"text"`
		CreatedAt     string `json:"created_at"` // ISO 8601: "2025-12-30T12:08:48.000Z"
		PublicMetrics struct {
			ImpressionCount int64 `json:"impression_count"`
			LikeCount       int64 `json:"like_count"`
			RetweetCount    int64 `json:"retweet_count"`
			ReplyCount      int64 `json:"reply_count"`
			BookmarkCount   int64 `json:"bookmark_count"`
		} `json:"public_metrics"`
		EditControls struct {
			EditsRemaining int    `json:"edits_remaining"`
			IsEditEligible bool   `json:"is_edit_eligible"`
			EditableUntil  string `json:"editable_until"`
		} `json:"edit_controls"`
	} `json:"data"`
}

// TweetObservation is returned from consensus HTTP calls to capture view counts, text, and creation timestamp.
type TweetObservation struct {
	ViewCount      int64  `consensus_aggregation:"median"`
	Text           string `consensus_aggregation:"identical"`
	CreatedAtUnix  int64  `consensus_aggregation:"median"`
	IsEditEligible bool   `consensus_aggregation:"identical"`
	EditsRemaining int    `consensus_aggregation:"median"`
}

// DeliveryActionResult represents the result of processing a delivery action.
type DeliveryActionResult struct {
	CampaignID   *big.Int
	Success      bool
	ViewsChecked int64
	MinViews     *big.Int
	Action       ActionType // Action taken: release, refund, or none
	Message      string
}

// HTTPTriggerInput defines the payload structure for HTTP triggers.
type HTTPTriggerInput struct {
	CampaignID string `json:"campaignId"`
	TweetURL   string `json:"tweetUrl,omitempty"`  // Backend provides tweet URL
	ChainName  string `json:"chainName,omitempty"` // Target chain (defaults to first EVM config)
}

// InitWorkflow initializes the workflow with HTTP trigger.
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	// Populate Bearer token from secrets if not specified directly in config
	if config.XApiBearerToken == "" && secretsProvider != nil {
		secret, err := secretsProvider.GetSecret(&cre.SecretRequest{Id: "BEARER_TOKEN"}).Await()
		if err != nil {
			logger.Warn("unable to load X_BEARER_TOKEN secret", "error", err)
		} else if secret.GetValue() != "" {
			config.XApiBearerToken = secret.GetValue()
		}
	}

	// Convert authorized addresses to HTTP trigger keys
	var authorizedKeys []*http.AuthorizedKey
	for _, addr := range config.AuthorizedKeys {
		authorizedKeys = append(authorizedKeys, &http.AuthorizedKey{
			Type:      http.KeyType_KEY_TYPE_ECDSA_EVM,
			PublicKey: addr,
		})
	}

	logger.Info("Initializing workflow with HTTP trigger",
		"authorizedKeys", len(authorizedKeys),
	)

	return cre.Workflow[*Config]{
		// HTTP trigger: Allow external systems to trigger campaign checks on-demand
		cre.Handler(
			http.Trigger(&http.Config{
				AuthorizedKeys: authorizedKeys,
			}),
			onHTTPTrigger,
		),
	}, nil
}

// findEvmConfig looks up EVM config by chain name, defaulting to first entry.
func (config *Config) findEvmConfig(chainName string) EvmConfig {
	if chainName != "" {
		for _, evm := range config.Evms {
			if evm.ChainName == chainName {
				return evm
			}
		}
	}
	return config.Evms[0]
}

// normalizeText deterministically normalizes text for content comparison across consensus nodes.
func normalizeText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return strings.Join(strings.Fields(s), " ")
}

// onHTTPTrigger is triggered when an HTTP request is received.
func onHTTPTrigger(config *Config, runtime cre.Runtime, trigger *http.Payload) (*DeliveryActionResult, error) {
	logger := runtime.Logger()

	// Parse the input payload to get the campaign ID
	var input HTTPTriggerInput
	if err := json.Unmarshal(trigger.Input, &input); err != nil {
		return nil, fmt.Errorf("failed to parse HTTP trigger input: %w", err)
	}

	triggeredBy := "simulation"
	if trigger.Key != nil && trigger.Key.PublicKey != "" {
		triggeredBy = trigger.Key.PublicKey
	}

	logger.Info("HTTP trigger received",
		"campaignID", input.CampaignID,
		"tweetURL", input.TweetURL,
		"chainName", input.ChainName,
		"triggeredBy", triggeredBy,
	)

	// Parse campaign ID from string to big.Int
	campaignID := new(big.Int)
	if _, ok := campaignID.SetString(input.CampaignID, 10); !ok {
		return nil, fmt.Errorf("invalid campaign ID format: %s", input.CampaignID)
	}

	// Resolve EVM config (multi-chain support)
	evmConfig := config.findEvmConfig(input.ChainName)
	chainSelector, err := evm.ChainSelectorFromName(evmConfig.ChainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	escrowAddress := common.HexToAddress(evmConfig.EscrowAddress)

	// Create escrow V2 contract instance
	escrowContract, err := escrow_v2.NewEscrowV2(evmClient, escrowAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create escrow contract instance: %w", err)
	}

	if input.TweetURL == "" {
		return nil, fmt.Errorf("tweetUrl is required in HTTP trigger payload")
	}

	// Process the specified campaign with tweet URL from trigger
	return processCampaignWithTweetURL(config, runtime, escrowContract, campaignID, input.TweetURL)
}

// processCampaignWithTweetURL checks a deal's criteria and either releases funds or refunds.
// tweetURLOverride allows the HTTP trigger to provide the tweet URL directly.
func processCampaignWithTweetURL(config *Config, runtime cre.Runtime, escrowContract *escrow_v2.EscrowV2, dealID *big.Int, tweetURLOverride string) (*DeliveryActionResult, error) {
	logger := runtime.Logger()
	logger.Info("Processing deal", "dealID", dealID.String())

	// Read deal data from contract (latest finalized block)
	codec, err := escrow_v2.NewCodec()
	if err == nil {
		calldata, encodeErr := codec.EncodeGetDealMethodCall(escrow_v2.GetDealInput{DealId: dealID})
		if encodeErr == nil {
			logger.Info("Calldata for getDeal",
				"dealID", dealID.String(),
				"calldata", fmt.Sprintf("0x%x", calldata),
				"contract", escrowContract.Address.Hex())
		}
	}

	callResult := escrowContract.GetDeal(runtime, escrow_v2.GetDealInput{DealId: dealID}, big.NewInt(-2))

	logger.Info("Calling getDeal", "dealID", dealID.String(), "contract", escrowContract.Address.Hex())

	campaign, err := callResult.Await()
	if err != nil {
		errMsg := strings.ToLower(err.Error())
		logger.Info("Error reading deal", "dealID", dealID.String(), "error", err.Error())

		if strings.Contains(errMsg, "deal does not exist") ||
			strings.Contains(errMsg, "execution reverted") ||
			strings.Contains(errMsg, "attempting to unmarshal an empty string") ||
			strings.Contains(errMsg, "failed to execute capability") {
			logger.Info("Deal not found, returning graceful result", "dealID", dealID.String())
			return &DeliveryActionResult{
				CampaignID:   dealID,
				Success:      false,
				ViewsChecked: 0,
				MinViews:     big.NewInt(0),
				Action:       ActionNone,
				Message:      fmt.Sprintf("Deal %s does not exist yet (may need to wait for block finality)", dealID.String()),
			}, nil
		}
		return nil, fmt.Errorf("failed to read deal: %w", err)
	}

	if campaign.Advertiser == (common.Address{}) {
		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     big.NewInt(0),
			Action:       ActionNone,
			Message:      fmt.Sprintf("Deal %s does not exist (zero address)", dealID.String()),
		}, nil
	}

	logger.Info("Retrieved deal data",
		"advertiser", campaign.Advertiser.Hex(),
		"influencer", campaign.Influencer.Hex(),
		"token", campaign.Token.Hex(),
		"amount", campaign.Amount.String(),
		"contentHash", fmt.Sprintf("%x", campaign.ContentHash),
		"minViews", campaign.MinViews.String(),
		"campaignDuration", campaign.CampaignDuration,
		"deadline", campaign.Deadline.String(),
		"state", campaign.State,
	)

	// Check campaign state - only process Funded or Accepted campaigns
	if campaign.State != CampaignStateFunded && campaign.State != CampaignStateAccepted {
		stateNames := map[uint8]string{
			CampaignStateFunded:    "funded",
			CampaignStateAccepted:  "accepted",
			CampaignStateCompleted: "completed",
			CampaignStateRefunded:  "refunded",
			CampaignStateDisputed:  "disputed",
			CampaignStateCancelled: "cancelled",
		}
		stateStr := stateNames[campaign.State]
		if stateStr == "" {
			stateStr = "unknown"
		}
		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      fmt.Sprintf("Deal is %s, not actionable", stateStr),
		}, nil
	}

	// Check if deadline has passed - if so, trigger refund
	currentTime := big.NewInt(time.Now().Unix())
	if currentTime.Cmp(campaign.Deadline) > 0 {
			logger.Info("Deal deadline passed, triggering refund",
				"dealID", dealID.String(),
				"deadline", campaign.Deadline.String(),
			"currentTime", currentTime.String(),
		)

		// Encode and submit refund report
		reportPayload, err := encodeRefundReport(dealID)
		if err != nil {
			return nil, fmt.Errorf("failed to encode refund report: %w", err)
		}

		logger.Info("Refund report encoded", "dealID", dealID.String(), "payload", fmt.Sprintf("%x", reportPayload))

		if err := submitReport(runtime, escrowContract, reportPayload); err != nil {
			return nil, fmt.Errorf("failed to submit refund report: %w", err)
		}
		logger.Info("Refund report submitted to forwarder", "dealID", dealID.String())

		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      true,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionRefund,
			Message:      "Campaign deadline passed, funds refunded to advertiser",
		}, nil
	}

	tweetURL := tweetURLOverride
	if tweetURL == "" {
		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      "No tweet URL provided",
		}, nil
	}

	logger.Info("Using tweet URL", "url", tweetURL)

	// Fetch tweet data from X API v2 (with consensus)
	obs, err := fetchXViewCount(config, runtime, tweetURL)
	if err != nil {
		logger.Warn("Failed to fetch tweet data", "error", err.Error())
		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      fmt.Sprintf("Failed to fetch tweet data: %s", err.Error()),
		}, nil
	}

	views := obs.ViewCount
	createdAtUnix := obs.CreatedAtUnix

	// Edit detection: if edits remaining < 5, the tweet was edited → reject
	if obs.EditsRemaining < 5 {
		logger.Warn("Tweet was edited",
			"dealID", dealID.String(),
			"editsRemaining", obs.EditsRemaining,
			"isEditEligible", obs.IsEditEligible,
		)
		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      false,
			ViewsChecked: views,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      fmt.Sprintf("Tweet was edited (edits remaining: %d/5)", obs.EditsRemaining),
		}, nil
	}

	// Content matching: V2 stores contentHash (keccak256 of normalized text). Compare observed text hash to deal's ContentHash.
	if obs.Text != "" {
		normalized := normalizeText(obs.Text)
		observedHash := crypto.Keccak256([]byte(normalized))
		if len(observedHash) != 32 {
			return &DeliveryActionResult{
				CampaignID:   dealID,
				Success:      false,
				ViewsChecked: views,
				MinViews:     campaign.MinViews,
				Action:       ActionNone,
				Message:      "Content hash computation failed",
			}, nil
		}
		if !bytes.Equal(observedHash, campaign.ContentHash[:]) {
			logger.Warn("Tweet content hash mismatch",
				"dealID", dealID.String(),
				"expectedHash", fmt.Sprintf("%x", campaign.ContentHash),
				"observedHash", fmt.Sprintf("%x", observedHash),
			)
			return &DeliveryActionResult{
				CampaignID:   dealID,
				Success:      false,
				ViewsChecked: views,
				MinViews:     campaign.MinViews,
				Action:       ActionNone,
				Message:      "Tweet content does not match expected deal content",
			}, nil
		}
	}

	logger.Info("Fetched view count", "views", views, "minViews", campaign.MinViews.String(), "createdAtUnix", createdAtUnix)

	// Check if minimum views are met
	viewsBig := big.NewInt(views)
	if viewsBig.Cmp(campaign.MinViews) >= 0 {
		// Criteria met! Prepare release report payload for Keystone forwarder
		logger.Info("Criteria met, preparing release report", "dealID", dealID.String())

		// Calculate postedDuration from actual tweet creation time
		var postedDuration uint64 = 0
		if createdAtUnix > 0 {
			// Calculate duration: current time - tweet creation time
			currentTime := time.Now().Unix()
			durationSeconds := currentTime - createdAtUnix
			if durationSeconds < 0 {
				// This shouldn't happen, but handle edge case
				logger.Warn("Tweet creation time is in the future", "createdAtUnix", createdAtUnix, "currentTime", currentTime)
				postedDuration = 0
			} else {
				postedDuration = uint64(durationSeconds)
			}
			logger.Info("Calculated posting duration", "postedDuration", postedDuration, "createdAtUnix", createdAtUnix, "currentTime", currentTime)
		} else {
			// If createdAt parsing failed, fall back to old behavior
			// Only set duration if campaign requires it (but this will fail validation)
			if campaign.CampaignDuration > 0 {
				logger.Warn("Cannot calculate duration (createdAt parsing failed), but campaign requires duration", "campaignDuration", campaign.CampaignDuration)
				// Set to 0, which will fail the contract validation if duration is required
				postedDuration = 0
			}
		}

		reportPayload, err := encodeReleaseReport(dealID, viewsBig, postedDuration)
		if err != nil {
			return nil, fmt.Errorf("failed to encode release report: %w", err)
		}

		logger.Info("Release report encoded", "dealID", dealID.String(), "payload", fmt.Sprintf("%x", reportPayload))

		if err := submitReport(runtime, escrowContract, reportPayload); err != nil {
			return nil, fmt.Errorf("failed to submit release report: %w", err)
		}
		logger.Info("Release report submitted to forwarder", "dealID", dealID.String())

		return &DeliveryActionResult{
			CampaignID:   dealID,
			Success:      true,
			ViewsChecked: views,
			MinViews:     campaign.MinViews,
			Action:       ActionRelease,
			Message:      "Criteria met, funds released to influencer",
		}, nil
	}

	return &DeliveryActionResult{
		CampaignID:   dealID,
		Success:      false,
		ViewsChecked: views,
		MinViews:     campaign.MinViews,
		Action:       ActionNone,
		Message:      fmt.Sprintf("Criteria not met: %d < %s views", views, campaign.MinViews.String()),
	}, nil
}

// fetchXViewCount fetches tweet data from X API v2 with consensus.
func fetchXViewCount(config *Config, runtime cre.Runtime, tweetURL string) (*TweetObservation, error) {
	if config.XApiBearerToken == "" {
		return nil, fmt.Errorf("X API Bearer token not configured")
	}

	tweetID, err := extractTweetID(tweetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to extract tweet ID: %w", err)
	}

	// X API v2 endpoint with required fields
	apiUrl := fmt.Sprintf("https://api.x.com/2/tweets/%s?tweet.fields=public_metrics,created_at,text,edit_controls", tweetID)

	client := &http.Client{}
	req := &http.Request{
		Url:    apiUrl,
		Method: "GET",
		Headers: map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", config.XApiBearerToken),
		},
	}

	tweetDataPromise := http.SendRequest(
		config,
		runtime,
		client,
		func(config *Config, logger *slog.Logger, sendRequester *http.SendRequester) (TweetObservation, error) {
			resp, err := sendRequester.SendRequest(req).Await()
			if err != nil {
				return TweetObservation{}, fmt.Errorf("failed to get X API response: %w", err)
			}

			var apiResp XAPIv2Response
			if err := json.Unmarshal(resp.Body, &apiResp); err != nil {
				return TweetObservation{}, fmt.Errorf("failed to parse X API v2 response: %w", err)
			}

			if apiResp.Data.ID == "" {
				return TweetObservation{}, fmt.Errorf("tweet %s not found in X API v2 response", tweetID)
			}

			// Parse created_at (ISO 8601 / RFC3339)
			createdAtUnix := int64(0)
			if apiResp.Data.CreatedAt != "" {
				parsedTime, err := time.Parse(time.RFC3339, apiResp.Data.CreatedAt)
				if err != nil {
					logger.Warn("Failed to parse created_at", "created_at", apiResp.Data.CreatedAt, "error", err)
				} else {
					createdAtUnix = parsedTime.Unix()
				}
			}

			logger.Info("Fetched tweet data from X API v2",
				"tweetID", tweetID,
				"views", apiResp.Data.PublicMetrics.ImpressionCount,
				"editsRemaining", apiResp.Data.EditControls.EditsRemaining,
				"createdAtUnix", createdAtUnix,
			)

			return TweetObservation{
				ViewCount:      apiResp.Data.PublicMetrics.ImpressionCount,
				Text:           apiResp.Data.Text,
				CreatedAtUnix:  createdAtUnix,
				IsEditEligible: apiResp.Data.EditControls.IsEditEligible,
				EditsRemaining: apiResp.Data.EditControls.EditsRemaining,
			}, nil
		},
		cre.ConsensusAggregationFromTags[TweetObservation](),
	)

	tweetData, err := tweetDataPromise.Await()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tweet data with consensus: %w", err)
	}

	return &tweetData, nil
}

// encodeReleaseReport encodes a CreReport for releasing funds to influencer.
// Contract expects: CreReport { action: Release, data: abi.encode(campaignId, actualViews, postedDuration) }
func encodeReleaseReport(campaignID, actualViews *big.Int, postedDuration uint64) ([]byte, error) {
	// First encode the inner data: (uint256 campaignId, uint256 actualViews, uint64 postedDuration)
	innerData, err := encodeReleaseData(campaignID, actualViews, postedDuration)
	if err != nil {
		return nil, fmt.Errorf("failed to encode release data: %w", err)
	}

	// Then encode the CreReport struct: (uint8 action, bytes data)
	return encodeCreReport(CreActionRelease, innerData)
}

// encodeRefundReport encodes a CreReport for refunding funds to advertiser.
// Contract expects: CreReport { action: Refund, data: abi.encode(campaignId) }
func encodeRefundReport(campaignID *big.Int) ([]byte, error) {
	// First encode the inner data: (uint256 campaignId)
	innerData, err := encodeRefundData(campaignID)
	if err != nil {
		return nil, fmt.Errorf("failed to encode refund data: %w", err)
	}

	// Then encode the CreReport struct: (uint8 action, bytes data)
	return encodeCreReport(CreActionRefund, innerData)
}

// encodeCreReport encodes a CreReport struct: (uint8 action, bytes data)
// Solidity's abi.decode expects structs with dynamic types to have an outer offset
func encodeCreReport(action uint8, data []byte) ([]byte, error) {
	uint8Type, err := abi.NewType("uint8", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint8 ABI type: %w", err)
	}

	bytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bytes ABI type: %w", err)
	}

	// Encode as tuple: (uint8, bytes)
	args := abi.Arguments{
		{Type: uint8Type}, // action
		{Type: bytesType}, // data
	}

	tupleData, err := args.Pack(action, data)
	if err != nil {
		return nil, fmt.Errorf("failed to pack tuple: %w", err)
	}

	// When Solidity decodes a struct containing dynamic types using abi.decode,
	// it expects an outer offset (0x20 = 32 bytes) pointing to the tuple data
	// Prepend the offset to match Solidity's struct encoding format
	offset := make([]byte, 32)
	offset[31] = 0x20 // Offset of 32 bytes to the tuple data

	return append(offset, tupleData...), nil
}

// encodeReleaseData encodes: (uint256 campaignId, uint256 actualViews, uint64 postedDuration)
func encodeReleaseData(campaignID, actualViews *big.Int, postedDuration uint64) ([]byte, error) {
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint256 ABI type: %w", err)
	}

	uint64Type, err := abi.NewType("uint64", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint64 ABI type: %w", err)
	}

	args := abi.Arguments{
		{Type: uint256Type}, // campaignId
		{Type: uint256Type}, // actualViews
		{Type: uint64Type},  // postedDuration
	}

	return args.Pack(campaignID, actualViews, postedDuration)
}

// encodeRefundData encodes: (uint256 campaignId)
func encodeRefundData(campaignID *big.Int) ([]byte, error) {
	uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint256 ABI type: %w", err)
	}

	args := abi.Arguments{
		{Type: uint256Type}, // campaignId
	}

	return args.Pack(campaignID)
}

// extractTweetID extracts the tweet ID from a Twitter/X URL using proper URL parsing.
// Supports: https://twitter.com/user/status/123, https://x.com/user/status/123
func extractTweetID(rawURL string) (string, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	host := strings.ToLower(parsed.Hostname())
	if host != "twitter.com" && host != "x.com" && host != "www.twitter.com" && host != "www.x.com" && host != "mobile.twitter.com" {
		return "", fmt.Errorf("not a Twitter/X URL: %s", host)
	}

	// Path: /username/status/1234567890 or /username/statuses/1234567890
	parts := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	for i, part := range parts {
		if (part == "status" || part == "statuses") && i+1 < len(parts) {
			tweetID := parts[i+1]
			// Validate it's numeric
			for _, c := range tweetID {
				if c < '0' || c > '9' {
					return "", fmt.Errorf("tweet ID contains non-numeric characters: %s", tweetID)
				}
			}
			if tweetID == "" {
				return "", fmt.Errorf("empty tweet ID in URL: %s", rawURL)
			}
			return tweetID, nil
		}
	}

	return "", fmt.Errorf("could not find /status/ segment in URL: %s", rawURL)
}

func submitReport(runtime cre.Runtime, escrowContract *escrow_v2.EscrowV2, payload []byte) error {
	reportRequest := &cre.ReportRequest{
		EncodedPayload: payload,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	}

	report, err := runtime.GenerateReport(reportRequest).Await()
	if err != nil {
		return fmt.Errorf("failed to generate report: %w", err)
	}

	_, err = escrowContract.WriteReport(runtime, report, nil).Await()
	if err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	return nil
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
