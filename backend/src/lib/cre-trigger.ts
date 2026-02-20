import { createHash } from 'node:crypto'
import { randomUUID } from 'node:crypto'
import { privateKeyToAccount, type Hex } from 'viem/accounts'
import { parseSignature } from 'viem'
import stringify from 'json-stable-stringify'

const CRE_GATEWAY_URL =
  process.env.CRE_GATEWAY_URL || 'https://01.gateway.zone-a.cre.chain.link'
const CRE_TRIGGER_PRIVATE_KEY = process.env.CRE_TRIGGER_PRIVATE_KEY as string
const CRE_WORKFLOW_ID = (process.env.CRE_TWITTER_WORKFLOW_ID || '').replace(/^0x/, '')

interface TriggerPayload {
  campaignId: string
  tweetUrl: string
  chainName?: string
}

function base64URLEncode(str: string): string {
  return str.replace(/\+/g, '-').replace(/\//g, '_').replace(/=/g, '')
}

function sha256Hex(data: object): string {
  const jsonString = stringify(data)
  return createHash('sha256').update(jsonString, 'utf8').digest('hex')
}

/** Create CRE JWT: header.payload.signature (ETH-signed message). */
async function createCREJWT(
  request: object,
  privateKey: Hex
): Promise<string> {
  const account = privateKeyToAccount(privateKey as Hex)

  const header = { alg: 'ETH', typ: 'JWT' }
  const now = Math.floor(Date.now() / 1000)
  const payload = {
    digest: `0x${sha256Hex(request)}`,
    iss: account.address,
    iat: now,
    exp: now + 300,
    jti: randomUUID(),
  }

  const encodedHeader = base64URLEncode(
    Buffer.from(JSON.stringify(header), 'utf8').toString('base64')
  )
  const encodedPayload = base64URLEncode(
    Buffer.from(JSON.stringify(payload), 'utf8').toString('base64')
  )
  const message = `${encodedHeader}.${encodedPayload}`

  const signature = await account.signMessage({ message })
  const { r, s, v } = parseSignature(signature)

  const signatureBytes = Buffer.concat([
    Buffer.from(r.slice(2), 'hex'),
    Buffer.from(s.slice(2), 'hex'),
    Buffer.from([Number(v)]),
  ])
  const encodedSignature = base64URLEncode(signatureBytes.toString('base64'))

  return `${encodedHeader}.${encodedPayload}.${encodedSignature}`
}

export async function triggerCREWorkflow(
  payload: TriggerPayload
): Promise<{ success: boolean; error?: string }> {
  if (!CRE_TRIGGER_PRIVATE_KEY?.trim()) {
    return { success: false, error: 'CRE trigger private key not configured' }
  }
  if (!CRE_WORKFLOW_ID) {
    return { success: false, error: 'CRE workflow ID not configured' }
  }

  const privateKey = CRE_TRIGGER_PRIVATE_KEY.startsWith('0x')
    ? (CRE_TRIGGER_PRIVATE_KEY as Hex)
    : (`0x${CRE_TRIGGER_PRIVATE_KEY}` as Hex)

  const request = {
    jsonrpc: '2.0' as const,
    id: randomUUID(),
    method: 'workflows.execute' as const,
    params: {
      input: {
        campaignId: payload.campaignId,
        tweetUrl: payload.tweetUrl,
      },
      workflow: {
        workflowID: CRE_WORKFLOW_ID,
      },
    },
  }

  const body = stringify(request)

  let jwt: string
  try {
    jwt = await createCREJWT(request, privateKey)
  } catch (err) {
    const msg = err instanceof Error ? err.message : String(err)
    return { success: false, error: `CRE JWT failed: ${msg}` }
  }

  try {
    const response = await fetch(CRE_GATEWAY_URL, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${jwt}`,
      },
      body,
    })

    const text = await response.text()
    let data: { error?: { code?: number; message?: string }; result?: unknown }
    try {
      data = JSON.parse(text)
    } catch {
      if (!response.ok) {
        return {
          success: false,
          error: `CRE gateway error: ${response.status} ${text.slice(0, 200)}`,
        }
      }
      return { success: false, error: `CRE gateway invalid JSON: ${text.slice(0, 100)}` }
    }

    if (data.error) {
      return {
        success: false,
        error: data.error.message || JSON.stringify(data.error),
      }
    }

    if (!response.ok) {
      return {
        success: false,
        error: `CRE gateway error: ${response.status} ${text.slice(0, 200)}`,
      }
    }

    return { success: true }
  } catch (err) {
    const msg = err instanceof Error ? err.message : String(err)
    return { success: false, error: `CRE trigger failed: ${msg}` }
  }
}
