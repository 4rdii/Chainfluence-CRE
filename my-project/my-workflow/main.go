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
	"github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"
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
	// Cron schedule for periodic campaign checks
	CheckSchedule string `json:"checkSchedule"`
	// X (Twitter) API configuration
	XApiBaseUrl string `json:"xApiBaseUrl"`
	// API key for twitterapi.io (should be kept secret)
	XApiKey string `json:"xApiKey"`
	// EVM chain configurations
	Evms []EvmConfig `json:"evms"`
	// Manual tweet URLs (temporary until backend provides them)
	ManualTweetUrls []ManualTweetConfig `json:"manualTweetUrls"`
}

// Campaign represents a campaign from the escrow contract.
// Using the generated type from escrow package
type Campaign = escrow.AdEscrowCampaign

// TwitterAPIResponse represents the response from twitterapi.io for tweet metrics.
type TwitterAPIResponse struct {
	Tweets []struct {
		ID        string `json:"id"`
		ViewCount int64  `json:"viewCount"`
		Text      string `json:"text"`
	} `json:"tweets"`
}

// TweetObservation is returned from consensus HTTP calls to capture view counts and text.
type TweetObservation struct {
	ViewCount int64  `consensus_aggregation:"median"`
	Text      string `consensus_aggregation:"identical"`
}

// DeliveryActionResult represents the result of processing a delivery action.
type DeliveryActionResult struct {
	CampaignID   *big.Int
	Success      bool
	ViewsChecked int64
	MinViews     *big.Int
	Withdrawn    bool
	Message      string
}

// InitWorkflow initializes the workflow with multiple triggers.
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

	return cre.Workflow[*Config]{
		// Cron trigger: Periodically check all active campaigns
		cre.Handler(
			cron.Trigger(&cron.Config{Schedule: config.CheckSchedule}),
			onPeriodicCheck,
		),
		// TODO: Add EVM log trigger for DeliveryActionCalled events
		// This would require the event trigger capability which may need additional setup
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

// onPeriodicCheck is triggered by cron to check all active campaigns.
func onPeriodicCheck(config *Config, runtime cre.Runtime, trigger *cron.Payload) (*DeliveryActionResult, error) {
	logger := runtime.Logger()
	logger.Info("Periodic check triggered", "time", time.Now())

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

	// For now, we'll process campaigns one by one
	// In production, you'd want to track campaign IDs or query them from events
	// This is a simplified version - you'd need to implement campaign ID tracking

	// Example: Check campaign ID 1 (in production, iterate through active campaigns)
	campaignID := big.NewInt(1)
	return processCampaign(config, runtime, escrowContract, campaignID)
}

// processCampaign checks a campaign's criteria and withdraws if met.
func processCampaign(config *Config, runtime cre.Runtime, escrowContract *escrow.Escrow, campaignID *big.Int) (*DeliveryActionResult, error) {
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

	callResult := escrowContract.GetCampaign(runtime, escrow.GetCampaignInput{CampaignId: campaignID}, big.NewInt(-3))

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
				Withdrawn:    false,
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
			Withdrawn:    false,
			Message:      fmt.Sprintf("Campaign %s does not exist (zero address)", campaignID.String()),
		}, nil
	}

	// Log campaign data for debugging
	logger.Info("Retrieved campaign data",
		"advertiser", campaign.Advertiser.Hex(),
		"influencer", campaign.Influencer.Hex(),
		"amount", campaign.Amount.String(),
		"contentText", campaign.ContentText,
		"minViews", campaign.MinViews.String(),
		"deadline", campaign.Deadline.String(),
		"fulfilled", campaign.Fulfilled,
		"withdrawn", campaign.Withdrawn,
	)

	// Check if already withdrawn
	if campaign.Withdrawn {
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     big.NewInt(0),
			Withdrawn:    false,
			Message:      "Campaign already withdrawn",
		}, nil
	}

	// Check if deadline has passed
	currentTime := big.NewInt(time.Now().Unix())
	if currentTime.Cmp(campaign.Deadline) > 0 {
		logger.Info("Campaign deadline passed", "deadline", campaign.Deadline.String())
		// Optionally, you might want to refund advertiser or handle differently
	}

	// Determine tweet URL to inspect
	tweetURL := config.tweetURLForCampaign(campaignID)
	if tweetURL == "" {
		tweetURL = campaign.ContentText
	}
	if tweetURL == "" {
		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      false,
			ViewsChecked: 0,
			MinViews:     campaign.MinViews,
			Withdrawn:    false,
			Message:      "No tweet URL configured for campaign",
		}, nil
	}

	// Check if criteria are met by fetching view count from X API
	views, tweetText, err := fetchXViewCount(config, runtime, tweetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch view count: %w", err)
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
			Withdrawn:    false,
			Message:      "Tweet content does not match expected campaign content",
		}, nil
	}

	logger.Info("Fetched view count", "views", views, "minViews", campaign.MinViews.String())

	// Check if minimum views are met
	viewsBig := big.NewInt(views)
	if viewsBig.Cmp(campaign.MinViews) >= 0 {
		// Criteria met! Prepare report payload for Keystone forwarder
		logger.Info("Criteria met, preparing report payload", "campaignID", campaignID.String())

		reportPayload, err := encodeFulfillmentReport(campaignID, viewsBig)
		if err != nil {
			return nil, fmt.Errorf("failed to encode fulfillment report: %w", err)
		}

		logger.Info("Fulfillment report encoded", "campaignID", campaignID.String(), "payload", fmt.Sprintf("%x", reportPayload))

		if err := submitReport(runtime, escrowContract, reportPayload); err != nil {
			return nil, fmt.Errorf("failed to submit report: %w", err)
		}
		logger.Info("Report submitted to forwarder", "campaignID", campaignID.String())

		return &DeliveryActionResult{
			CampaignID:   campaignID,
			Success:      true,
			ViewsChecked: views,
			MinViews:     campaign.MinViews,
			Withdrawn:    true,
			Message:      "Funds withdrawn successfully",
		}, nil
	}

	return &DeliveryActionResult{
		CampaignID:   campaignID,
		Success:      false,
		ViewsChecked: views,
		MinViews:     campaign.MinViews,
		Withdrawn:    false,
		Message:      fmt.Sprintf("Criteria not met: %d < %s views", views, campaign.MinViews.String()),
	}, nil
}

// fetchXViewCount fetches the view count and text for a tweet from X API.
func fetchXViewCount(config *Config, runtime cre.Runtime, tweetURL string) (int64, string, error) {

	if config.XApiKey == "" {
		return 0, "", fmt.Errorf("twitter API key not configured")
	}

	// Extract tweet ID from URL
	tweetID, err := extractTweetID(tweetURL)
	if err != nil {
		return 0, "", fmt.Errorf("failed to extract tweet ID: %w", err)
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

			viewCount := apiResp.Tweets[0].ViewCount
			tweetText := apiResp.Tweets[0].Text
			logger.Info("Fetched tweet data from API", "tweetID", tweetID, "views", viewCount)
			return TweetObservation{
				ViewCount: viewCount,
				Text:      tweetText,
			}, nil
		},
		cre.ConsensusAggregationFromTags[TweetObservation](),
	)

	tweetData, err := tweetDataPromise.Await()
	if err != nil {
		return 0, "", fmt.Errorf("failed to fetch tweet data with consensus: %w", err)
	}

	return tweetData.ViewCount, tweetData.Text, nil
}

func encodeFulfillmentReport(campaignID, actualViews *big.Int) ([]byte, error) {
	return abiEncodeUint256Pair(campaignID, actualViews)
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
