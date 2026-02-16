# Chainfluence — Figma AI Design Prompt (Part 1)

Design a dark-themed web3 app called "Chainfluence" — a decentralized Twitter/X advertising marketplace. Advertisers create campaigns, influencers apply, payments are held in smart contract escrow and released when Chainlink oracle verification confirms the tweet meets criteria.

## Branding & Colors

**Name**: Chainfluence. **Tagline**: "Trustless influencer marketing, powered by Chainlink". **Logo**: shield/vault + chain-link motif, geometric, minimal. **Fonts**: Space Grotesk (headings), Inter (body), JetBrains Mono (addresses/hashes).

**Dark theme**: BG #0A0A0F, cards #12121A, elevated #1A1A2E, borders #2A2A3E. **Accent**: violet #7C3AED, gradient CTA: linear-gradient(135deg, #7C3AED, #3B82F6). **Status colors**: success #10B981, warning #F59E0B, error #EF4444. **Text**: primary #F8FAFC, secondary #94A3B8, muted #64748B.

## Design System

Cards: #12121A, 1px border, 12px radius, 24px padding, subtle glassmorphism. Hover: lift -2px + purple border glow. Buttons primary: gradient bg, white text, 8px radius, 44px height. Buttons secondary: transparent, 1px #3A3A4E border. Inputs: #12121A, 1px border, 44px height. Focus: purple border glow. Badges: colored pills (green=completed, purple=verifying, blue=funded, amber=posted, gray=proposed, red=disputed).

## Navigation

**Header** (fixed 64px): Logo left. Right: "Sepolia" network pill, wallet address pill with Jazzicon avatar, notification bell.

**Sidebar** (240px, collapsible <1024px): Dashboard, Marketplace, My Campaigns, My Channels, My Deals, separator, Settings, Docs. Active: purple left border + #1A1A2E bg.

## Screen 1: Connect Wallet

Full viewport, no shell. Subtle dark gradient mesh background. Centered 480px card:
- Logo 48px centered
- H1: "Welcome to Chainfluence"
- Subtitle: "Connect your wallet to start creating or completing ad campaigns with trustless escrow."
- Full-width "Connect Wallet" gradient button (52px)
- Three feature pills: "Trustless Escrow" | "Twitter Verified" | "Chainlink Powered"

**After connect**: checkmark + address shown, "Sign message" button for SIWE verification.
**After sign**: "Choose your role" — three selectable cards: Advertiser (bullhorn icon), Influencer (megaphone icon), Both (handshake icon). Selected = purple border. Optional display name input. "Get Started" button.

## Screen 2: Dashboard

Full app shell. **Stats row**: 4 equal cards — Active Deals (count), Pending Verifications (amber tint), Total Spent/Earned (green), Success Rate (purple progress ring).

**Two columns (2:1)**: Left: "Active Deals" list — compact horizontal deal cards with Twitter icon, campaign title, counterparty address, status badge pill, next action hint text, deadline countdown in amber. Right: "Recent Activity" — vertical timeline with colored dots (green=success, blue=proposal, purple=verification, amber=post) and timestamped event descriptions.

**Quick Actions row**: "Create Campaign" (gradient) + "Browse Marketplace" (outline).

## Screen 3: Marketplace — Campaigns Tab

Tabs: "Campaigns" (active, purple underline) | "Influencers". Search bar (48px, search icon) + filter row: category multi-select, budget range min/max, min views input, sort dropdown (Newest/Highest Budget/Ending Soon).

**Campaign cards** (3-col grid, 16px gap): Twitter icon top-left in sky-blue circle, deadline badge top-right (green >7d, amber 1-7d, red <1d). Title bold 16px white (2 lines max). Description muted 14px (2 lines). Category tag pills (#1A1A2E bg). Thin separator. Three metrics row: Budget "500 USDC", Min Views "10K", Duration "7 days" — values bold white, labels 12px muted below. Buttons: "View Details" (secondary) + "Apply" (gradient). Card hover: translateY(-2px) + faint purple shadow.

## Screen 4: Marketplace — Influencers Tab

Same layout, different cards: 40px Jazzicon avatar + @handle in purple + display name bold + Twitter verified badge. Category tags. Bio 14px muted (2 lines). Separator. Metrics: Followers "125K", Rate "$50/post", Rating "4.8". Buttons: "View Profile" (secondary) + "Propose Deal" (gradient).

Filters adapt: follower range, price range, category instead of budget/views.
