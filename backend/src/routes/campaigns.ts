import { Hono } from 'hono'
import { eq, desc, and, inArray, sql, or, isNull } from 'drizzle-orm'
import { db } from '../db'
import { campaigns, users, deals } from '../db/schema'
import { authMiddleware, type AppEnv } from '../lib/auth'

const app = new Hono<AppEnv>()

// All routes require auth
app.use('*', authMiddleware)

// GET /api/campaigns — list my campaigns (with deal count). Excludes direct-deal campaigns (proposed via channel); those only show as deals.
app.get('/', async (c) => {
  const { userId } = c.get('user')
  const rows = await db
    .select()
    .from(campaigns)
    .where(
      and(
        eq(campaigns.advertiserId, userId),
        or(eq(campaigns.isDirectDeal, false), isNull(campaigns.isDirectDeal))
      )
    )
    .orderBy(desc(campaigns.createdAt))
  if (rows.length === 0) return c.json({ campaigns: [] })
  const campaignIds = rows.map((r) => r.id)
  const countRows = await db
    .select({ campaignId: deals.campaignId, count: sql<number>`count(*)::int` })
    .from(deals)
    .where(inArray(deals.campaignId, campaignIds))
    .groupBy(deals.campaignId)
  const countByCampaign = Object.fromEntries(countRows.map((r) => [r.campaignId, r.count]))
  const campaignsWithCount = rows.map((camp) => ({
    ...camp,
    dealCount: countByCampaign[camp.id] ?? 0,
  }))
  return c.json({ campaigns: campaignsWithCount })
})

// POST /api/campaigns — create a draft campaign
app.post('/', async (c) => {
  try {
    const { userId } = c.get('user')
    const body = await c.req.json<{
      title: string
      contentText: string
      contentHash?: string
      tokenAddress: string
      amount: string
      minViews?: number
      campaignDuration?: number
      expiryDeadline: string
      categories?: string[]
    }>()

    const [campaign] = await db
      .insert(campaigns)
      .values({
        advertiserId: userId,
        title: body.title,
        contentText: body.contentText,
        contentHash: body.contentHash,
        tokenAddress: body.tokenAddress,
        amount: body.amount,
        minViews: body.minViews ?? 0,
        campaignDuration: body.campaignDuration ?? 0,
        expiryDeadline: new Date(body.expiryDeadline),
        categories: body.categories ?? [],
        status: 'draft',
      })
      .returning()

    return c.json({ campaign }, 201)
  } catch (err: any) {
    console.error('POST /api/campaigns error:', err)
    return c.json({ error: err.message, detail: err.stack?.split('\n')[0] }, 500)
  }
})

// GET /api/campaigns/:id/deals — list all deals for this campaign (campaign can have multiple deals)
app.get('/:id/deals', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const [campaign] = await db.select().from(campaigns).where(eq(campaigns.id, id)).limit(1)
  if (!campaign) return c.json({ error: 'Campaign not found' }, 404)
  // Only campaign owner (advertiser) or involved influencers can see deals
  if (campaign.advertiserId !== userId) {
    const myDeals = await db.select().from(deals).where(and(eq(deals.campaignId, id), eq(deals.influencerId, userId))).orderBy(desc(deals.createdAt))
    return c.json({ deals: myDeals })
  }
  const rows = await db.select().from(deals).where(eq(deals.campaignId, id)).orderBy(desc(deals.createdAt))
  return c.json({ deals: rows })
})

// GET /api/campaigns/:id — get campaign detail (includes advertiser wallet for deal view)
app.get('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const rows = await db
    .select({
      campaign: campaigns,
      advertiserWalletAddress: users.walletAddress,
    })
    .from(campaigns)
    .leftJoin(users, eq(campaigns.advertiserId, users.id))
    .where(eq(campaigns.id, id))
    .limit(1)
  const row = rows[0]
  if (!row?.campaign) return c.json({ error: 'Campaign not found' }, 404)
  return c.json({
    campaign: {
      ...row.campaign,
      advertiserWalletAddress: row.advertiserWalletAddress ?? undefined,
    },
  })
})

// PATCH /api/campaigns/:id — update after deposit (advertiser only)
app.patch('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const body = await c.req.json<{
    onchainCampaignId?: number
    txHash?: string
    status?: string
  }>()

  // Verify caller is the campaign's advertiser
  const [campaign] = await db.select().from(campaigns).where(eq(campaigns.id, id)).limit(1)
  if (!campaign) return c.json({ error: 'Campaign not found' }, 404)
  if (campaign.advertiserId !== userId) return c.json({ error: 'Only the advertiser can update this campaign' }, 403)

  const [updated] = await db
    .update(campaigns)
    .set({ ...body, updatedAt: new Date() })
    .where(eq(campaigns.id, id))
    .returning()

  if (!updated) return c.json({ error: 'Campaign not found' }, 404)
  return c.json({ campaign: updated })
})

export default app
