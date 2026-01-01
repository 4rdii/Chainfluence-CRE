# HTTP Trigger Migration Guide

## Overview

The workflow has been successfully converted from a **cron-based trigger** (polling every 10 seconds) to an **HTTP trigger** (event-driven, on-demand execution).

## What Changed

### 1. **Trigger Mechanism**

**Before (Cron):**
- Workflow ran automatically every 5-10 seconds
- Checked campaign ID 1 on every execution
- No external control over when checks happen

**After (HTTP):**
- Workflow is triggered on-demand via HTTP POST requests
- Backend/external system specifies which campaign to check
- More efficient: only runs when needed
- Better control and scalability

### 2. **Code Changes**

#### Main Workflow (`main.go`)

**Removed:**
- `import "github.com/smartcontractkit/cre-sdk-go/capabilities/scheduler/cron"`
- `CheckSchedule` field from `Config` struct
- `onPeriodicCheck()` function
- Cron trigger handler

**Added:**
- `HTTPTriggerInput` struct to define the payload format
- `AuthorizedKeys` field to `Config` struct
- `onHTTPTrigger()` function to handle HTTP requests
- HTTP trigger handler with authorization

#### Configuration Files

**Removed:**
- `checkSchedule` field (was `"*/5 * * * * *"`)

**Added:**
- `authorizedKeys` array with your linked wallet address: `"0xFF3f9d0E76D30577D7A0bD35E14259a76078230e"`

### 3. **Workflow Execution Flow**

```
┌─────────────────┐
│  Backend/API    │
│  (Authorized)   │
└────────┬────────┘
         │ HTTP POST
         │ {"campaignId": "1"}
         │ + JWT Signature
         ▼
┌────────────────────┐
│  CRE Gateway       │
│  Validates JWT     │
└────────┬───────────┘
         │ Authorized ✓
         ▼
┌────────────────────┐
│  Workflow Executes │
│  - Fetch campaign  │
│  - Check views     │
│  - Validate        │
│  - Withdraw funds  │
└────────────────────┘
```

## How to Use

### Testing Locally (Simulation)

**Basic test (uses config mapping or on-chain contentText):**
```bash
cre workflow simulate my-workflow --target=staging-settings --input='{"campaignId":"1"}'
```

**Test with tweet URL override:**
```bash
cre workflow simulate my-workflow --target=staging-settings --input='{"campaignId":"1","tweetUrl":"https://x.com/username/status/1234567890"}'
```

The `--input` flag simulates the HTTP trigger payload.

### Triggering Deployed Workflows

After deployment, trigger the workflow with an HTTP request:

**Endpoint:**
```
https://01.gateway.zone-a.cre.chain.link
```

**Request Body (JSON-RPC 2.0):**
```json
{
  "jsonrpc": "2.0",
  "method": "workflow_execute",
  "params": {
    "workflowId": "your-workflow-id-here",
    "input": {
      "campaignId": "123",
      "tweetUrl": "https://x.com/username/status/1234567890"
    }
  },
  "id": 1
}
```

**Input Parameters:**
- `campaignId` (required): The campaign ID to check
- `tweetUrl` (optional): The tweet URL to verify. If not provided, the workflow will:
  1. Check the `manualTweetUrls` config mapping
  2. Fall back to the on-chain `contentText` field

**Authentication:**
The request must include a JWT token signed by the private key of an authorized address (the one linked earlier: `0xFF3f9d0E76D30577D7A0bD35E14259a76078230e`).

See the [CRE HTTP Trigger documentation](https://docs.chain.link/cre/guides/workflow/using-triggers/http-trigger/overview-go) for JWT generation details.

## Benefits of HTTP Trigger

1. **Cost Efficient**: Only runs when campaigns need checking, not on a fixed schedule
2. **Event-Driven**: Backend can trigger immediately when:
   - A new campaign is created
   - An influencer reports completion
   - Deadline is approaching
3. **Scalable**: Can handle multiple campaigns without increasing base load
4. **Flexible**: Campaign ID is provided dynamically
5. **Secure**: Only authorized addresses can trigger the workflow

## Authorization

The workflow is configured to accept triggers from:
- `0xFF3f9d0E76D30577D7A0bD35E14259a76078230e` (your linked account)

To add more authorized addresses, update the `authorizedKeys` array in `config.staging.json` or `config.production.json`.

## Integration with Backend

Your backend should:

1. **Monitor contract events** for new campaigns or `DeliveryActionCalled` events
2. **Generate JWT tokens** using the authorized wallet's private key
3. **Send HTTP POST** to the CRE gateway with the campaign ID
4. **Handle responses** and retry logic as needed

Example workflow:
```
Contract emits DeliveryActionCalled(campaignId=5)
  ↓
Backend listens for event
  ↓
Backend generates JWT with private key
  ↓
Backend POSTs to CRE gateway with campaignId=5
  ↓
Workflow executes and checks campaign 5
```

## Next Steps

1. ✅ Workflow converted to HTTP trigger
2. ✅ Configuration updated
3. ✅ Documentation updated
4. 🔲 Test simulation with `--input` flag
5. 🔲 Deploy to staging
6. 🔲 Implement backend integration
7. 🔲 Test end-to-end flow

## References

- [CRE HTTP Trigger Overview](https://docs.chain.link/cre/guides/workflow/using-triggers/http-trigger/overview-go)
- [CRE HTTP Trigger Configuration](https://docs.chain.link/cre/guides/workflow/using-triggers/http-trigger/overview-go)
- [JWT Authentication Guide](https://docs.chain.link/cre/guides/workflow/using-triggers/http-trigger/triggering-deployed-workflows-go)

