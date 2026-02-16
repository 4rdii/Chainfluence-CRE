# Chainfluence — Figma AI Design Prompt

Design a dark-themed web3 app called "Chainfluence" — a decentralized Twitter/X advertising marketplace. Advertisers create campaigns, influencers apply, payments are held in smart contract escrow and released when Chainlink oracle verification confirms the tweet meets criteria.

## Branding & Colors

**Name**: Chainfluence. **Tagline**: "Trustless influencer marketing, powered by Chainlink". **Logo**: shield/vault + chain-link motif, geometric, minimal. **Fonts**: Space Grotesk (headings), Inter (body), JetBrains Mono (addresses/hashes).

**Dark theme**: BG #0A0A0F, cards #12121A, elevated #1A1A2E, borders #2A2A3E. **Accent**: violet #7C3AED, gradient CTA: linear-gradient(135deg, #7C3AED, #3B82F6). **Status colors**: success #10B981, warning #F59E0B, error #EF4444. **Text**: primary #F8FAFC, secondary #94A3B8, muted #64748B.

## Design System

Cards: #12121A, 1px border, 12px radius, 24px padding, subtle glassmorphism. Hover: lift -2px + purple border glow. Buttons primary: gradient bg, white text, 8px radius, 44px height. Buttons secondary: transparent, 1px #3A3A4E border. Inputs: #12121A, 1px border, 44px height. Focus: purple border glow. Badges: colored pills (green=completed, purple=verifying, blue=funded, amber=posted, gray=proposed, red=disputed).

## Navigation

**Header** (fixed 64px): Logo left. Right: "Sepolia" network pill, wallet address pill with Jazzicon, notification bell.

**Sidebar** (240px, collapsible <1024px): Dashboard, Marketplace, My Campaigns, My Channels, My Deals, separator, Settings, Docs. Active: purple left border + #1A1A2E bg.

## Screen 1: Connect Wallet

Full viewport, no shell. Subtle dark gradient mesh background. Centered 480px card:
- Logo 48px centered
- H1: "Welcome to Chainfluence"
- Subtitle: "Connect your wallet to start..."
- Full-width "Connect Wallet" gradient button (52px)
- Three feature pills: "Trustless Escrow" | "Twitter Verified" | "Chainlink Powered"

**After connect**: checkmark + address shown, "Sign message" button for SIWE.
**After sign**: "Choose your role" — three selectable cards: Advertiser (bullhorn), Influencer (megaphone), Both (handshake). Selected = purple border. Optional display name input. "Get Started" button.

## Screen 2: Dashboard

Full app shell. **Stats row**: 4 equal cards — Active Deals (count), Pending Verifications (amber), Total Spent/Earned (green), Success Rate (purple progress ring).

**Two columns (2:1)**: Left: "Active Deals" — compact horizontal deal cards with Twitter icon, title, counterparty address, status badge, next action hint, deadline countdown. Right: "Recent Activity" — vertical timeline with colored dots and timestamped events.

**Quick Actions**: "Create Campaign" (gradient) + "Browse Marketplace" (outline).

## Screen 3: Marketplace — Campaigns

Tabs: "Campaigns" (active) | "Influencers". Search bar + filter row (category, budget range, min views, sort dropdown).

**Campaign cards** (3-col grid): Twitter icon top-left, deadline badge top-right (color-coded by urgency). Title (bold, 2 lines max). Description (muted, 2 lines). Category tag pills. Separator. Three metrics: Budget "500 USDC", Min Views "10K", Duration "7 days" with labels below. Two buttons: "View Details" (secondary) + "Apply" (gradient). Hover: lift + purple shadow.

## Screen 4: Marketplace — Influencers

Same layout, different cards: Avatar circle + @handle (purple) + display name + verified badge. Category tags. Bio (2 lines). Separator. Metrics: Followers, Rate ($/post), Rating. Buttons: "View Profile" + "Propose Deal".

## Screen 5: Create Campaign (Multi-step)

640px centered. **Horizontal stepper**: (1) Platform → (2) Content → (3) Requirements → (4) Payment → (5) Confirm. Green checkmarks for completed, purple active, gray upcoming.

**Step 1**: Single selectable Twitter/X card (selected = purple border + checkmark).
**Step 2**: Title input, description textarea, content text (mono font, char count), categories multi-select chips. Preview card below.
**Step 3**: Min views input, duration dropdown (1d/3d/7d/14d/30d/custom), expiry date picker. Blue info box about duration requirements.
**Step 4**: Token dropdown, large amount input. Breakdown card: influencer receives / platform fee 2.5% / total. Wallet balance shown.
**Step 5**: Full summary in 2-col grid. Amber warning about gas fees. "Deposit & Create" large gradient button. **Pending state**: spinner + tx hash link. **Success state**: green checkmark animation + campaign ID + "View Campaign" button.

## Screen 6: Deal View

**Status tracker** (full width, 6 steps): Proposed → Accepted → Funded → Posted → Verifying → Completed. Green checkmarks, purple active pulse, gray upcoming. Red X for disputed.

**Left column (2/3)**:
- Campaign details card (title, content in code block, metrics, addresses)
- **Action panel** (changes per status): proposed = Accept/Decline buttons; accepted = "Deposit" button; funded = tweet URL input + Submit; posted = "Verify Post" gradient button with Chainlink icon; verifying = spinner + progress; completed = green success + tx link; disputed = red alert
- Verification history table: #, Date, Views (actual/required), Content Match, Duration, Edited, Action, Status badges

**Right column (1/3)**:
- Deal info card: status badge, amount, dates, advertiser/influencer addresses with Jazzicons, on-chain ID + tx hashes as Etherscan links
- Activity timeline: vertical line with colored dots + timestamped events

## Screen 7: My Campaigns

Tabs: Active | Completed | Draft | All. "Create Campaign" button top-right. Table: Title, Budget, Deals, Status badge, Deadline, View link. **Empty state**: megaphone illustration + "No campaigns yet" + CTA.

## Screen 8: My Channels

"Register Channel" button top-right. Full-width horizontal cards: Twitter icon + @handle + follower count + price/post + categories + active/inactive toggle + edit/delete. **Empty state**: broadcast tower illustration + "Add your Twitter channels" + CTA.

## Global States

**Loading**: skeleton shimmer loaders matching card/table shapes. **Empty**: line-art illustration (purple tint) + text + CTA. **Error**: red-bordered alert, red input borders. **Success**: green alerts, checkmark animations. **Toasts** (top-right overlay, 360px): colored left border (green/blue/amber/red), auto-dismiss 5s, stackable.

## Responsive

Desktop >1280px: sidebar + 3-col grid + 2-col deal. Tablet 768-1279px: hamburger + 2-col grid + 1-col deal. Mobile <768px: overlay nav + 1-col everything. Stats: 4-col → 2x2 → stack. Tables: scroll horizontal. Stepper: horizontal → vertical. Forms: always single-col.
