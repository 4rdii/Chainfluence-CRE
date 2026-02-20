import { Hono } from 'hono'
import { eq, desc } from 'drizzle-orm'
import { db } from '../db'
import { channels, campaigns, deals } from '../db/schema'
import { authMiddleware, type AppEnv } from '../lib/auth'

const app = new Hono<AppEnv>()
app.use('*', authMiddleware)

// GET /api/channels — list my channels
app.get('/', async (c) => {
  const { userId } = c.get('user')
  const rows = await db
    .select()
    .from(channels)
    .where(eq(channels.userId, userId))
    .orderBy(desc(channels.createdAt))
  return c.json({ channels: rows })
})

// POST /api/channels — register a channel
app.post('/', async (c) => {
  const { userId } = c.get('user')
  const body = await c.req.json<{
    platformHandle: string
    followerCount?: number
    categories?: string[]
    pricePerPost?: number
  }>()

  const [channel] = await db
    .insert(channels)
    .values({
      userId,
      platformHandle: body.platformHandle,
      followerCount: body.followerCount ?? 0,
      categories: body.categories ?? [],
      pricePerPost: body.pricePerPost ?? 0,
    })
    .returning()

  return c.json({ channel }, 201)
})

// POST /api/channels/:id/propose — advertiser proposes a deal directly to this channel (creates campaign + deal)
app.post('/:id/propose', async (c) => {
  const channelId = parseInt(c.req.param('id'))
  const { userId, sub } = c.get('user')
  const body = await c.req.json<{
    title: string
    contentText: string
    tokenAddress: string
    amount: string
    minViews?: number
    campaignDuration?: number
    expiryDeadline: string
    categories?: string[]
  }>()

  const [channel] = await db.select().from(channels).where(eq(channels.id, channelId)).limit(1)
  if (!channel) return c.json({ error: 'Channel not found' }, 404)

  const [campaign] = await db
    .insert(campaigns)
    .values({
      advertiserId: userId,
      title: body.title,
      contentText: body.contentText,
      tokenAddress: body.tokenAddress,
      amount: body.amount,
      minViews: body.minViews ?? 0,
      campaignDuration: body.campaignDuration ?? 0,
      expiryDeadline: new Date(body.expiryDeadline),
      categories: body.categories ?? [],
      status: 'draft',
      isDirectDeal: true, // hidden from campaign list; both sides only see the deal
    })
    .returning()

  const [deal] = await db
    .insert(deals)
    .values({
      campaignId: campaign.id,
      channelId,
      influencerId: channel.userId,
      proposedBy: sub,
      status: 'proposed',
    })
    .returning()

  return c.json({ campaign, deal }, 201)
})

// GET /api/channels/:id — get channel by id
app.get('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const [channel] = await db.select().from(channels).where(eq(channels.id, id)).limit(1)
  if (!channel) return c.json({ error: 'Channel not found' }, 404)
  return c.json({ channel })
})

// PATCH /api/channels/:id — update channel (owner only)
app.patch('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const body = await c.req.json<{
    categories?: string[]
    pricePerPost?: number
    isActive?: boolean
  }>()

  const [channel] = await db.select().from(channels).where(eq(channels.id, id)).limit(1)
  if (!channel) return c.json({ error: 'Channel not found' }, 404)
  if (channel.userId !== userId) return c.json({ error: 'Not authorized' }, 403)

  const [updated] = await db
    .update(channels)
    .set({ ...body, updatedAt: new Date() })
    .where(eq(channels.id, id))
    .returning()

  return c.json({ channel: updated })
})

// DELETE /api/channels/:id
app.delete('/:id', async (c) => {
  const id = parseInt(c.req.param('id'))
  const { userId } = c.get('user')
  const [channel] = await db.select().from(channels).where(eq(channels.id, id)).limit(1)
  if (!channel) return c.json({ error: 'Channel not found' }, 404)
  if (channel.userId !== userId) return c.json({ error: 'Not authorized' }, 403)

  await db.delete(channels).where(eq(channels.id, id))
  return c.json({ success: true })
})

export default app
