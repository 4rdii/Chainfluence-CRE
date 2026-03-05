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

// POST /api/channels/:id/verify-worldid — verify channel owner is a unique human via World ID
app.post('/:id/verify-worldid', async (c) => {
  const channelId = parseInt(c.req.param('id'))
  const { userId } = c.get('user')

  // Verify caller owns the channel
  const [channel] = await db.select().from(channels).where(eq(channels.id, channelId)).limit(1)
  if (!channel) return c.json({ error: 'Channel not found' }, 404)
  if (channel.userId !== userId) return c.json({ error: 'Not authorized' }, 403)
  if (channel.worldIdNullifier) return c.json({ error: 'Channel already verified with World ID' }, 400)

  const body = await c.req.json<{
    merkle_root: string
    nullifier_hash: string
    proof: string
    verification_level: string
  }>()

  // Verify the proof with World ID API
  const appId = process.env.WORLD_ID_APP_ID
  const actionId = process.env.WORLD_ID_ACTION_ID || 'verify-channel'

  if (!appId) return c.json({ error: 'World ID not configured on server' }, 500)

  const verifyRes = await fetch(
    `https://developer.worldcoin.org/api/v2/verify/${appId}`,
    {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        merkle_root: body.merkle_root,
        nullifier_hash: body.nullifier_hash,
        proof: body.proof,
        action: actionId,
        verification_level: body.verification_level,
      }),
    }
  )

  if (!verifyRes.ok) {
    const errData = await verifyRes.json().catch(() => ({}))
    return c.json({ error: 'World ID verification failed', detail: (errData as { detail?: string }).detail || verifyRes.statusText }, 400)
  }

  // Check uniqueness — same nullifier can't verify two different channels
  const [existing] = await db
    .select()
    .from(channels)
    .where(eq(channels.worldIdNullifier, body.nullifier_hash))
    .limit(1)
  if (existing) return c.json({ error: 'This World ID has already been used to verify another channel' }, 409)

  // Mark channel as verified
  const [updated] = await db
    .update(channels)
    .set({
      worldIdNullifier: body.nullifier_hash,
      worldIdVerifiedAt: new Date(),
      updatedAt: new Date(),
    })
    .where(eq(channels.id, channelId))
    .returning()

  return c.json({ channel: updated })
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
