const API_BASE = import.meta.env.VITE_API_URL || 'http://localhost:3001'

export async function apiClient<T = any>(
  path: string,
  options?: RequestInit
): Promise<T> {
  const token = localStorage.getItem('jwt')
  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...options?.headers,
    },
  })

  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: res.statusText }))
    throw new Error(body.error || `API error: ${res.status}`)
  }

  return res.json()
}

export function setAuthToken(token: string) {
  localStorage.setItem('jwt', token)
}

export function clearAuthToken() {
  localStorage.removeItem('jwt')
}
