//go:build wasip1

package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"my-project/contracts/evm/src/generated/escrow"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/networking/http"
	"github.com/smartcontractkit/cre-sdk-go/cre"
	"github.com/smartcontractkit/cre-sdk-go/cre/wasm"
)

// EvmConfig defines the configuration for a single EVM chain.
type EvmConfig struct {
	EscrowAddress string `json:"escrowAddress"`
	ChainName     string `json:"chainName"`
}

// ManualTweetConfig pairs a campaign ID with the tweet URL to inspect.
type ManualTweetConfig struct {
	CampaignID string `json:"campaignId"`
	TweetURL   string `json:"tweetUrl"`
}

// Config contains workflow configuration.
type Config struct {
	// X (Twitter) API configuration
	XApiBaseUrl string `json:"xApiBaseUrl"`
	// API key for twitterapi.io (should be kept secret)
	XApiKey string `json:"xApiKey"`
	// EVM chain configurations
	Evms []EvmConfig `json:"evms"`
	// Manual tweet URLs (temporary until backend provides them)
	ManualTweetUrls []ManualTweetConfig `json:"manualTweetUrls"`
	// Authorized keys for HTTP trigger (EVM addresses that can trigger the workflow)
	AuthorizedKeys []string `json:"authorizedKeys"`
}

// Campaign represents a campaign from the escrow contract.
// Using the generated type from escrow package
type Campaign = escrow.AdEscrowCampaign

