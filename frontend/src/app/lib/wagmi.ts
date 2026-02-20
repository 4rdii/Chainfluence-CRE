import { getDefaultConfig } from '@rainbow-me/rainbowkit'
import { sepolia } from 'wagmi/chains'

export const wagmiConfig = getDefaultConfig({
  appName: 'Chainfluence',
  projectId: import.meta.env.VITE_WALLETCONNECT_PROJECT_ID || 'chainfluence-dev',
  chains: [sepolia],
})

// Single source of truth — redeploy picks this up; no env override so host env can't stick to old contract
export const ESCROW_ADDRESS = '0x09431B4603E4Aa7511dD9341f2852fD031eA8C71' as `0x${string}`

export const CHAIN_ID = parseInt(import.meta.env.VITE_CHAIN_ID || '11155111')

export const MTK_TOKEN = {
  address: '0xeab216ca4381a5c19e751f1471c55b452db0758a' as `0x${string}`,
  symbol: 'MTK',
  name: 'MyToken',
  decimals: 18,
} as const
