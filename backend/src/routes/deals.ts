import { Hono } from 'hono'
import { eq, or, desc, inArray } from 'drizzle-orm'
import { db } from '../db'
import { deals, verifications, campaigns, users } from '../db/schema'
import { authMiddleware, type AppEnv } from '../lib/auth'
import { triggerCREWorkflow } from '../lib/cre-trigger'
import { getDealStateFromChain } from '../lib/escrow-chain'

const app = new Hono<AppEnv>()
app.use('*', authMiddleware)

// GET /api/deals — list my deals (as advertiser or influencer)
app.get('/', async (c) => {
  const { userId } = c.get('user')

  // Find campaign IDs owned by this user (advertiser)
  const myCampaigns = await db
    .select({ id: campaigns.id })
    .from(campaigns)
    .where(eq(campaigns.advertiserId, userId))
  const myCampaignIds = myCampaigns.map((c) => c.id)

  // Get deals where user is influencer OR where the campaign belongs to them
  const conditions = [eq(deals.influencerId, userId)]
  if (myCampaignIds.length > 0) {
    conditions.push(inArray(deals.campaignId, myCampaignIds))
  }

  const rows = await db
    .select({
      deal: deals,
      campaignTitle: campaigns.title,
      campaignAmount: campaigns.amount,
    })
    .from(deals)
    .leftJoin(campaigns, eq(deals.campaignId, campaigns.id))
    .where(or(...conditions))
    .orderBy(desc(deals.createdAt))

  return c.json({
    deals: rows.map((r) => ({
      ...r.deal,
      campaignTitle: r.campaignTitle,
      campaignAmount: r.campaignAmount,
    })),
  })
})

// POST /api/deals — propose a deal
app.post('/', async (c) => {
  const { userId, sub } = c.get('user')
  const body = await c.req.json<{
    campaignId: number
    channelId?: number
  }>()

  const [deal] = await db
    .insert(deals)
    .values({
      campaignId: body.campaignId,
      channelId: body.channelId,
      influencerId: userId,
      proposedBy: sub,
      status: 'proposed',
    })
    .returning()

  return c.json({ deal }, 201)
})

// GET /api/deals/:id — get deal detail + role flags + advertiser wallet; sync status from chain when verifying/posted
app.get('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const rows = await db
    .select({
      deal: deals,
      campaignAdvertiserId: campaigns.advertiserId,
      advertiserWalletAddress: users.walletAddress,
    })
    .from(deals)
    .leftJoin(campaigns, eq(deals.campaignId, campaigns.id))
    .leftJoin(users, eq(campaigns.advertiserId, users.id))
    .where(eq(deals.id, id))
    .limit(1)
  const row = rows[0]
  if (!row?.deal) return c.json({ error: 'Deal not found' }, 404)

  let deal = row.deal

  // When deal has on-chain id and is not already terminal, sync status from contract (source of truth)
  if (deal.onchainCampaignId != null && !['completed', 'refunded', 'disputed', 'cancelled'].includes(deal.status)) {
    console.log('[deals] GET /:id syncing from chain', { dealId: id, onchainCampaignId: deal.onchainCampaignId, dbStatus: deal.status })
    const onChainStatus = await getDealStateFromChain(deal.onchainCampaignId)
    console.log('[deals] GET /:id chain sync result', { dealId: id, onChainStatus, dbStatus: deal.status, willUpdate: Boolean(onChainStatus && ['completed', 'refunded', 'disputed', 'cancelled'].includes(onChainStatus) && onChainStatus !== deal.status) })
    if (onChainStatus && ['completed', 'refunded', 'disputed', 'cancelled'].includes(onChainStatus) && onChainStatus !== deal.status) {
      const [updated] = await db
        .update(deals)
        .set({
          status: onChainStatus,
          updatedAt: new Date(),
          ...(onChainStatus === 'completed' ? { completedAt: new Date() } : {}),
        })
        .where(eq(deals.id, id))
        .returning()
      if (updated) {
        deal = updated
        console.log('[deals] GET /:id updated deal from chain', { dealId: id, newStatus: onChainStatus })
      }
    }
  }

  const isAdvertiser = row.campaignAdvertiserId != null && row.campaignAdvertiserId === userId
  const isInfluencer = deal.influencerId === userId
  return c.json({
    deal,
    isAdvertiser,
    isInfluencer,
    advertiserWalletAddress: row.advertiserWalletAddress ?? undefined,
  })
})