// CampaignState enum values matching contract
const (
	CampaignStateActive    uint8 = 0 // Campaign is active, waiting for criteria to be met
	CampaignStateWithdrawn uint8 = 1 // Criteria met and funds successfully withdrawn
	CampaignStateRefunded  uint8 = 2 // Campaign expired and funds refunded to the advertiser
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

// TwitterAPIResponse represents the response from twitterapi.io for tweet metrics.
type TwitterAPIResponse struct {
	Tweets []struct {
		ID        string `json:"id"`
		ViewCount int64  `json:"viewCount"`
		Text      string `json:"text"`
		CreatedAt string `json:"createdAt"` // Format: "Tue Dec 30 12:08:48 +0000 2025"
	} `json:"tweets"`
}

// TweetObservation is returned from consensus HTTP calls to capture view counts, text, and creation timestamp.
type TweetObservation struct {
	ViewCount     int64  `consensus_aggregation:"median"`
	Text          string `consensus_aggregation:"identical"`
	CreatedAtUnix int64  `consensus_aggregation:"median"` // Unix timestamp in seconds
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
	TweetURL   string `json:"tweetUrl,omitempty"` // Optional: Backend can provide tweet URL directly
}

// InitWorkflow initializes the workflow with HTTP trigger.
func InitWorkflow(config *Config, logger *slog.Logger, secretsProvider cre.SecretsProvider) (cre.Workflow[*Config], error) {
	// Populate API key from secrets if not specified directly in config
	if config.XApiKey == "" && secretsProvider != nil {
		secret, err := secretsProvider.GetSecret(&cre.SecretRequest{Id: "API_KEY"}).Await()
		if err != nil {
			logger.Warn("unable to load X_API_KEY secret", "error", err)
		} else if secret.GetValue() != "" {
			config.XApiKey = secret.GetValue()
		}
	}
	if config.XApiKey == "" {
		if envKey := os.Getenv("X_API_KEY"); envKey != "" {
			config.XApiKey = envKey
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

func (config *Config) tweetURLForCampaign(campaignID *big.Int) string {
	for _, entry := range config.ManualTweetUrls {
		if entry.CampaignID == campaignID.String() {
			return entry.TweetURL
		}
	}
	return ""
}

// getTweetURLSource returns a human-readable description of where the tweet URL came from
func getTweetURLSource(override string, config *Config, campaignID *big.Int, contentText string) string {
	if override != "" {
		return "http-trigger"
	}
	if config.tweetURLForCampaign(campaignID) != "" {
		return "config-mapping"
	}
	if contentText != "" {
		return "onchain-contentText"
	}
	return "unknown"
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
		"triggeredBy", triggeredBy,
	)

	// Parse campaign ID from string to big.Int
	campaignID := new(big.Int)
	if _, ok := campaignID.SetString(input.CampaignID, 10); !ok {
		return nil, fmt.Errorf("invalid campaign ID format: %s", input.CampaignID)
	}

	// Get the first EVM configuration
	evmConfig := config.Evms[0]
	chainSelector, err := evm.ChainSelectorFromName(evmConfig.ChainName)
	if err != nil {
		return nil, fmt.Errorf("invalid chain name: %w", err)
	}

	evmClient := &evm.Client{
		ChainSelector: chainSelector,
	}

	escrowAddress := common.HexToAddress(evmConfig.EscrowAddress)

	// Create escrow contract instance
	escrowContract, err := escrow.NewEscrow(evmClient, escrowAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create escrow contract instance: %w", err)
	}

	// Process the specified campaign with optional tweet URL override
	return processCampaignWithTweetURL(config, runtime, escrowContract, campaignID, input.TweetURL)
}

// processCampaignWithTweetURL checks a campaign's criteria and either releases funds or refunds.
// tweetURLOverride allows the HTTP trigger to provide the tweet URL directly.
func processCampaignWithTweetURL(config *Config, runtime cre.Runtime, escrowContract *escrow.Escrow, campaignID *big.Int, tweetURLOverride string) (*DeliveryActionResult, error) {
	logger := runtime.Logger()
	logger.Info("Processing campaign", "campaignID", campaignID.String())

	// Read campaign data from contract
	// Use nil for blockNumber to get the latest finalized block
	// If the campaign was just created, we might need to wait for finality

	// Encode the call to see the calldata
	codec, err := escrow.NewCodec()
	if err == nil {
		calldata, encodeErr := codec.EncodeGetCampaignMethodCall(escrow.GetCampaignInput{CampaignId: campaignID})
		if encodeErr == nil {
			logger.Info("Calldata for getCampaign",
				"campaignID", campaignID.String(),
				"calldata", fmt.Sprintf("0x%x", calldata),
				"contract", escrowContract.Address.Hex())
		}
	}

	callResult := escrowContract.GetCampaign(runtime, escrow.GetCampaignInput{CampaignId: campaignID}, big.NewInt(-2))

	logger.Info("Calling getCampaign", "campaignID", campaignID.String(), "contract", escrowContract.Address.Hex())

	campaign, err := callResult.Await()
	if err != nil {
		// Check if the error is due to contract revert (campaign doesn't exist)
		errMsg := strings.ToLower(err.Error())
		logger.Info("Error reading campaign", "campaignID", campaignID.String(), "error", err.Error())

		// Check for various error patterns that indicate campaign doesn't exist
		if strings.Contains(errMsg, "campaign does not exist") ||
			strings.Contains(errMsg, "execution reverted") ||
			strings.Contains(errMsg, "attempting to unmarshal an empty string") ||
			strings.Contains(errMsg, "failed to execute capability") {
			// Campaign doesn't exist - this is expected for campaigns that haven't been created yet
			// or if we're querying a finalized block before the campaign was created
			// Return a graceful result instead of failing
			logger.Info("Campaign not found, returning graceful result", "campaignID", campaignID.String())
			return &DeliveryActionResult{
				CampaignID:   campaignID,
				Success:      false,
				ViewsChecked: 0,
				MinViews:     big.NewInt(0),
				Action:       ActionNone,
				Message:      fmt.Sprintf("Campaign %s does not exist yet (may need to wait for block finality)", campaignID.String()),
			}, nil
		}
		return nil, fmt.Errorf("failed to read campaign: %w", err)
	}

	// Check if campaign exists (advertiser address will be zero if campaign doesn't exist)
	if campaign.Advertiser == (common.Address{}) {
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     big.NewInt(0),
			Action:       ActionNone,
			Message:      fmt.Sprintf("Campaign %s does not exist (zero address)", campaignID.String()),
		}, nil
	}

	// Log campaign data for debugging
	logger.Info("Retrieved campaign data",
		"advertiser", campaign.Advertiser.Hex(),
		"influencer", campaign.Influencer.Hex(),
		"token", campaign.Token.Hex(),
		"amount", campaign.Amount.String(),
		"contentText", campaign.ContentText,
		"minViews", campaign.MinViews.String(),
		"campaignDuration", campaign.CampaignDuration,
		"deadline", campaign.Deadline.String(),
		"state", campaign.State,
	)

	// Check campaign state - only process Active campaigns
	if campaign.State != CampaignStateActive {
		stateStr := "unknown"
		switch campaign.State {
		case CampaignStateWithdrawn:
			stateStr = "withdrawn"
		case CampaignStateRefunded:
			stateStr = "refunded"
		}
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      fmt.Sprintf("Campaign already %s", stateStr),
		}, nil
	}

	// Check if deadline has passed - if so, trigger refund
	currentTime := big.NewInt(time.Now().Unix())
	if currentTime.Cmp(campaign.Deadline) > 0 {
		logger.Info("Campaign deadline passed, triggering refund",
			"campaignID", campaignID.String(),
			"deadline", campaign.Deadline.String(),
			"currentTime", currentTime.String(),
		)

		// Encode and submit refund report
		reportPayload, err := encodeRefundReport(campaignID)
		if err != nil {
			return nil, fmt.Errorf("failed to encode refund report: %w", err)
		}

		logger.Info("Refund report encoded", "campaignID", campaignID.String(), "payload", fmt.Sprintf("%x", reportPayload))

		if err := submitReport(runtime, escrowContract, reportPayload); err != nil {
			return nil, fmt.Errorf("failed to submit refund report: %w", err)
		}
		logger.Info("Refund report submitted to forwarder", "campaignID", campaignID.String())

		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      true,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionRefund,
			Message:      "Campaign deadline passed, funds refunded to advertiser",
		}, nil
	}

	// Determine tweet URL to inspect (priority order):
	// 1. HTTP trigger override (if provided)
	// 2. Config mapping (manualTweetUrls)
	// 3. On-chain contentText
	tweetURL := tweetURLOverride
	if tweetURL == "" {
		tweetURL = config.tweetURLForCampaign(campaignID)
	}
	if tweetURL == "" {
		tweetURL = campaign.ContentText
	}
	if tweetURL == "" {
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      "No tweet URL configured for campaign",
		}, nil
	}

	logger.Info("Using tweet URL", "url", tweetURL, "source", getTweetURLSource(tweetURLOverride, config, campaignID, campaign.ContentText))

	// Check if criteria are met by fetching view count from X API
	views, tweetText, createdAtUnix, err := fetchXViewCount(config, runtime, tweetURL)
	if err != nil {
		// If we can't fetch tweet data, it might mean the tweet doesn't exist
		// Log but continue - don't refund automatically here (let caller decide)
		logger.Warn("Failed to fetch tweet data", "error", err.Error())
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      fmt.Sprintf("Failed to fetch tweet data: %s", err.Error()),
		}, nil
	}

	if campaign.ContentText != "" && tweetText != "" && tweetText != campaign.ContentText {
		logger.Warn("Tweet content mismatch",
			"campaignID", campaignID.String(),
			"expectedContent", campaign.ContentText,
			"actualContent", tweetText,
		)
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: views,
			MinViews:     campaign.MinViews,
			Action:       ActionNone,
			Message:      "Tweet content does not match expected campaign content",
		}, nil
	}

	logger.Info("Fetched view count", "views", views, "minViews", campaign.MinViews.String(), "createdAtUnix", createdAtUnix)

	// Check if minimum views are met
	viewsBig := big.NewInt(views)
	if viewsBig.Cmp(campaign.MinViews) >= 0 {
		// Criteria met! Prepare release report payload for Keystone forwarder
		logger.Info("Criteria met, preparing release report", "campaignID", campaignID.String())

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

		reportPayload, err := encodeReleaseReport(campaignID, viewsBig, postedDuration)
		if err != nil {
			return nil, fmt.Errorf("failed to encode release report: %w", err)
		}

		logger.Info("Release report encoded", "campaignID", campaignID.String(), "payload", fmt.Sprintf("%x", reportPayload))

		if err := submitReport(runtime, escrowContract, reportPayload); err != nil {
			return nil, fmt.Errorf("failed to submit release report: %w", err)
		}
		logger.Info("Release report submitted to forwarder", "campaignID", campaignID.String())

		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      true,
			ViewsChecked: views,
			MinViews:     campaign.MinViews,
			Action:       ActionRelease,
			Message:      "Criteria met, funds released to influencer",
		}, nil
	}

	return &DeliveryActionResult{
		CampaignID:   campaignID,
		Success:      false,
		ViewsChecked: views,
		MinViews:     campaign.MinViews,
		Action:       ActionNone,
		Message:      fmt.Sprintf("Criteria not met: %d < %s views", views, campaign.MinViews.String()),
	}, nil
}

