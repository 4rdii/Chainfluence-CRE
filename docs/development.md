# Development Guide

## Repository Layout

```
chainlink-cre/
├── contracts/          Foundry project — AdEscrowV2 + deploy scripts
├── my-project/         CRE workflow (Go + WASM)
│   ├── my-workflow/    X/Twitter verification workflow
│   └── contracts/      Go bindings generated from contract ABI
├── backend/            Hono API (Vercel serverless, TypeScript)
├── frontend/           React app (Vite, wagmi v2, RainbowKit)
└── docs/               This documentation
```

## Prerequisites

- **Go** 1.24+ (with WASM target support)
- **Foundry** (`foundryup`) for contracts
- **Node.js** 18+ and npm for backend and frontend
- **`cre` CLI** — install and `cre login` before working on the workflow

## Backend

```bash
cd backend

npm install

# Development (local)
npm run dev

# Build for production
npm run build

# Deploy to Vercel
vercel deploy
```

**Environment variables** (`backend/.env`):
```
DATABASE_URL=postgresql://...
JWT_SECRET=...
X_CLIENT_ID=...
X_CLIENT_SECRET=...
X_REDIRECT_URI=...
CRE_GATEWAY_URL=https://01.gateway.zone-a.cre.chain.link
CRE_WORKFLOW_ID=...
CRE_PRIVATE_KEY=...
ESCROW_ADDRESS=0xDdaD4bfB7338E6b7d5B818Bb721E2B65e3bA2e71
RPC_URL=https://ethereum-sepolia-rpc.publicnode.com
WORLD_ID_APP_ID=app_23fa3c8124a793479d513fa793a3349c
WORLD_ID_ACTION_ID=verify-channel
```

## Frontend

```bash
cd frontend

npm install

# Development server
npm run dev

# Production build
npm run build

# Deploy to Vercel
vercel deploy
```

**Environment variables** (`frontend/.env`):
```
VITE_API_URL=https://your-backend.vercel.app
VITE_WALLETCONNECT_PROJECT_ID=...
VITE_CHAIN_ID=11155111
VITE_WORLD_ID_APP_ID=app_23fa3c8124a793479d513fa793a3349c
VITE_WORLD_ID_ACTION_ID=verify-channel
```

Contract addresses are hardcoded in `frontend/src/app/lib/wagmi.ts` — update there after a redeploy.

## CRE Workflow

```bash
cd my-project

# Format and lint
go fmt ./...
go vet ./...

# Run unit tests
go test ./...

# Build WASM binary
GOOS=wasip1 GOARCH=wasm go build -o my-workflow/tmp.wasm ./my-workflow

# Simulate (requires cre CLI and staging secrets)
cre workflow simulate my-workflow \
  --target=staging-settings \
  --input='{"campaignId":"1","tweetUrl":"https://x.com/user/status/123"}'

# Deploy
cre workflow deploy my-workflow --target=staging-settings
```

## Contracts

```bash
cd contracts

forge build
forge test
forge fmt
```

See [contracts.md](./contracts.md) for deployment instructions.

## Coding Conventions

### Go (workflow)
- Standard `gofmt` style; keep the `//go:build wasip1` build tag intact
- Exported types PascalCase, helpers camelCase
- Wrap errors: `fmt.Errorf("context: %w", err)`
- Structured logging via `slog`
- Config structs in `main.go` must stay aligned with the JSON config files

### TypeScript (backend / frontend)
- Backend uses Hono with Drizzle ORM + Neon Postgres
- All routes live in `backend/src/routes/`; shared logic in `backend/src/lib/`
- Frontend pages in `frontend/src/app/pages/`; reusable components in `frontend/src/app/components/`
- wagmi v2 + @wagmi/core v2 (do not upgrade to v3 — breaks RainbowKit v2)

### Backend Authorization
All deal/campaign mutation endpoints enforce role-based access:

| Endpoint | Allowed by |
|----------|------------|
| `POST /deals/:id/fund` | Advertiser (campaign owner) |
| `POST /deals/:id/post` | Influencer (deal assignee) |
| `POST /deals/:id/verify` | Advertiser or influencer |
| `POST /deals/:id/dispute` | Advertiser or influencer |
| `GET /deals/:id/verifications` | Advertiser or influencer |
| `PATCH /campaigns/:id` | Advertiser (campaign owner) |

Error stack traces are only returned in development (`NODE_ENV !== 'production'`).

### World ID Integration

Influencers can verify their humanity via World ID on the **My Channels** page. The flow:

1. Influencer clicks "Verify Human" → opens IDKit widget (Orb-level proof)
2. Frontend sends proof to `POST /api/channels/:id/verify-worldid`
3. Backend verifies proof with World ID API (`developer.worldcoin.org/api/v2/verify`)
4. Backend checks nullifier uniqueness (one human = one channel)
5. Channel is marked verified → "Human" badge appears in marketplace, channel detail, and my-channels

**DB columns:** `world_id_nullifier` (varchar, unique), `world_id_verified_at` (timestamp) on `channels` table.

**World ID Developer Portal setup:**
1. Create a Cloud app at https://developer.worldcoin.org
2. Create an action with identifier `verify-channel`, max verifications = 1, level = Orb
3. Set `WORLD_ID_APP_ID` and `WORLD_ID_ACTION_ID` env vars on both backend and frontend

### Solidity
- OpenZeppelin imports for SafeERC20, ReentrancyGuard, Ownable
- Custom errors preferred over `require` strings (gas efficient)
- All state-changing functions emit events

## Secrets & Security

- Never commit `.env` files or private keys — all subdirectories have `.gitignore` entries
- CRE secrets are stored via `cre secrets set <name> --target=<target>`; `secrets.yaml` only maps names
- Rotate secrets and rebuild WASM after any contract redeploy or key rotation
- `contracts/.env` and `my-project/.env` contain real keys — keep out of source control

## Commit Style

Imperative, present-tense summaries:
```
add duration check to workflow verification
fix advertiser deal visibility in GET /api/deals
update escrow address after redeployment
```
