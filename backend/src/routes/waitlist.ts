import { Hono } from 'hono'
import { db } from '../db'
import { waitlist } from '../db/schema'

const app = new Hono()

app.post('/', async (c) => {
  const { email, role } = await c.req.json<{ email: string; role?: string }>()

  if (!email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
    return c.json({ error: 'Valid email is required' }, 400)
  }

  try {
    await db.insert(waitlist).values({
      email: email.toLowerCase().trim(),
      role: role || 'unknown',
    })
  } catch (err: any) {
    // Duplicate email — treat as success (don't leak info)
    if (err.code === '23505') {
      return c.json({ success: true })
    }
    throw err
  }

  return c.json({ success: true })
})

export default app
