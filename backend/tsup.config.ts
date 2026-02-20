import { defineConfig } from 'tsup'

export default defineConfig({
  entry: ['src/index.ts'],
  format: ['cjs'],
  outDir: 'api',
  clean: true,
  splitting: false,
  external: [
    // CJS packages that break when bundled into ESM
    'siwe',
    '@spruceid/siwe-parser',
    'apg-js',
    // Only used locally
    '@hono/node-server',
  ],
  noExternal: [/.*/],  // Bundle all other npm packages
})
