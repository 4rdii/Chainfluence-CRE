import { v4 as uuidv4 } from 'uuid'
import { type Hex, parseSignature } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { base64URLEncode, sha256 } from './utils'

export interface JSONRPCRequest {
	jsonrpc: string
	id: string
	method: string
	params: {
		input: any
		workflow: {
			workflowID: string
		}
	}
}

export const createJWT = async (request: JSONRPCRequest, privateKey: Hex): Promise<string> => {
	const account = privateKeyToAccount(privateKey)
	const address = account.address

	// Create JWT header
	const header = {
		alg: 'ETH',
		typ: 'JWT',
	}

	// Create JWT payload with request and metadata
	const now = Math.floor(Date.now() / 1000)

	const payload = {
		digest: `0x${sha256(request)}`,
		iss: address,
		iat: now,
		exp: now + 300, // 5 minutes expiration
		jti: uuidv4(),
	}

	// Encode header and payload to base64url
	const encodedHeader = base64URLEncode(
		Buffer.from(JSON.stringify(header), 'utf8').toString('base64'),
	)
	const encodedPayload = base64URLEncode(
		Buffer.from(JSON.stringify(payload), 'utf8').toString('base64'),
	)

	// Create message to sign: base64url(header) + "." + base64url(payload)
	const message = `${encodedHeader}.${encodedPayload}`

	// Sign message using Ethereum signed message format
	const signature = await account.signMessage({
		message,
	})

	// Parse signature to get r, s, v
	const { r, s, v } = parseSignature(signature)

	// Concatenate r || s || v (65 bytes total)
	const signatureBytes = Buffer.concat([
		Buffer.from(r.slice(2), 'hex'), // 32 bytes
		Buffer.from(s.slice(2), 'hex'), // 32 bytes
		Buffer.from([Number(v)]), // 1 byte - convert BigInt to number
	])

	// Encode signature to base64url
	const encodedSignature = base64URLEncode(signatureBytes.toString('base64'))

	// Combine JWT: header.payload.signature
	return `${encodedHeader}.${encodedPayload}.${encodedSignature}`
}

