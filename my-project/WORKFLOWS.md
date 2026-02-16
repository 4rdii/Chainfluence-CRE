# Multiple Workflows in CRE Project

This project contains **multiple workflows** that can be managed independently. Each workflow is a separate directory with its own configuration and code.

## Project Structure

```
my-project/
├── project.yaml              # Project-level settings (shared RPCs, etc.)
├── secrets.yaml              # Shared secrets
├── my-workflow/              # Twitter/X workflow
│   ├── workflow.yaml         # Workflow-specific settings
│   ├── main.go               # Workflow code
│   ├── config.staging.json   # Staging config
│   └── config.production.json
└── telegram-workflow/        # Telegram workflow
    ├── workflow.yaml
    ├── main.go
    ├── config.staging.json
    └── config.production.json
```

## Current Workflows

### 1. `my-workflow` - Twitter/X Escrow Workflow
- **Purpose**: Validates Twitter/X posts for advertisement campaigns
- **Trigger**: HTTP trigger with `campaignId` and optional `tweetUrl`
- **Validates**: View counts, content text, duration, deadline

### 2. `telegram-workflow` - Telegram Escrow Workflow
- **Purpose**: Validates Telegram channel posts for advertisement campaigns
- **Trigger**: HTTP trigger with `campaignId`, `chatId`, and `messageId`
- **Validates**: View counts, content text, duration, subscriber count, deadline

## Creating a New Workflow

### Option 1: Using `cre init` (Recommended)

```bash
cd /home/ardeshir/Desktop/chainlink-cre/my-project
cre init --workflow-name <workflow-name>
```

This will create a new workflow directory with template files.

### Option 2: Manual Creation

1. **Create workflow directory**:
   ```bash
   mkdir -p <workflow-name>
   cd <workflow-name>
   ```

2. **Create `workflow.yaml`** (see `telegram-workflow/workflow.yaml` for example)

3. **Create `main.go`** with your workflow code

4. **Create config files** (`config.staging.json`, `config.production.json`)

## Working with Multiple Workflows

### Simulate a Workflow

```bash
# Simulate Twitter workflow
cre workflow simulate my-workflow --target=staging-settings --input='{"campaignId":"1","tweetUrl":"https://x.com/..."}'

# Simulate Telegram workflow
cre workflow simulate telegram-workflow --target=staging-settings --input='{"campaignId":"1","chatId":"@channel","messageId":123}'
```

### Deploy a Workflow

```bash
# Deploy Twitter workflow
cre workflow deploy my-workflow --target=staging-settings

# Deploy Telegram workflow
cre workflow deploy telegram-workflow --target=staging-settings
```

## Workflow Commands Summary

All CRE workflow commands work with the workflow directory name:

| Command | Description | Example |
|---------|-------------|---------|
| `cre workflow simulate <workflow-dir>` | Test workflow locally | `cre workflow simulate my-workflow --target=staging-settings` |
| `cre workflow deploy <workflow-dir>` | Deploy to CRE network | `cre workflow deploy telegram-workflow --target=staging-settings` |
| `cre workflow activate <workflow-dir>` | Activate deployed workflow | `cre workflow activate my-workflow --target=staging-settings` |
| `cre workflow pause <workflow-dir>` | Pause a workflow | `cre workflow pause my-workflow --target=staging-settings` |
| `cre workflow delete <workflow-dir>` | Delete workflow | `cre workflow delete my-workflow --target=staging-settings` |

## Shared Resources

- **`project.yaml`**: Shared RPC endpoints and project-level settings
- **`secrets.yaml`**: Shared secrets (API keys, tokens, etc.)
- **`contracts/`**: Shared contract bindings (both workflows use the same Escrow contract)

## Workflow-Specific Configuration

Each workflow has its own:
- **`workflow.yaml`**: Workflow registry name and artifact paths
- **`config.staging.json`**: Staging environment configuration
- **`config.production.json`**: Production environment configuration
- **`main.go`**: Workflow implementation code

## Notes

- Each workflow is **independent** - they can be deployed, updated, and managed separately
- Workflows can share the same contract addresses (as configured in their config files)
- Each workflow has its own **workflow-name** in the CRE registry (defined in `workflow.yaml`)
- The `--target` flag selects which settings to use (staging-settings or production-settings)
