# CRE HTTP Trigger Script

Script to trigger CRE workflows via HTTP with JWT authentication.

## Setup

1. Install dependencies:
```bash
npm install
```

2. Copy `.env.example` to `.env` and fill in your values:
```bash
cp .env.example .env
```

3. Edit `.env` with your values:
```env
WORKFLOW_ID=your-64-character-workflow-id
PRIVATE_KEY=your-private-key-hex
CAMPAIGN_ID=1
TWEET_URL=https://x.com/username/status/1234567890
GATEWAY_URL=https://01.gateway.zone-a.cre.chain.link
```

## Usage

```bash
npm run trigger
```

Or directly with tsx:
```bash
npx tsx trigger.ts
```

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `WORKFLOW_ID` | Yes | Your 64-character workflow ID (from deployment output) |
| `PRIVATE_KEY` | Yes | Private key in hex format (with or without 0x prefix) |
| `CAMPAIGN_ID` | Yes | Campaign ID to check |
| `TWEET_URL` | No | Optional tweet URL to verify |
| `GATEWAY_URL` | No | Gateway URL (defaults to https://01.gateway.zone-a.cre.chain.link) |

## Example

```bash
export WORKFLOW_ID="a1b2c3d4e5f67890a1b2c3d4e5f67890a1b2c3d4e5f67890a1b2c3d4e5f67890"
export PRIVATE_KEY="0x1234567890abcdef..."
export CAMPAIGN_ID="2"
export TWEET_URL="https://x.com/user/status/1234567890"
npm run trigger
```

## How It Works

1. Reads environment variables for workflow ID, private key, and campaign details
2. Creates a JSON-RPC 2.0 request with the campaign ID and optional tweet URL
3. Generates a JWT token signed with your private key (ECDSA)
4. Sends authenticated HTTP POST request to CRE gateway
5. Displays the execution ID for tracking in CRE UI

## Security Notes

- **Never commit your `.env` file** - it contains your private key
- Keep your private key secure
- The private key must correspond to an address in the workflow's `authorizedKeys` configuration

## References

- [CRE HTTP Trigger Documentation](https://docs.chain.link/cre/guides/workflow/using-triggers/http-trigger/triggering-deployed-workflows)
- [CRE HTTP Trigger SDK](https://github.com/smartcontractkit/cre-sdk-typescript/tree/main/packages/cre-http-trigger)