// fetchXViewCount fetches the view count, text, and creation timestamp for a tweet from X API.
// Returns: viewCount, text, createdAtUnix (Unix timestamp in seconds), error
func fetchXViewCount(config *Config, runtime cre.Runtime, tweetURL string) (int64, string, int64, error) {

	if config.XApiKey == "" {
		return 0, "", 0, fmt.Errorf("twitter API key not configured")
	}

	// Extract tweet ID from URL
	tweetID, err := extractTweetID(tweetURL)
	if err != nil {
		return 0, "", 0, fmt.Errorf("failed to extract tweet ID: %w", err)
	}

	// Build twitterapi.io request
	apiUrl := fmt.Sprintf("%s/twitter/tweets?tweet_ids=%s", strings.TrimRight(config.XApiBaseUrl, "/"), tweetID)

	// Create HTTP client and send request with consensus
	client := &http.Client{}
	req := &http.Request{
		Url:    apiUrl,
		Method: "GET",
		Headers: map[string]string{
			"X-API-Key": config.XApiKey,
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

			var apiResp TwitterAPIResponse
			if err := json.Unmarshal(resp.Body, &apiResp); err != nil {
				return TweetObservation{}, fmt.Errorf("failed to parse X API response: %w", err)
			}

			if len(apiResp.Tweets) == 0 {
				return TweetObservation{}, fmt.Errorf("tweet %s missing in response", tweetID)
			}

			tweet := apiResp.Tweets[0]
			viewCount := tweet.ViewCount
			tweetText := tweet.Text

			// Parse createdAt timestamp (format: "Tue Dec 30 12:08:48 +0000 2025")
			// Twitter uses RFC1123Z format: "Mon Jan 2 15:04:05 -0700 2006"
			createdAtUnix := int64(0)
			if tweet.CreatedAt != "" {
				parsedTime, err := time.Parse("Mon Jan 2 15:04:05 -0700 2006", tweet.CreatedAt)
				if err != nil {
					logger.Warn("Failed to parse createdAt timestamp", "createdAt", tweet.CreatedAt, "error", err)
					// Continue with createdAtUnix = 0, which will be handled in duration calculation
				} else {
					createdAtUnix = parsedTime.Unix()
				}
			}

			logger.Info("Fetched tweet data from API", "tweetID", tweetID, "views", viewCount, "createdAt", tweet.CreatedAt, "createdAtUnix", createdAtUnix)
			return TweetObservation{
				ViewCount:     viewCount,
				Text:          tweetText,
				CreatedAtUnix: createdAtUnix,
			}, nil
		},
		cre.ConsensusAggregationFromTags[TweetObservation](),
	)

	tweetData, err := tweetDataPromise.Await()
	if err != nil {
		return 0, "", 0, fmt.Errorf("failed to fetch tweet data with consensus: %w", err)
	}

	return tweetData.ViewCount, tweetData.Text, tweetData.CreatedAtUnix, nil
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

// extractTweetID extracts the tweet ID from a Twitter/X URL.
func extractTweetID(url string) (string, error) {
	// Simple extraction - in production, use proper URL parsing
	// Format: https://twitter.com/username/status/1234567890
	// or: https://x.com/username/status/1234567890
	// We'll look for the last numeric segment after /status/

	// This is a simplified version - you'd want to use url.Parse and proper regex
	// For now, we'll assume the URL format is consistent
	statusIndex := -1
	if idx := findSubstring(url, "/status/"); idx != -1 {
		statusIndex = idx + len("/status/")
	} else if idx := findSubstring(url, "/statuses/"); idx != -1 {
		statusIndex = idx + len("/statuses/")
	} else {
		return "", fmt.Errorf("invalid Twitter/X URL format: %s", url)
	}

	// Extract the tweet ID (numeric)
	tweetID := ""
	for i := statusIndex; i < len(url); i++ {
		if url[i] >= '0' && url[i] <= '9' {
			tweetID += string(url[i])
		} else if url[i] == '?' || url[i] == '#' || url[i] == '/' {
			break
		}
	}

	if tweetID == "" {
		return "", fmt.Errorf("could not extract tweet ID from URL: %s", url)
	}

	return tweetID, nil
}

// Helper function to find substring
func findSubstring(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func submitReport(runtime cre.Runtime, escrowContract *escrow.Escrow, payload []byte) error {
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

func abiEncodeUint256Pair(a, b *big.Int) ([]byte, error) {
	uintType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create uint256 ABI type: %w", err)
	}

	args := abi.Arguments{
		{Type: uintType},
		{Type: uintType},
	}

	return args.Pack(a, b)
}

func main() {
	wasm.NewRunner(cre.ParseJSON[Config]).Run(InitWorkflow)
}