// POST /api/deals/:id/accept — only influencer can accept when deal is channel-originated (campaign.isDirectDeal); otherwise non-proposer (e.g. advertiser for campaign apply) may accept
app.post('/:id/accept', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const [row] = await db
    .select({ deal: deals, isDirectDeal: campaigns.isDirectDeal })
    .from(deals)
    .leftJoin(campaigns, eq(deals.campaignId, campaigns.id))
    .where(eq(deals.id, id))
    .limit(1)
  const deal = row?.deal
  if (!deal) return c.json({ error: 'Deal not found' }, 404)
  if (deal.status !== 'proposed') return c.json({ error: 'Deal is not in proposed state' }, 400)
  // Channel-originated (direct deal): only the influencer (channel owner) may accept
  if (row?.isDirectDeal === true && deal.influencerId !== userId) {
    return c.json({ error: 'Only the influencer can accept this deal' }, 403)
  }
  const [updated] = await db
    .update(deals)
    .set({ status: 'accepted', acceptedAt: new Date(), updatedAt: new Date() })
    .where(eq(deals.id, id))
    .returning()
  if (!updated) return c.json({ error: 'Deal not found' }, 404)
  return c.json({ deal: updated })
})

// POST /api/deals/:id/fund — mark funded after on-chain tx
app.post('/:id/fund', async (c) => {
  const id = parseInt(c.req.param('id'))
  const body = await c.req.json<{ onchainCampaignId: number }>()
  const [updated] = await db
    .update(deals)
    .set({
      status: 'funded',
      onchainCampaignId: body.onchainCampaignId,
      fundedAt: new Date(),
      updatedAt: new Date(),
    })
    .where(eq(deals.id, id))
    .returning()
  if (!updated) return c.json({ error: 'Deal not found' }, 404)
  return c.json({ deal: updated })
})

// POST /api/deals/:id/post — submit tweet URL
app.post('/:id/post', async (c) => {
  const id = parseInt(c.req.param('id'))
  const body = await c.req.json<{ postUrl: string }>()
  const [updated] = await db
    .update(deals)
    .set({ status: 'posted', postUrl: body.postUrl, postedAt: new Date(), updatedAt: new Date() })
    .where(eq(deals.id, id))
    .returning()
  if (!updated) return c.json({ error: 'Deal not found' }, 404)
  return c.json({ deal: updated })
})

// POST /api/deals/:id/verify — trigger CRE workflow
app.post('/:id/verify', async (c) => {
  const id = parseInt(c.req.param('id'))
  const [deal] = await db.select().from(deals).where(eq(deals.id, id)).limit(1)
  if (!deal) return c.json({ error: 'Deal not found' }, 404)
  if (!deal.postUrl) return c.json({ error: 'No tweet URL submitted' }, 400)
  if (!deal.onchainCampaignId) return c.json({ error: 'Deal not funded on-chain' }, 400)

  // Update status to verifying
  await db.update(deals).set({ status: 'verifying', updatedAt: new Date() }).where(eq(deals.id, id))

  // Trigger CRE
  const result = await triggerCREWorkflow({
    campaignId: deal.onchainCampaignId.toString(),
    tweetUrl: deal.postUrl,
  })

  if (!result.success) {
    return c.json({ error: result.error }, 502)
  }

  return c.json({ message: 'Verification triggered', dealId: id })
})

// POST /api/deals/:id/dispute
app.post('/:id/dispute', async (c) => {
  const id = parseInt(c.req.param('id'))
  const [updated] = await db
    .update(deals)
    .set({ status: 'disputed', updatedAt: new Date() })
    .where(eq(deals.id, id))
    .returning()
  if (!updated) return c.json({ error: 'Deal not found' }, 404)
  return c.json({ deal: updated })
})

// GET /api/deals/:id/verifications — verification history
app.get('/:id/verifications', async (c) => {
  const dealId = parseInt(c.req.param('id'))
  const rows = await db
    .select()
    .from(verifications)
    .where(eq(verifications.dealId, dealId))
    .orderBy(desc(verifications.createdAt))
  return c.json({ verifications: rows })
})

export default app
