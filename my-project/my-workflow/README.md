# Trustless Advertisement Protocol Workflow

This CRE workflow implements a trustless advertisement protocol where:
- Advertisers deposit funds into an escrow contract with campaign details
- Influencers post content on platforms (X/Twitter, Telegram)
- The workflow automatically checks if criteria are met (view counts, time-based)
- Funds are automatically withdrawn to influencers when criteria are satisfied

## Architecture

### Components

1. **Escrow Contract** (`AdEscrow.sol`):
   - Advertisers deposit funds with campaign parameters
   - Stores campaign data (influencer, content text, minimum views, deadline)
   - Allows CRE workflow to withdraw funds when criteria are met

2. **CRE Workflow**:
   - **Cron Trigger**: Periodically checks active campaigns
   - **Twitter API Integration**: Fetches view counts & tweet text from twitterapi.io
   - **Content Validation**: Compares on-chain `contentText` with the live tweet text
   - **Criteria Validation**: Checks if minimum views are met and deadline hasn't passed
   - **Automatic Withdrawal**: Transfers funds to influencer when criteria are satisfied

### Workflow Flow

```
1. Advertiser deposits funds → Escrow Contract
2. Influencer posts content → Platform (X/Telegram)
3. Cron trigger fires → Check campaign criteria
4. Fetch tweet data → twitterapi.io (with consensus)
5. Validate content + criteria → Tweet text matches on-chain `contentText`, views >= minViews, deadline not passed
6. Withdraw funds → Transfer to influencer
```

## Setup

### 1. Contract Deployment

Deploy the `AdEscrow` contract to your target chain and update the address in `config.staging.json`:

```json
{
  "escrowAddress": "0xYourDeployedContractAddress"
}
```

### 2. Generate Contract Bindings

You need to generate Go bindings from the ABI:

```bash
# The contract bindings should be generated using the CRE contract generation tool
# Place the generated bindings in: contracts/evm/src/generated/adescrow/
```

### 3. Configure Twitter API (twitterapi.io)

Add your twitterapi.io API key to secrets:

1. Update `secrets.yaml`:
```yaml
secretsNames:
  staging-settings:
    - X_API_KEY
  production-settings:
    - X_API_KEY
```

2. Set the secret value using CRE CLI:
```bash
cre secrets set X_API_KEY
```

Or set it directly in `config.staging.json` (not recommended for production), or export `X_API_KEY` in your environment/.env file. The workflow will first check config, then secrets, then environment.
```json
{
  "xApiKey": "your-api-key"
}
```

### 4. Provide manual tweet URLs (temporary)

Until the backend supplies tweet URLs automatically, list each campaign ID and corresponding tweet URL in your workflow config:

```json
"manualTweetUrls": [
  { "campaignId": "1", "tweetUrl": "https://x.com/..." }
]
```

### 5. Update General Configuration

Edit `config.staging.json`:
- `checkSchedule`: Cron schedule for periodic checks (e.g., `*/5 * * * * *` for every 5 seconds)
- `xApiBaseUrl`: twitterapi.io base URL (default: `https://api.twitterapi.io`)
- `manualTweetUrls`: Temporary mapping from campaign IDs to tweet URLs
- `escrowAddress`: Deployed escrow contract address
- `chainName`: Target blockchain (e.g., `ethereum-testnet-sepolia`)

## Running the Workflow

### Simulation

From the project root:

```bash
cre workflow simulate my-workflow --target=staging-settings
```

### Deployment

Once tested:

```bash
cre workflow deploy my-workflow --target=staging-settings
```

## Contract Interface

The workflow expects the following contract functions:

- `deposit(uint256 campaignId, address influencer, string contentText, uint256 minViews, uint256 deadline)`: Deposit funds for a campaign
- `getCampaign(uint256 campaignId)`: Get campaign details
- `onReport(bytes metadata, bytes report)`: Called by the Keystone forwarder to process fulfillment reports
- `deliveryAction(uint256 campaignId)`: Trigger manual check (emits event)

## Current Limitations & TODOs

1. **Contract Bindings**: Currently using placeholder - need to generate from ABI
2. **Campaign Tracking**: Currently hardcoded to check campaign ID 1 - need to implement campaign discovery
3. **EVM Log Trigger**: TODO - Add trigger for `DeliveryActionCalled` events
4. **Multiple Platforms**: Currently only supports X/Twitter - can be extended to Telegram, etc.
5. **Error Handling**: Add more robust error handling and retry logic

## Extending the Workflow

### Add Telegram Support

1. Add Telegram API client
2. Implement `fetchTelegramViewCount()` similar to `fetchXViewCount()`
3. Update criteria checking logic

### Add EVM Log Trigger

Listen for `DeliveryActionCalled` events to trigger immediate checks:

```go
// In InitWorkflow, add:
cre.Handler(
    evm.LogTrigger(&evm.LogTriggerConfig{
        Address: escrowAddress,
        EventSignature: "DeliveryActionCalled(uint256)",
    }),
    onDeliveryActionTriggered,
)
```

### Campaign Discovery

Implement logic to discover active campaigns:
- Query contract events for `CampaignDeposited`
- Maintain a list of active campaign IDs
- Filter out fulfilled/withdrawn campaigns

## Security Considerations

- Twitter API calls use consensus aggregation for trustless verification
- Contract withdrawals are only executable by the CRE workflow (enforced on-chain)
- All onchain operations are verified through BFT consensus
- Secrets (API keys) should be stored securely using CRE secrets management

## References

- [CRE Documentation](https://docs.chain.link/cre)
- [CRE SDK Reference](https://docs.chain.link/cre/reference)
- [twitterapi.io Documentation](https://docs.twitterapi.io/)
