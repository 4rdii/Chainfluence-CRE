import { createPublicClient, http, type Address } from 'viem'
import { sepolia } from 'viem/chains'

const ESCROW_ADDRESS = (process.env.ESCROW_ADDRESS || '0x09431B4603E4Aa7511dD9341f2852fD031eA8C71') as Address
// Default: public Sepolia RPC (no API key). Override with RPC_URL in env (e.g. Vercel project settings).
const RPC_URL = process.env.RPC_URL || 'https://ethereum-sepolia-rpc.publicnode.com'

// AdEscrowV2.CampaignState enum (matches contract lines 23–30)
const CampaignState = {
  Funded: 0,
  Accepted: 1,
  Completed: 2,
  Refunded: 3,
  Disputed: 4,
  Cancelled: 5,
} as const

// Minimal ABI for getDeal(uint256) -> (Campaign struct) — matches AdEscrowV2.getDeal
const getDealAbi = [
  {
    name: 'getDeal',
    type: 'function',
    stateMutability: 'view',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [
      { name: 'advertiser', type: 'address' },
      { name: 'influencer', type: 'address' },
      { name: 'token', type: 'address' },
      { name: 'amount', type: 'uint256' },
      { name: 'contentHash', type: 'bytes32' },
      { name: 'minViews', type: 'uint256' },
      { name: 'campaignDuration', type: 'uint64' },
      { name: 'deadline', type: 'uint256' },
      { name: 'state', type: 'uint8' },
      { name: 'platformFee', type: 'uint256' },
      { name: 'influencerAccepted', type: 'bool' },
    ],
  },
] as const

const publicClient = createPublicClient({
  chain: sepolia,
  transport: http(RPC_URL),
})

export type OnChainDealState = 'funded' | 'accepted' | 'posted' | 'verifying' | 'completed' | 'refunded' | 'disputed' | 'cancelled'

/**
 * Read deal state from the escrow contract. Returns our app status string or null if read fails.
 */
export async function getDealStateFromChain(dealId: number): Promise<OnChainDealState | null> {
  console.log('[escrow-chain] getDealStateFromChain called', { dealId, escrowAddress: ESCROW_ADDRESS, rpcUrl: RPC_URL })
  try {
    const result = await publicClient.readContract({
      address: ESCROW_ADDRESS,
      abi: getDealAbi,
      functionName: 'getDeal',
      args: [BigInt(dealId)],
    })
    // result is array: [advertiser, influencer, token, amount, contentHash, minViews, campaignDuration, deadline, state, platformFee, influencerAccepted]
    const state = Number(result[8])
    console.log('[escrow-chain] getDeal contract result', { dealId, rawState: state })
    let status: OnChainDealState
    switch (state) {
      case CampaignState.Completed:
        status = 'completed'
        break
      case CampaignState.Refunded:
        status = 'refunded'
        break
      case CampaignState.Disputed:
        status = 'disputed'
        break
      case CampaignState.Cancelled:
        status = 'cancelled'
        break
      case CampaignState.Funded:
        status = 'funded'
        break
      case CampaignState.Accepted:
        status = 'accepted'
        break
      default:
        status = 'funded'
    }
    console.log('[escrow-chain] getDealStateFromChain returning', { dealId, status })
    return status
  } catch (err) {
    const msg = err instanceof Error ? err.message : String(err)
    console.error('[escrow-chain] getDealStateFromChain failed', { dealId, error: msg })
    return null
  }
}
