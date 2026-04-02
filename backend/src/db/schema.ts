import {
  pgTable,
  serial,
  text,
  varchar,
  integer,
  bigint,
  boolean,
  timestamp,
  jsonb,
  uniqueIndex,
} from 'drizzle-orm/pg-core'

export const users = pgTable(
  'users',
  {
    id: serial('id').primaryKey(),
    walletAddress: varchar('wallet_address', { length: 42 }).notNull(),
    role: varchar('role', { length: 20 }).notNull().default('both'), // advertiser | influencer | both
    displayName: varchar('display_name', { length: 100 }),
    nonce: varchar('nonce', { length: 64 }),
    createdAt: timestamp('created_at').defaultNow().notNull(),
    updatedAt: timestamp('updated_at').defaultNow().notNull(),
  },
  (table) => [uniqueIndex('wallet_address_idx').on(table.walletAddress)]
)

export const channels = pgTable('channels', {
  id: serial('id').primaryKey(),
  userId: integer('user_id')
    .references(() => users.id)
    .notNull(),
  platform: varchar('platform', { length: 20 }).notNull().default('twitter'),
  platformHandle: varchar('platform_handle', { length: 100 }).notNull(),
  platformUserId: varchar('platform_user_id', { length: 100 }),
  profileImageUrl: varchar('profile_image_url', { length: 500 }),
  followerCount: integer('follower_count').default(0),
  categories: jsonb('categories').$type<string[]>().default([]),
  pricePerPost: integer('price_per_post').default(0), // in cents USD
  isActive: boolean('is_active').default(true),
  verifiedAt: timestamp('verified_at'),
  worldIdNullifier: varchar('world_id_nullifier', { length: 100 }),
  worldIdVerifiedAt: timestamp('world_id_verified_at'),
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})

export const campaigns = pgTable('campaigns', {
  id: serial('id').primaryKey(),
  onchainCampaignId: integer('onchain_campaign_id'), // null until deposited
  advertiserId: integer('advertiser_id')
    .references(() => users.id)
    .notNull(),
  title: varchar('title', { length: 200 }).notNull(),
  contentText: text('content_text').notNull(),
  contentHash: varchar('content_hash', { length: 66 }), // 0x + 64 hex chars
  tokenAddress: varchar('token_address', { length: 42 }).notNull(),
  amount: varchar('amount', { length: 78 }).notNull(), // raw wei string — avoids bigint overflow for 18-decimal tokens
  minViews: integer('min_views').default(0),
  campaignDuration: integer('campaign_duration').default(0), // seconds
  expiryDeadline: timestamp('expiry_deadline').notNull(),
  chainId: integer('chain_id').default(11155111), // Sepolia
  status: varchar('status', { length: 20 }).notNull().default('draft'), // draft | funded | active | completed | refunded | cancelled | disputed
  txHash: varchar('tx_hash', { length: 66 }),
  categories: jsonb('categories').$type<string[]>().default([]),
  isDirectDeal: boolean('is_direct_deal').default(false), // true when created via channel propose — hidden from campaign list, only deal is shown
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})

export const deals = pgTable('deals', {
  id: serial('id').primaryKey(),
  campaignId: integer('campaign_id')
    .references(() => campaigns.id)
    .notNull(),
  channelId: integer('channel_id').references(() => channels.id),
  influencerId: integer('influencer_id')
    .references(() => users.id)
    .notNull(),
  status: varchar('status', { length: 20 }).notNull().default('proposed'), // proposed | accepted | funded | posted | verifying | completed | refunded | disputed
  proposedBy: varchar('proposed_by', { length: 42 }).notNull(), // wallet address
  postUrl: text('post_url'),
  onchainCampaignId: integer('onchain_campaign_id'),
  acceptedAt: timestamp('accepted_at'),
  fundedAt: timestamp('funded_at'),
  postedAt: timestamp('posted_at'),
  completedAt: timestamp('completed_at'),
  createdAt: timestamp('created_at').defaultNow().notNull(),
  updatedAt: timestamp('updated_at').defaultNow().notNull(),
})

export const verifications = pgTable('verifications', {
  id: serial('id').primaryKey(),
  dealId: integer('deal_id')
    .references(() => deals.id)
    .notNull(),
  status: varchar('status', { length: 20 }).notNull().default('pending'), // pending | passed | failed
  viewsChecked: integer('views_checked'),
  contentMatched: boolean('content_matched'),
  durationChecked: boolean('duration_checked'),
  isEdited: boolean('is_edited'),
  action: varchar('action', { length: 20 }), // release | refund | none
  message: text('message'),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})

export const waitlist = pgTable(
  'waitlist',
  {
    id: serial('id').primaryKey(),
    email: varchar('email', { length: 255 }).notNull(),
    role: varchar('role', { length: 20 }).default('unknown'), // influencer | advertiser | unknown
    createdAt: timestamp('created_at').defaultNow().notNull(),
  },
  (table) => [uniqueIndex('waitlist_email_idx').on(table.email)]
)

export const chainEvents = pgTable('chain_events', {
  id: serial('id').primaryKey(),
  eventName: varchar('event_name', { length: 100 }).notNull(),
  txHash: varchar('tx_hash', { length: 66 }).notNull(),
  blockNumber: bigint('block_number', { mode: 'bigint' }),
  rawData: jsonb('raw_data'),
  processed: boolean('processed').default(false),
  createdAt: timestamp('created_at').defaultNow().notNull(),
})
