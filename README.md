# Chainfluence

> Trustless influencer marketing, powered by Chainlink CRE and World ID.

Chainfluence is a decentralized Twitter/X advertising marketplace that removes trust from the advertiser-influencer relationship. Advertisers escrow funds in a smart contract, influencers post the agreed content, and a **Chainlink CRE workflow** running on a Decentralized Oracle Network automatically verifies the post — releasing payment only when all criteria are met. Influencers prove their humanity via **World ID** to prevent bot farms and Sybil attacks.

## Demo Video

[Watch the full demo walkthrough](docs/chainflunce-demo-cre.mp4)

## Live Demo

| | URL |
|--|-----|
| Frontend | [www.chain-fluence.xyz](https://www.chain-fluence.xyz) |
| Backend API | [chainfluence-cre-back.vercel.app/api](https://chainfluence-cre-back.vercel.app/api) |
| Escrow Contract | [`0xDdaD4bfB7338E6b7d5B818Bb721E2B65e3bA2e71`](https://sepolia.etherscan.io/address/0xDdaD4bfB7338E6b7d5B818Bb721E2B65e3bA2e71) (Sepolia) |
| MTK Token | [`0xeab216ca4381a5c19e751f1471c55b452db0758a`](https://sepolia.etherscan.io/address/0xeab216ca4381a5c19e751f1471c55b452db0758a) (Sepolia) |

## Tech Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| Smart Contract | Solidity, Foundry, OpenZeppelin | `AdEscrowV2` — escrow with deal lifecycle, platform fees, CRE forwarder integration |
| CRE Workflow | Go (WASM), Chainlink CRE SDK | Reads deal from chain, fetches tweet via X API with DON consensus, submits signed release/refund report |
| Backend | Hono, Drizzle ORM, Neon Postgres, Vercel Serverless | Auth (SIWE), campaigns/deals/channels CRUD, CRE trigger via signed JWT |
| Frontend | React, wagmi v2, RainbowKit, shadcn/ui, Vite | Full marketplace UI — wallet connect, campaign creation, deal lifecycle, World ID verification |
| Identity | World ID (Orb-level proof) | Sybil-resistant influencer verification — one human per channel |

## How It Works

```
Advertiser                    Smart Contract              CRE Workflow (DON)         X API v2
    |                              |                            |                        |
    |-- create campaign ---------->|                            |                        |
    |-- deposit funds ------------>| [Funded]                   |                        |
    |                              |                            |                        |
Influencer                         |                            |                        |
    |-- verify via World ID        |                            |                        |
    |-- accept deal + tweet URL -->| [Accepted]                 |                        |
    |-- post tweet                 |                            |                        |
    |                              |                            |                        |
Backend                            |                            |                        |
    |-- trigger CRE workflow ------|------------------------->  |                        |
    |                              |                            |                        |
    |                              |<-- getDeal() ------------- |                        |
    |                              |   (contract = immutable    |                        |
    |                              |    source of truth for     |                        |
    |                              |    deal criteria, tweet    |                        |
    |                              |    URL, content hash,      |                        |
    |                              |    min views, duration)    |                        |
    |                              |                            |                        |
    |                              |                            |-- fetch tweet -------> |
    |                              |                            |<- views, text, edits - |
    |                              |                            |   (DON consensus)      |
    |                              |                            |                        |
    |                              |                            | compare content hash   |
    |                              |                            | check edit status      |
    |                              |                            | check post duration    |
    |                              |                            | check view count       |
    |                              |                            |                        |
    |                              |<-- signed report --------- |                        |
    |                              | [Completed] funds released |                        |
```

1. **Advertiser** creates a campaign and deposits tokens into the escrow contract — deal criteria (content hash, min views, duration, deadline, tweet URL) are stored on-chain as the immutable source of truth
2. **Influencer** (verified human via World ID) accepts the deal and provides the tweet URL — also stored on-chain via `acceptDeal()`
3. **Backend** triggers the CRE workflow via HTTP with a signed JWT (only the deal ID is passed — no off-chain data)
4. **CRE Workflow** runs on the DON:
   - **Reads all deal criteria directly from the contract** via `getDeal()` — the contract is the only source of truth, not the backend database
   - **Fetches the live tweet** from the X API with multi-node consensus (median for view count, identical match for content)
   - **Verifies on-chain criteria against live tweet data:**
     - **Content integrity** — `keccak256(normalizedText)` matches the on-chain content hash
     - **Edit detection** — tweet was not modified after posting
     - **Post duration** — tweet has been live for the required number of seconds
     - **View count** — impressions meet the on-chain minimum threshold
     - **Deadline** — refunds automatically if the campaign has expired
5. **Contract** receives the signed fulfillment report via the Keystone forwarder — releases funds to the influencer (minus platform fee) or refunds the advertiser

### Sybil Resistance — World ID

Influencers verify their humanity via [World ID](https://world.org/world-id) (Orb-level proof) when registering a channel. A unique nullifier hash ensures **one human = one channel**, preventing bot farms and duplicate accounts. Verified channels display a "Verified Human" badge throughout the marketplace.

### Authorization Model

All deal and campaign mutation endpoints enforce role-based access control — only the advertiser can fund a deal, only the influencer can submit a tweet, and only involved parties can trigger verification or raise disputes.

## Repository Layout

| Path | Description |
|------|-------------|
| `contracts/` | Foundry project — `AdEscrowV2` escrow contract + deploy scripts |
| `my-project/` | Chainlink CRE workflow (Go + WASM) — tweet verification + fulfillment reports |
| `backend/` | Hono API on Vercel Serverless — auth, campaigns, deals, channels, World ID |
| `frontend/` | React app — wagmi v2, RainbowKit, World ID IDKit, full marketplace UI |
| `docs/` | Extended documentation |

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
  --input='{"campaignId":"1"}'
```

See `.env.example` files in `backend/` and `frontend/` for required environment variables.

## Documentation

- [**Contracts**](docs/contracts.md) — escrow architecture, deal lifecycle, on-chain struct, deployment
- [**CRE Workflow**](docs/workflow.md) — verification logic, simulation, deployment, HTTP trigger
- [**Development**](docs/development.md) — setup, env vars, authorization model, World ID integration
- [**Design System**](docs/design.md) — colors, components, badges, layout, screen specs

## Key Design Decisions

- **Content stored as hash, not string** — the contract stores `keccak256(normalizedText)` to save gas. The CRE workflow applies the same normalization to the live tweet and compares hashes.
- **Tweet URL stored on-chain** — set by the influencer when calling `acceptDeal()`, ensuring the CRE workflow reads the URL from the contract (source of truth), not from off-chain state.
- **DON consensus for tweet data** — multiple CRE nodes independently fetch the X API and reach consensus on view count (median) and content (identical match), preventing single-node manipulation.
- **World ID at the channel level** — verification happens once per channel, not per deal. This balances UX (low friction) with Sybil resistance (one human, one channel).
- **Backend syncs state from chain** — the deal detail endpoint reads on-chain state and updates the database if the contract has moved to a terminal state, keeping off-chain and on-chain in sync.

## Secrets & Security

- Never commit `.env` files — all subdirectories have `.gitignore` entries
- CRE secrets are managed via `cre secrets set <name> --target=<target>`; `secrets.yaml` maps names only
- Error stack traces are stripped in production
- All deal/campaign mutations require authenticated wallet + role authorization

## License

MIT — see [`LICENSE`](LICENSE).
