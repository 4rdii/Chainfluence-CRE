# Design System

## Branding

**Name:** Chainfluence
**Tagline:** "Trustless influencer marketing, powered by Chainlink"
**Fonts:** Space Grotesk (headings), Inter (body), JetBrains Mono (addresses/hashes/code)

## Colors

```
Background:   #0A0A0F
Cards:        #12121A
Elevated:     #1A1A2E
Borders:      #2A2A3E

Accent violet:  #7C3AED
CTA gradient:   linear-gradient(135deg, #7C3AED, #3B82F6)

Text primary:   #F8FAFC
Text secondary: #94A3B8
Text muted:     #64748B

Success:  #10B981
Warning:  #F59E0B
Error:    #EF4444
```

In Tailwind these map to custom classes defined in `frontend/src/styles/theme.css`:
`bg-background`, `bg-card`, `bg-elevated`, `border-border-color`, `text-text-primary`, `text-text-secondary`, `text-text-muted`, `text-accent-violet`, `text-accent-blue`.

## Component Patterns

**Cards:** `bg-card border border-border-color rounded-xl p-6`. Hover: `translateY(-2px)` + faint violet shadow.

**Buttons primary:** gradient background, white text, 8px radius, 44px height.

**Buttons secondary:** transparent, `border border-border-color`.

**Inputs:** `bg-card border border-border-color`, 44px height, focus: violet border glow.

**Status badges (deal states):**

| State | Color |
|-------|-------|
| proposed | blue |
| accepted | yellow |
| funded | violet |
| posted | cyan |
| verifying | orange |
| completed | green |
| refunded | red |
| disputed | gray |

## Layout

**Header** (fixed 64px): Logo left. Right: network pill ("Sepolia"), wallet address with Jazzicon avatar, notification bell.

**Sidebar** (240px, collapsible below 1024px): Dashboard, Marketplace, My Campaigns, My Channels, My Deals, Settings, Docs. Active item: violet left border + elevated background.

**Responsive breakpoints:**
- Desktop `>1280px`: sidebar + 3-col grid + 2-col deal view
- Tablet `768–1279px`: hamburger + 2-col grid + 1-col deal
- Mobile `<768px`: overlay nav + 1-col everything

## Key Screens

### Connect Wallet
Full viewport (no shell). Centered 480px card: logo, headline, connect button, three feature pills ("Trustless Escrow", "Twitter Verified", "Chainlink Powered"). After connect: SIWE sign prompt. After sign: role selection (Advertiser / Influencer / Both).

### Marketplace
Tabs: Campaigns | Influencers. Search + filters. Campaign cards show: budget, min views, duration, deadline badge (green >7d, amber 1–7d, red <1d). Influencer cards show: avatar, handle, followers, rate, categories.

### Create Campaign
Centered 640px, multi-step: Platform → Content → Requirements → Payment → Confirm. No wallet transaction at creation — funding happens later on the deal view.

### Deal View
Status tracker (6 steps: Proposed → Accepted → Funded → Posted → Verifying → Completed). 2/3 + 1/3 column split. Left: campaign details + role-specific action panel. Right: deal metadata + activity timeline.

### Action Panel (changes per deal state)
- **proposed** (influencer): Accept / Decline buttons
- **accepted** (advertiser): Fund Deal button → approve + deposit transactions
- **funded** (influencer): Tweet URL input + Submit
- **verifying**: spinner, polling every 15s, Refresh Status button
- **completed**: green success + tx hash
- **disputed**: red alert

## Global States

- **Loading:** skeleton shimmer matching card/table shapes
- **Empty:** line-art illustration (violet tint) + description + CTA per section
- **Error:** red-bordered alert, red input borders
- **Toasts:** top-right overlay, colored left border, auto-dismiss 5s, stackable
