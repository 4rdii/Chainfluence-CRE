import { Hono } from 'hono'
import { generateNonce, SiweMessage } from 'siwe'
import { eq, and } from 'drizzle-orm'
import { db } from '../db'
import { users, channels } from '../db/schema'
import { signJWT, authMiddleware, type AppEnv } from '../lib/auth'

// In-memory PKCE store (TTL managed manually)
const oauthStates = new Map<string, { codeVerifier: string; userId: number; expiresAt: number }>()

const auth = new Hono<AppEnv>()

// GET /api/auth/nonce — generate a SIWE nonce
auth.get('/nonce', async (c) => {
  const nonce = generateNonce()
  return c.json({ nonce })
})

// POST /api/auth/verify — verify SIWE signature, upsert user, return JWT
auth.post('/verify', async (c) => {
  const { message, signature } = await c.req.json<{ message: string; signature: string }>()

  try {
    const siweMessage = new SiweMessage(message)
    const { data } = await siweMessage.verify({ signature })
    const walletAddress = data.address.toLowerCase()

    // Upsert user
    let [user] = await db.select().from(users).where(eq(users.walletAddress, walletAddress)).limit(1)
    if (!user) {
      ;[user] = await db
        .insert(users)
        .values({ walletAddress, role: 'both' })
        .returning()
    }

    const token = await signJWT({
      sub: walletAddress,
      userId: user.id,
      role: user.role,
    })

    return c.json({ token, user: { id: user.id, walletAddress, role: user.role, displayName: user.displayName } })
  } catch {
    return c.json({ error: 'Invalid SIWE signature' }, 401)
  }
})

// GET /api/auth/me — current user from JWT
auth.get('/me', authMiddleware, async (c) => {
  const { sub } = c.get('user')
  const [user] = await db.select().from(users).where(eq(users.walletAddress, sub)).limit(1)
  if (!user) return c.json({ error: 'User not found' }, 404)
  return c.json({ user })
})

// PATCH /api/auth/me — update display name / role
auth.patch('/me', authMiddleware, async (c) => {
  const { sub } = c.get('user')
  const body = await c.req.json<{ displayName?: string; role?: string }>()

  const [updated] = await db
    .update(users)
    .set({ ...body, updatedAt: new Date() })
    .where(eq(users.walletAddress, sub))
    .returning()

  return c.json({ user: updated })
})

// --- X OAuth 2.0 (PKCE) for channel registration ---

function generateRandomString(length: number): string {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~'
  const arr = new Uint8Array(length)
  crypto.getRandomValues(arr)
  return Array.from(arr, (b) => chars[b % chars.length]).join('')
}

