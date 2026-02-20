import { Hono } from 'hono'
import { cors } from 'hono/cors'
import { logger } from 'hono/logger'

import auth from './routes/auth'
import campaigns from './routes/campaigns'
import deals from './routes/deals'
import channels from './routes/channels'
import marketplace from './routes/marketplace'
import webhook from './routes/webhook'

const app = new Hono().basePath('/api')

// Middleware
app.use('*', logger())
app.use(
  '*',
  cors({
    origin: (origin) => {
      const allowed = [
        'http://localhost:5173',
        'http://localhost:3000',
        process.env.FRONTEND_URL?.trim(),
      ].filter(Boolean) as string[]
      return allowed.includes(origin) ? origin : allowed[0]
    },
    credentials: true,
  })
)

// Global error handler — surfaces errors instead of generic 500
app.onError((err, c) => {
  console.error('Unhandled error:', err)
  return c.json({ error: err.message, stack: err.stack?.split('\n').slice(0, 3) }, 500)
})

// Health check
app.get('/health', (c) => c.json({ status: 'ok', timestamp: Date.now() }))

// Routes
app.route('/auth', auth)
app.route('/campaigns', campaigns)
app.route('/deals', deals)
app.route('/channels', channels)
app.route('/marketplace', marketplace)
app.route('/webhook', webhook)

// Local dev server (VERCEL env var is always set on Vercel)
if (!process.env.VERCEL) {
  import('@hono/node-server').then(({ serve }) => {
    const port = parseInt(process.env.PORT || '3001')
    serve({ fetch: app.fetch, port })
    console.log(`Server running on http://localhost:${port}`)
  })
}

// Vercel serverless handler (Node.js CJS format)
import type { IncomingMessage, ServerResponse } from 'http'

async function vercelHandler(req: IncomingMessage, res: ServerResponse) {
  try {
    // Convert Node.js request to Web API Request
    const protocol = req.headers['x-forwarded-proto'] || 'https'
    const host = req.headers['x-forwarded-host'] || req.headers.host || 'localhost'
    const url = `${protocol}://${host}${req.url}`

    const headers = new Headers()
    for (const [key, value] of Object.entries(req.headers)) {
      if (value) headers.set(key, Array.isArray(value) ? value.join(', ') : value)
    }

    const body = ['GET', 'HEAD'].includes(req.method || 'GET')
      ? undefined
      : await new Promise<Buffer>((resolve) => {
          const chunks: Buffer[] = []
          req.on('data', (chunk: Buffer) => chunks.push(chunk))
          req.on('end', () => resolve(Buffer.concat(chunks)))
        })

    const request = new Request(url, {
      method: req.method,
      headers,
      body,
    })

    const response = await app.fetch(request)

    // Convert Web API Response to Node.js response
    res.statusCode = response.status
    response.headers.forEach((value, key) => {
      res.setHeader(key, value)
    })

    const responseBody = await response.text()
    res.end(responseBody)
  } catch (e: any) {
    res.statusCode = 500
    res.setHeader('Content-Type', 'application/json')
    res.end(JSON.stringify({ error: e.message, stack: e.stack }))
  }
}

export default vercelHandler

