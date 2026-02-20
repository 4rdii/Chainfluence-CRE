import { Hono } from 'hono'
import { eq, desc, gte, lte, and, or, ilike, isNull } from 'drizzle-orm'
import { db } from '../db'
import { campaigns, channels, users } from '../db/schema'

const app = new Hono()

// GET /api/marketplace/campaigns — public campaign listings (excludes direct-deal campaigns)
app.get('/campaigns', async (c) => {
  const { search, minBudget, maxBudget, sort } = c.req.query()

  // List draft and funded campaigns so influencers can propose; exclude direct-deal (channel-propose) campaigns
  const conditions = [
    or(eq(campaigns.status, 'draft'), eq(campaigns.status, 'funded')),
    or(eq(campaigns.isDirectDeal, false), isNull(campaigns.isDirectDeal)),
  ]
  if (search) conditions.push(ilike(campaigns.title, `%${search}%`))
  if (minBudget) conditions.push(gte(campaigns.amount, BigInt(minBudget)))
  if (maxBudget) conditions.push(lte(campaigns.amount, BigInt(maxBudget)))

  const orderBy =
    sort === 'budget' ? desc(campaigns.amount) :
    sort === 'ending' ? campaigns.expiryDeadline :
    desc(campaigns.createdAt)

  const rows = await db
    .select()
    .from(campaigns)
    .where(and(...conditions))
    .orderBy(orderBy)
    .limit(50)

  return c.json({ campaigns: rows })
})

// GET /api/marketplace/channels — public influencer listings
app.get('/channels', async (c) => {
  const { search, minFollowers, maxPrice } = c.req.query()

  const conditions = [eq(channels.isActive, true)]
  if (search) conditions.push(ilike(channels.platformHandle, `%${search}%`))
  if (minFollowers) conditions.push(gte(channels.followerCount, parseInt(minFollowers)))
  if (maxPrice) conditions.push(lte(channels.pricePerPost, parseInt(maxPrice)))

  const rows = await db
    .select({
      channel: channels,
      user: {
        walletAddress: users.walletAddress,
        displayName: users.displayName,
      },
    })
    .from(channels)
    .innerJoin(users, eq(channels.userId, users.id))
    .where(and(...conditions))
    .orderBy(desc(channels.followerCount))
    .limit(50)

  return c.json({ channels: rows })
})

export default app
