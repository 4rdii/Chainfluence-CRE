# CRE Workflow

## Overview

The CRE workflow (`my-project/my-workflow/`) is a Go program compiled to WASM and deployed on the Chainlink CRE (Compute-Runtime Environment) network. It is triggered on-demand via HTTP, fetches tweet data from the X API using DON consensus, validates the deal criteria, and submits a signed fulfillment report back to the escrow contract.

## How It Works

```
Backend triggers workflow via HTTP POST
    → Workflow reads deal from escrow contract (getDeal)
    → Fetches tweet data from X API v2 (with DON consensus)
    → Checks: edit detection, content hash, min views
    → Submits CreReport via Keystone forwarder → onReport()
    → Contract releases or refunds funds
```

## Verification Checks

| Check | How | Fail behavior |
|-------|-----|--------------|
| Edit detection | `editsRemaining < 5` → tweet was edited | Returns `ActionNone` (no release, no refund) |
| Content match | `keccak256(normalize(tweetText)) == contentHash` | Returns `ActionNone` |
| Min views | `impressionCount >= minViews` | Returns `ActionNone` (retry later) |
| Deadline | `now > campaign.Deadline` | Submits refund report |

> **Note:** `campaignDuration` (post uptime) is calculated and passed to the contract but is not yet enforced as a blocking condition. See TODO in README.

## Simulating Locally

```bash
cd my-project

# Simulate with a campaign ID and tweet URL
cre workflow simulate my-workflow \
  --target=staging-settings \
  --input='{"campaignId":"1","tweetUrl":"https://x.com/user/status/123"}'

# Simulate without tweet URL (will error — URL is required)
cre workflow simulate my-workflow \
  --target=staging-settings \
  --input='{"campaignId":"1"}'
```

## Building the WASM Binary

```bash
cd my-project

# Format and vet first
go fmt ./...
go vet ./...

# Compile to WASM
GOOS=wasip1 GOARCH=wasm go build -o my-workflow/tmp.wasm ./my-workflow
```

The compiled binary (`tmp.wasm` / `binary.wasm.br.b64`) is what gets deployed to CRE.

## Deploying

```bash
cd my-project

# Deploy to staging
cre workflow deploy my-workflow --target=staging-settings

# Deploy to production
cre workflow deploy my-workflow --target=production-settings

# Other lifecycle commands
cre workflow activate my-workflow --target=staging-settings
cre workflow pause my-workflow --target=staging-settings
cre workflow delete my-workflow --target=staging-settings
```

## Triggering a Deployed Workflow

The backend calls the CRE HTTP gateway with a signed JWT. The `http-trigger/` helper handles JWT generation.

**Endpoint:** `https://01.gateway.zone-a.cre.chain.link`

**Payload (JSON-RPC 2.0):**
```json
{
  "jsonrpc": "2.0",
  "method": "workflow_execute",
  "params": {
    "workflowId": "<workflow-id>",
    "input": {
      "campaignId": "42",
      "tweetUrl": "https://x.com/user/status/1234567890"
    }
  },
  "id": 1
}
```

**Authentication:** JWT signed by the private key of an address listed in `authorizedKeys` in the workflow config.

## Configuration

### `config.staging.json` / `config.production.json`

```json
{
  "xApiBearerToken": "",
  "evms": [
    {
      "escrowAddress": "0x09431B4603E4Aa7511dD9341f2852fD031eA8C71",
      "chainName": "ethereum-testnet-sepolia"
    }
  ],
  "authorizedKeys": [
    "0xYourBackendWalletAddress"
  ]
}
```

### Secrets

Bearer token is loaded from CRE secrets at runtime (not stored in config):

```bash
cre secrets set BEARER_TOKEN --target=staging-settings
cre secrets set BEARER_TOKEN --target=production-settings
```

`secrets.yaml` maps the secret name — it does not store the value:
```yaml
secretsNames:
  BEARER_TOKEN:
    - X_BEARER_TOKEN
```

## Authorization

Only addresses listed in `authorizedKeys` can trigger the workflow. To add the backend wallet:

1. Edit `config.staging.json` / `config.production.json`
2. Add the backend wallet address to `authorizedKeys`
3. Rebuild and redeploy

## Adding More Workflows

Each workflow is a separate directory under `my-project/`:

```
my-project/
├── my-workflow/          # X/Twitter verification workflow
│   ├── main.go
│   ├── workflow.yaml
│   ├── config.staging.json
│   └── config.production.json
└── <new-workflow>/       # add new ones here
    ├── main.go
    └── ...
```

```bash
# Initialize a new workflow
cd my-project
cre init --workflow-name <name>
```

All workflows share `project.yaml`, `secrets.yaml`, and the `contracts/` bindings.
