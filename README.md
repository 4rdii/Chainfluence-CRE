## Chainlink CRE Trustless Advertising Prototype

This repo ties together the on-chain escrow contracts and the Chainlink CRE workflow that automatically verifies social metrics and releases advertiser funds to influencers once campaign criteria are satisfied.

### Repository layout

| Path | Description |
| ---- | ----------- |
| `contracts/` | Foundry project containing `AdEscrow` plus deployment/interaction scripts. |
| `my-project/` | CRE workflow (Go + WASM) that polls social metrics, validates content, and submits fulfillment reports through the Keystone forwarder. |
| `install.sh` | Convenience script for installing toolchain prerequisites on fresh environments. |

### Quick start

1. **Install dependencies**
   - Foundry (`foundryup`)
   - Go 1.24+
   - `cre` CLI (login + init)

2. **Contracts**
   ```bash
   cd contracts
   forge build
   PRIVATE_KEY=<hex> forge script script/DeployEscrow.s.sol:DeployEscrow \
     --rpc-url <rpc> --broadcast
   ```

3. **CRE workflow**
   ```bash
   cd my-project
   cre workflow simulate my-workflow --target=staging-settings
   # Once validated:
   cre workflow deploy my-workflow --target=production-settings
   ```

### Secrets & environment variables

- `contracts/.env` and `my-project/.env` contain **real** private keys/API keys and are already listed in their respective `.gitignore` files. Keep them out of source control and rotate any key that may have been exposed.
- For CRE, store API keys with `cre secrets set <name> --target=<target>` rather than embedding them in configs.
- Never commit `.env`, `secrets.yaml` values, or broadcast logs that may leak private data.

### License

This project is released under the MIT License (see [`LICENSE`](LICENSE)). Use at your own risk; audit/modify before deploying to production networks.
