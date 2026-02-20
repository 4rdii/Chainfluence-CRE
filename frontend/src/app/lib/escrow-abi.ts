// AdEscrowV2: dealId is backend deal.id, passed in when depositing (no counter)
export const escrowAbi = [
  {
    type: 'function',
    name: 'deposit',
    inputs: [
      { name: 'dealId', type: 'uint256' },
      { name: 'token', type: 'address' },
      { name: 'influencer', type: 'address' },
      { name: 'amount', type: 'uint256' },
      { name: 'contentHash', type: 'bytes32' },
      { name: 'minViews', type: 'uint256' },
      { name: 'expiryDeadline', type: 'uint256' },
      { name: 'campaignDuration', type: 'uint64' },
    ],
    outputs: [],
    stateMutability: 'payable',
  },
  {
    type: 'function',
    name: 'acceptDeal',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'cancelDeal',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'raiseDispute',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'getDeal',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [
      {
        name: '',
        type: 'tuple',
        components: [
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
    ],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'isExpired',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'claimExpired',
    inputs: [{ name: 'dealId', type: 'uint256' }],
    outputs: [],
    stateMutability: 'nonpayable',
  },
  {
    type: 'event',
    name: 'DealCreated',
    inputs: [
      { name: 'dealId', type: 'uint256', indexed: true },
      { name: 'advertiser', type: 'address', indexed: true },
      { name: 'influencer', type: 'address', indexed: true },
      { name: 'token', type: 'address', indexed: false },
      { name: 'amount', type: 'uint256', indexed: false },
      { name: 'contentHash', type: 'bytes32', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'FundsReleased',
    inputs: [
      { name: 'dealId', type: 'uint256', indexed: true },
      { name: 'influencer', type: 'address', indexed: true },
      { name: 'token', type: 'address', indexed: false },
      { name: 'influencerAmount', type: 'uint256', indexed: false },
      { name: 'feeAmount', type: 'uint256', indexed: false },
    ],
  },
  {
    type: 'event',
    name: 'FundsRefunded',
    inputs: [
      { name: 'dealId', type: 'uint256', indexed: true },
      { name: 'advertiser', type: 'address', indexed: true },
      { name: 'token', type: 'address', indexed: false },
      { name: 'amount', type: 'uint256', indexed: false },
    ],
  },
] as const

export const erc20Abi = [
  {
    type: 'function',
    name: 'approve',
    inputs: [
      { name: 'spender', type: 'address' },
      { name: 'amount', type: 'uint256' },
    ],
    outputs: [{ name: '', type: 'bool' }],
    stateMutability: 'nonpayable',
  },
  {
    type: 'function',
    name: 'allowance',
    inputs: [
      { name: 'owner', type: 'address' },
      { name: 'spender', type: 'address' },
    ],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
  {
    type: 'function',
    name: 'balanceOf',
    inputs: [{ name: 'account', type: 'address' }],
    outputs: [{ name: '', type: 'uint256' }],
    stateMutability: 'view',
  },
] as const
