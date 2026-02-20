import { Hono } from 'hono'
import { createHmac } from 'node:crypto'
import { eq } from 'drizzle-orm'
import { getEventSelector } from 'viem'
import { db } from '../db'
import { chainEvents, deals } from '../db/schema'

const ALCHEMY_WEBHOOK_SIGNING_KEY = process.env.ALCHEMY_WEBHOOK_SIGNING_KEY || ''

const app = new Hono()

// POST /api/webhook/chain-events — Alchemy webhook receiver
app.post('/chain-events', async (c) => {
  // Verify Alchemy HMAC signature
  if (ALCHEMY_WEBHOOK_SIGNING_KEY) {
    const signature = c.req.header('x-alchemy-signature')
    const rawBody = await c.req.text()

    const hmac = createHmac('sha256', ALCHEMY_WEBHOOK_SIGNING_KEY)
    hmac.update(rawBody)
    const expectedSig = hmac.digest('hex')

    if (signature !== expectedSig) {
      return c.json({ error: 'Invalid webhook signature' }, 401)
    }

    // Re-parse body since we consumed it
    const body = JSON.parse(rawBody)
    return processWebhookEvent(c, body)
  }

  const body = await c.req.json()
  return processWebhookEvent(c, body)
})

async function processWebhookEvent(c: any, body: any) {
  const logs = body.event?.data?.block?.logs || []

  for (const log of logs) {
    const eventName = parseEventName(log.topics?.[0])
    const txHash = log.transaction?.hash || ''

    // Store raw event
    await db.insert(chainEvents).values({
      eventName,
      txHash,
      blockNumber: BigInt(log.blockNumber || 0),
      rawData: log,
      processed: false,
    })

    // Process known events (V2 contract uses dealId in events; we store it in onchainCampaignId)
    switch (eventName) {
      case 'DealCreated':
      case 'CampaignFunded': {
        const dealId = parseInt(log.topics?.[1], 16)
        await db
          .update(deals)
          .set({ status: 'funded', fundedAt: new Date(), updatedAt: new Date() })
          .where(eq(deals.onchainCampaignId, dealId))
        break
      }
      case 'FundsReleased': {
        const dealId = parseInt(log.topics?.[1], 16)
        await db
          .update(deals)
          .set({ status: 'completed', completedAt: new Date(), updatedAt: new Date() })
          .where(eq(deals.onchainCampaignId, dealId))
        break
      }
      case 'FundsRefunded': {
        const dealId = parseInt(log.topics?.[1], 16)
        await db
          .update(deals)
          .set({ status: 'refunded', updatedAt: new Date() })
          .where(eq(deals.onchainCampaignId, dealId))
        break
      }
    }

    // Mark as processed
    await db
      .update(chainEvents)
      .set({ processed: true })
      .where(eq(chainEvents.txHash, txHash))
  }

  return c.json({ processed: logs.length })
}

// Map known event topic hashes (topic[0]) to names — AdEscrowV2 events
const EVENT_TOPICS: Record<string, string> = {
  [getEventSelector('DealCreated(uint256,address,address,address,uint256,bytes32)')]: 'DealCreated',
  [getEventSelector('FundsReleased(uint256,address,address,uint256,uint256)')]: 'FundsReleased',
  [getEventSelector('FundsRefunded(uint256,address,address,uint256)')]: 'FundsRefunded',
  // Legacy V1 name for funded
  [getEventSelector('CampaignFunded(uint256,address,address,uint256)')]: 'CampaignFunded',
}

function parseEventName(topic: string): string {
  if (!topic) return 'Unknown'
  const key = topic.startsWith('0x') ? topic : `0x${topic}`
  return EVENT_TOPICS[key] ?? EVENT_TOPICS[key.toLowerCase()] ?? 'Unknown'
}

export default app
