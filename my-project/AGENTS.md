# Repository Guidelines

## Project Structure & Module Organization
- Go module root at `go.mod`; primary workflow code in `my-workflow/` (`main.go`, `workflow.yaml`, env configs `config.staging.json`/`config.production.json`).
- Contract bindings live under `contracts/evm/src/generated/`; regenerate when ABI changes and keep WASM output (`my-workflow/tmp.wasm`) untracked.
- Secrets map is `secrets.yaml`; network targets are set in `project.yaml` and overridden per target in `my-workflow/workflow.yaml`.

## Build, Test, and Development Commands
- `go fmt ./...` then `go vet ./...` — format and run basic static checks.
- `go test ./...` — run unit tests for workflow helpers.
- `GOOS=wasip1 GOARCH=wasm go build -o my-workflow/tmp.wasm ./my-workflow` — compile the workflow WASM with the pinned Go toolchain.
- `cre workflow simulate my-workflow --target=staging-settings` — local run using staging config/secrets.
- `cre workflow deploy my-workflow --target=staging-settings` — deploy after simulation passes.

## Coding Style & Naming Conventions
- Standard Go style (tabs, gofmt); keep the `//go:build wasip1` tag intact for WASM builds.
- Exported types/functions in PascalCase; helpers camelCase. Prefer explicit contexts and wrapped errors for capability calls (`fmt.Errorf("...: %w", err)`).
- Keep config structs in `main.go` aligned with JSON configs; structured logging via `slog`.

## Testing Guidelines
- Place tests next to sources (`my-workflow/foo_test.go`); use table-driven cases for helpers like URL parsing or ABI encoding.
- For onchain interactions, mock/stub EVM client calls where feasible to avoid network dependency.
- Run `cre workflow simulate ...` to validate cron scheduling, HTTP consensus, and secrets resolution before deploying.

## Commit & Pull Request Guidelines
- Commits: imperative, present-tense summaries (e.g., `add cron trigger validation`, `fix escrow report payload`).
- PRs should describe behavior changes, note config/secret impacts, and include logs from `cre workflow simulate` when relevant.
- Run `go fmt`, `go vet`, `go test`, and a staging simulation before review; highlight any ABI or config schema updates.

## Security & Configuration Tips
- Do not commit API keys; manage via `secrets.yaml` and `cre secrets set` or env vars consumed in `InitWorkflow`.
- Verify RPC URLs, chain selectors, and escrow addresses in `project.yaml`, `workflow.yaml`, and `config.*.json` match the target network.
- Rotate secrets and rebuild WASM after capability/library upgrades or contract redeploys.
