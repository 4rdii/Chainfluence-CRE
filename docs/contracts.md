# Smart Contracts

## Overview

The escrow contract (`AdEscrowV2`) holds all campaign funds on Ethereum Sepolia until the CRE workflow submits a verified fulfillment report. It uses Chainlink's Keystone forwarder for trustless report delivery.

## Deployed Addresses (Sepolia)

| Contract | Address |
|----------|---------|
| AdEscrowV2 | `0x09431B4603E4Aa7511dD9341f2852fD031eA8C71` |
| MTK Token | `0xeab216ca4381a5c19e751f1471c55b452db0758a` |

## Deal Lifecycle (on-chain states)

```
Funded → Accepted → [Completed | Refunded | Disputed | Cancelled]
```

| State | Value | Description |
|-------|-------|-------------|
| Funded | 0 | Advertiser deposited, waiting for influencer to accept |
| Accepted | 1 | Influencer accepted, can now post |
| Completed | 2 | CRE verified, funds released to influencer |
| Refunded | 3 | Expired or failed verification, funds returned |
| Disputed | 4 | Dispute raised by either party |
| Cancelled | 5 | Advertiser cancelled before influencer accepted |

## Key Functions

```solidity
// Advertiser locks funds (requires ERC-20 approval first)
function deposit(
    uint256 dealId,        // must match backend deal ID
    address token,
    address influencer,
    uint256 amount,
    bytes32 contentHash,   // keccak256(normalizedText)
    uint256 minViews,
    uint256 expiryDeadline,
    uint64  campaignDuration
) external payable

// Influencer accepts a funded deal
function acceptDeal(uint256 dealId) external

// Advertiser cancels before influencer accepts
function cancelDeal(uint256 dealId) external

// Either party raises a dispute
function raiseDispute(uint256 dealId) external

// Owner resolves dispute
function resolveDispute(uint256 dealId, bool releaseToInfluencer) external onlyOwner

// Advertiser reclaims after deadline
function claimExpired(uint256 dealId) external

// Read deal state
function getDeal(uint256 dealId) external view returns (Campaign memory)
```

## CRE Report Processing

The workflow submits reports via the Keystone forwarder, which calls `onReport`. The contract decodes the `CreReport` struct:

```solidity
struct CreReport {
    CreActions action; // Release (1) or Refund (0)
    bytes data;        // abi.encode(dealId)
}
```

- **Release**: transfers `amount - platformFee` to influencer, `platformFee` to fee recipient
- **Refund**: returns full `amount` to advertiser

## Content Hash

Content is never stored as a string on-chain. Instead, a `keccak256` hash of the normalized text is stored:

```
contentHash = keccak256(toLower(trim(collapseWhitespace(content))))
```

The workflow applies the same normalization to the live tweet text and compares hashes. This prevents content tampering without storing the text on-chain.

## Building & Deploying

```bash
cd contracts

# Build
forge build

# Deploy AdEscrowV2 to Sepolia
PRIVATE_KEY=<hex> forge script script/DeployEscrowV2.s.sol:DeployEscrowV2 \
  --rpc-url https://ethereum-sepolia-rpc.publicnode.com \
  --broadcast

# Verify on Etherscan
ETHERSCAN_API_KEY=<key> forge verify-contract \
  <deployed_address> \
  src/AdEscrowV2.sol:AdEscrowV2 \
  --chain sepolia \
  --constructor-args $(cast abi-encode "constructor(address,address[],uint256,address)" \
    <forwarder> [] 250 <fee_recipient>)
```

## Regenerating Contract Bindings (for workflow)

After any ABI change, regenerate the Go bindings used by the CRE workflow:

```bash
# From my-project/
# Place updated ABI at: contracts/evm/src/abi/EscrowV2.abi
# Then regenerate using the CRE contract generation tool
# Output goes to: contracts/evm/src/generated/escrow_v2/
```

## Foundry Reference

```bash
forge build        # compile
forge test         # run tests
forge fmt          # format
forge snapshot     # gas snapshots
anvil              # local node
cast <subcommand>  # chain interaction
```
