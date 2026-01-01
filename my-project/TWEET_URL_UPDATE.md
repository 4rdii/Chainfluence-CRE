# Tweet URL in HTTP Trigger - Update

## Overview

The workflow now accepts the **tweet URL directly in the HTTP trigger payload**, making it fully dynamic and eliminating the need for the `manualTweetUrls` config mapping.

## Changes

### HTTP Trigger Input Structure

**Before:**
```json
{
  "campaignId": "1"
}
```

**After:**
```json
{
  "campaignId": "1",
  "tweetUrl": "https://x.com/username/status/1234567890"
}
```

The `tweetUrl` field is **optional**.

## Tweet URL Resolution Priority

The workflow uses the following priority order to determine which tweet URL to check:

1. **HTTP Trigger Payload** (highest priority)
   - If `tweetUrl` is provided in the HTTP request, it will be used

2. **Config Mapping** (`manualTweetUrls`)
   - Falls back to the config file mapping if no URL in the trigger

3. **On-Chain `contentText`** (lowest priority)
   - Uses the campaign's `contentText` field from the smart contract

4. **Error if none found**
   - Returns error if no tweet URL is available from any source

## Benefits

✅ **Fully Dynamic**: Backend can provide the tweet URL for each request  
✅ **No Config Updates**: No need to maintain `manualTweetUrls` mappings  
✅ **Flexible**: Config mapping and on-chain fallbacks still work  
✅ **Scalable**: Works with unlimited campaigns without config changes  
✅ **Backend Control**: Backend has full control over what content to verify  

## Usage Examples

### Example 1: Backend provides tweet URL (recommended)

```bash
# Simulation
echo '{"campaignId":"1","tweetUrl":"https://x.com/user/status/123"}' | \
  cre workflow simulate my-workflow --target=staging-settings

# Production (HTTP POST to CRE gateway)
{
  "jsonrpc": "2.0",
  "method": "workflow_execute",
  "params": {
    "workflowId": "your-workflow-id",
    "input": {
      "campaignId": "1",
      "tweetUrl": "https://x.com/user/status/123"
    }
  },
  "id": 1
}
```

### Example 2: Fallback to config mapping

```bash
# If tweetUrl is not provided, checks manualTweetUrls in config
echo '{"campaignId":"1"}' | \
  cre workflow simulate my-workflow --target=staging-settings
```

### Example 3: Fallback to on-chain contentText

```json
// If no tweetUrl in trigger and no config mapping,
// uses campaign.contentText from the smart contract
{
  "campaignId": "1"
}
```

## Backend Integration

Your backend should:

1. **Listen for contract events** (CampaignDeposited, DeliveryActionCalled)
2. **Store tweet URLs** in your database mapped to campaign IDs
3. **Include tweet URL** when triggering the workflow:

```javascript
// Example backend code
const triggerWorkflow = async (campaignId, tweetUrl) => {
  const jwt = await generateJWT(privateKey);
  
  await fetch('https://01.gateway.zone-a.cre.chain.link', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${jwt}`
    },
    body: JSON.stringify({
      jsonrpc: '2.0',
      method: 'workflow_execute',
      params: {
        workflowId: process.env.WORKFLOW_ID,
        input: {
          campaignId: campaignId.toString(),
          tweetUrl: tweetUrl
        }
      },
      id: 1
    })
  });
};
```

## Migration Path

The `manualTweetUrls` config is now **optional** and can be removed once your backend provides tweet URLs:

**Current config.staging.json:**
```json
{
  "manualTweetUrls": [
    {"campaignId": "1", "tweetUrl": "https://x.com/..."}
  ]
}
```

**After backend integration (can be removed):**
```json
{
  "manualTweetUrls": []
}
```

## Logging

The workflow logs which source the tweet URL came from:

```
msg="Using tweet URL" 
  url="https://x.com/user/status/123" 
  source="http-trigger"
```

Possible sources:
- `http-trigger` - URL from HTTP request (highest priority)
- `config-mapping` - URL from manualTweetUrls config
- `onchain-contentText` - URL from smart contract
- `unknown` - No URL found (will error)

## Testing

Test all three scenarios:

```bash
# 1. HTTP trigger with URL
echo '{"campaignId":"1","tweetUrl":"https://x.com/test/status/123"}' | \
  cre workflow simulate my-workflow --target=staging-settings

# 2. Config fallback (requires manualTweetUrls in config)
echo '{"campaignId":"1"}' | \
  cre workflow simulate my-workflow --target=staging-settings

# 3. On-chain fallback (requires campaign.contentText to be a valid URL)
echo '{"campaignId":"2"}' | \
  cre workflow simulate my-workflow --target=staging-settings
```

## Summary

The workflow is now **fully event-driven and dynamic**:
- Backend triggers with campaign ID + tweet URL
- No config file updates needed for new campaigns
- Flexible fallback mechanisms for backward compatibility
- Ready for production deployment! 🚀