async function sha256Base64url(plain: string): Promise<string> {
  const encoder = new TextEncoder()
  const data = encoder.encode(plain)
  const hash = await crypto.subtle.digest('SHA-256', data)
  const base64 = btoa(String.fromCharCode(...new Uint8Array(hash)))
  return base64.replace(/\+/g, '-').replace(/\//g, '_').replace(/=+$/, '')
}

// GET /api/auth/x/connect — start X OAuth flow (requires JWT)
auth.get('/x/connect', authMiddleware, async (c) => {
  const { userId } = c.get('user')
  const clientId = process.env.X_CLIENT_ID
  const backendUrl = process.env.BACKEND_URL || 'http://localhost:3001'

  if (!clientId) return c.json({ error: 'X OAuth not configured' }, 500)

  const state = generateRandomString(32)
  const codeVerifier = generateRandomString(64)
  const codeChallenge = await sha256Base64url(codeVerifier)

  // Store with 5 min TTL
  oauthStates.set(state, { codeVerifier, userId, expiresAt: Date.now() + 5 * 60 * 1000 })

  // Clean up expired entries
  for (const [key, val] of oauthStates) {
    if (val.expiresAt < Date.now()) oauthStates.delete(key)
  }

  const redirectUri = `${backendUrl}/api/auth/x/callback`
  const params = new URLSearchParams({
    response_type: 'code',
    client_id: clientId,
    redirect_uri: redirectUri,
    scope: 'tweet.read users.read offline.access',
    state,
    code_challenge: codeChallenge,
    code_challenge_method: 'S256',
  })

  return c.json({ url: `https://x.com/i/oauth2/authorize?${params}` })
})

// GET /api/auth/x/callback — X redirects here after user authorizes
auth.get('/x/callback', async (c) => {
  const code = c.req.query('code')
  const state = c.req.query('state')
  const error = c.req.query('error')
  const frontendUrl = process.env.FRONTEND_URL?.trim() || 'http://localhost:5173'

  if (error) {
    return c.redirect(`${frontendUrl}/my-channels?error=${encodeURIComponent(error)}`)
  }

  if (!code || !state) {
    return c.redirect(`${frontendUrl}/my-channels?error=missing_params`)
  }

  const stored = oauthStates.get(state)
  if (!stored || stored.expiresAt < Date.now()) {
    oauthStates.delete(state!)
    return c.redirect(`${frontendUrl}/my-channels?error=invalid_state`)
  }

  const { codeVerifier, userId } = stored
  oauthStates.delete(state)

  const clientId = process.env.X_CLIENT_ID!
  const clientSecret = process.env.X_CLIENT_SECRET!
  const backendUrl = process.env.BACKEND_URL || 'http://localhost:3001'
  const redirectUri = `${backendUrl}/api/auth/x/callback`

  try {
    // Exchange code for access token
    const tokenRes = await fetch('https://api.x.com/2/oauth2/token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        Authorization: `Basic ${btoa(`${clientId}:${clientSecret}`)}`,
      },
      body: new URLSearchParams({
        code,
        grant_type: 'authorization_code',
        redirect_uri: redirectUri,
        code_verifier: codeVerifier,
      }),
    })

    if (!tokenRes.ok) {
      const err = await tokenRes.text()
      console.error('X token exchange failed:', err)
      return c.redirect(`${frontendUrl}/my-channels?error=token_exchange_failed`)
    }

    const { access_token } = await tokenRes.json() as { access_token: string }

    // Fetch user profile
    const profileRes = await fetch(
      'https://api.x.com/2/users/me?user.fields=public_metrics,profile_image_url',
      { headers: { Authorization: `Bearer ${access_token}` } }
    )

    if (!profileRes.ok) {
      console.error('X profile fetch failed:', await profileRes.text())
      return c.redirect(`${frontendUrl}/my-channels?error=profile_fetch_failed`)
    }

    const { data: xUser } = await profileRes.json() as {
      data: {
        id: string
        username: string
        name: string
        profile_image_url?: string
        public_metrics?: { followers_count: number }
      }
    }

    // Upsert channel: if this X user already connected, update; otherwise insert
    const existing = await db
      .select()
      .from(channels)
      .where(and(eq(channels.platformUserId, xUser.id), eq(channels.userId, userId)))
      .limit(1)

    if (existing.length > 0) {
      await db
        .update(channels)
        .set({
          platformHandle: xUser.username,
          followerCount: xUser.public_metrics?.followers_count ?? 0,
          profileImageUrl: xUser.profile_image_url,
          verifiedAt: new Date(),
          updatedAt: new Date(),
        })
        .where(eq(channels.id, existing[0].id))
    } else {
      await db.insert(channels).values({
        userId,
        platform: 'twitter',
        platformHandle: xUser.username,
        platformUserId: xUser.id,
        followerCount: xUser.public_metrics?.followers_count ?? 0,
        profileImageUrl: xUser.profile_image_url,
        verifiedAt: new Date(),
      })
    }

    return c.redirect(`${frontendUrl}/my-channels?connected=true`)
  } catch (e: any) {
    console.error('X OAuth callback error:', e)
    return c.redirect(`${frontendUrl}/my-channels?error=callback_failed`)
  }
})

export default auth
