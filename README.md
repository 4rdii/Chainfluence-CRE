# Chainfluence

> Trustless influencer marketing, powered by Chainlink.

Chainfluence is a decentralized Twitter/X advertising marketplace. Advertisers create campaigns, influencers apply, and payments are held in a smart contract escrow — released automatically when a Chainlink CRE workflow verifies the post met the agreed criteria (views, content, duration).

## Repository Layout

| Path | Description |
|------|-------------|
| `contracts/` | Foundry project — `AdEscrowV2` escrow contract + deploy scripts |
| `my-project/` | Chainlink CRE workflow (Go + WASM) — verifies tweets and submits fulfillment reports |
| `backend/` | Hono API on Vercel Serverless — auth, campaigns, deals, channels, X OAuth |
| `frontend/` | React app — wagmi v2, RainbowKit, full marketplace UI |
| `docs/` | Extended documentation |

## How It Works

1. **Advertiser** creates a campaign (off-chain draft) and proposes or waits for influencer applications
2. **Influencer** accepts a deal; advertiser funds it — locking MTK tokens in the escrow contract
3. **Influencer** posts the agreed tweet and submits the URL
4. **Backend** triggers the CRE workflow via HTTP with the deal ID and tweet URL
5. **Workflow** fetches tweet data via X API (DON consensus), checks content hash, edit status, and view count
6. **Contract** receives the signed fulfillment report — releases funds to influencer or refunds advertiser

## Deployed Contracts (Sepolia)

| Contract | Address |
|----------|---------|
| AdEscrowV2 | `0xDdaD4bfB7338E6b7d5B818Bb721E2B65e3bA2e71` |
| MTK Token | `0xeab216ca4381a5c19e751f1471c55b452db0758a` |

## Quick Start

**Prerequisites:** Go 1.24+, Foundry (`foundryup`), Node.js 18+, `cre` CLI

```bash
# Contracts
cd contracts && forge build

# Backend
cd backend && npm install && npm run dev

# Frontend
cd frontend && npm install && npm run dev

# CRE workflow (simulate)
cd my-project
cre workflow simulate my-workflow \
  --target=staging-settings \
  --input='{"campaignId":"1","tweetUrl":"https://x.com/user/status/123"}'
```

## Documentation

- [**Contracts**](docs/contracts.md) — escrow architecture, deal lifecycle, deployment, ABI reference
- [**CRE Workflow**](docs/workflow.md) — verification logic, simulation, deployment, HTTP trigger
- [**Development**](docs/development.md) — setup, environment variables, coding conventions
- [**Design System**](docs/design.md) — colors, components, layout, screen specifications

## TODO

- [ ] **Enforce post duration in workflow** — `main.go` calculates `postedDuration` but never compares it against `campaign.CampaignDuration` before releasing. Add: `if postedDuration < campaign.CampaignDuration { return ActionNone }`.
- [ ] **Channel owner check on propose** — `POST /api/channels/:id/propose` does not verify the requesting user owns that channel. Add an ownership guard to prevent spoofed proposals.

## Secrets & Security

- Never commit `.env` files — all subdirectories have `.gitignore` entries covering them
- CRE secrets are managed via `cre secrets set <name> --target=<target>`; `secrets.yaml` maps names only
- Rotate keys and rebuild WASM after any contract redeploy or credential rotation

## License

MIT — see [`LICENSE`](LICENSE). Use at your own risk; audit before deploying to production networks.
