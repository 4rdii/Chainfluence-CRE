# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Chainfluence is a decentralized Twitter/X advertising marketplace. Advertisers fund escrow deals on-chain, influencers post tweets, and Chainlink CRE (Compute Routing Engine) workflows automatically verify tweet content/views/duration and release or refund funds. World ID provides Sybil resistance (one human per channel).

## Repository Structure

Four independent subprojects, each with its own dependencies:

- **`contracts/`** — Solidity escrow contract (Foundry). `AdEscrowV2.sol` receives CRE verification reports via `IReceiverTemplate`.
- **`my-project/`** — CRE workflow in Go, compiled to WASM. Reads deal criteria from the contract, fetches tweet data from X API via DON consensus, verifies content hash + views + duration, and submits a report to release or refund funds.
- **`backend/`** — Hono API deployed as Vercel serverless. Neon Postgres via Drizzle ORM. SIWE authentication. Syncs on-chain state to off-chain DB.
- **`frontend/`** — React + Vite SPA. wagmi v2 + RainbowKit for wallet, World ID IDKit for identity, Tailwind CSS + shadcn/ui for styling.
- **`http-trigger/`** — Standalone TypeScript script to trigger CRE workflows via signed HTTP request.

## Common Commands

### Contracts (Foundry)
```bash
cd contracts
forge build --sizes        # compile
forge fmt --check          # check formatting
forge fmt                  # auto-format
forge test -vvv            # run all tests
forge test --match-test testFunctionName -vvv  # run single test
```

### Backend (Hono)
```bash
cd backend
npm install
npm run dev                # local dev server (tsx watch, reads .env)
npm run build              # production build (tsup → api/ directory)
npx drizzle-kit generate   # generate DB migrations
npx drizzle-kit migrate    # apply migrations
npx drizzle-kit studio     # visual DB explorer
```

### Frontend (Vite + React)
```bash
cd frontend
npm install --legacy-peer-deps   # required due to peer dep conflicts
npm run dev                      # dev server at localhost:5173
npm run build                    # production build → dist/
```

### CRE Workflow (Go → WASM)
```bash
cd my-project
go build ./...             # check compilation
go test ./...              # run tests
# Deployment uses CRE CLI (cre secrets set, cre deploy, etc.)
```

### HTTP Trigger
```bash
cd http-trigger
npm install
npm run trigger            # sends signed request to CRE gateway
```

## Architecture Flow

1. **Advertiser** creates a campaign (backend DB) and funds a deal on-chain (`AdEscrowV2.fundDeal`)
2. **Influencer** accepts the deal, posts a tweet, and submits the tweet URL
3. **Backend** triggers the CRE workflow via ETH-signed JWT to the CRE gateway
4. **CRE workflow** (runs on DON nodes): reads deal from contract → fetches tweet from X API → verifies content hash (`keccak256(normalizedText)`), view count, post duration, edit status → submits signed report on-chain
5. **Contract** processes the CRE report: releases funds to influencer or refunds advertiser

## Key Design Decisions

- Content stored as `keccak256` hash on-chain, not raw text (gas efficiency)
- DON consensus: multiple CRE nodes independently fetch X API data
- World ID verification is per-channel (one-time), not per-deal
- Backend syncs chain state on relevant endpoints; on-chain is source of truth
- Deal amounts stored as wei strings in Postgres to avoid JS number overflow

## Environment Setup

Each subproject has a `.env.example` — copy to `.env` and fill in values:
- **Backend**: `DATABASE_URL`, `JWT_SECRET`, `CRE_TRIGGER_PRIVATE_KEY`, `RPC_URL`, World ID vars
- **Frontend**: `VITE_API_URL`, `VITE_WALLETCONNECT_PROJECT_ID`, `VITE_ESCROW_ADDRESS`, World ID vars
- **HTTP Trigger**: `WORKFLOW_ID`, `PRIVATE_KEY`, `GATEWAY_URL`

Target chain is Sepolia (chain ID 11155111).

## Deployment

- **Backend/Frontend**: Vercel (each has `vercel.json`). Frontend uses `--legacy-peer-deps` install.
- **Contracts**: Foundry scripts in `contracts/script/`. Uses `via_ir = true` in foundry.toml.
- **CRE Workflow**: Deployed via CRE CLI. Config in `my-project/workflow.yaml` with staging/production targets. Secrets managed via `cre secrets set`.

## Extended Documentation

Detailed docs in `docs/`: `contracts.md` (escrow architecture), `workflow.md` (CRE verification logic), `development.md` (full setup guide), `design.md` (UI design system).
