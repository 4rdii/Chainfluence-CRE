# Chainfluence — Figma AI Design Prompt (Part 2)

## Screen 5: Create Campaign (Multi-step Form)

640px max-width, centered. **Horizontal stepper** at top: (1) Platform → (2) Content → (3) Requirements → (4) Payment → (5) Confirm. Completed = green circle checkmark + green line. Active = purple pulsing circle + bold label. Upcoming = gray circle + muted label.

**Step 1 Platform**: H2 "Choose Platform". Single large selectable Twitter/X card with logo, name, description "Verify tweets, view counts, posting duration, and edit detection". Selected state: 2px purple border + checkmark. "Next" button.

**Step 2 Content**: Title text input. Description textarea (optional). Content Text textarea (mono font JetBrains Mono, #0A0A0F bg, char count "124/280"). Categories multi-select with removable pill chips. Live preview card showing marketplace appearance. Back/Next buttons.

**Step 3 Requirements**: Min Views number input (helper: "0 = no requirement"). Duration dropdown (No requirement/1d/3d/7d/14d/30d/Custom). Expiry deadline date+time picker. Blue info box: "The influencer must keep the tweet posted for the full campaign duration. Chainlink verification checks it hasn't been deleted or edited." Back/Next buttons.

**Step 4 Payment**: Token selector dropdown (icon + name + symbol). Large amount input (24px font, token symbol inside). Breakdown card: "Influencer receives: 487.50 USDC / Platform fee (2.5%): 12.50 USDC / Total deposit: 500.00 USDC". Wallet balance shown; red "Insufficient balance" if needed. Back/"Review & Deposit" buttons.

**Step 5 Confirm**: Summary in 2-col label:value grid (platform, title, content in mono, views, duration, deadline, token, amount, fee). Amber warning box about blockchain transaction + gas fees. "Deposit & Create Campaign" large full-width gradient button with lock icon. **Pending**: spinner + tx hash link to Etherscan. **Success**: green checkmark animation + "Campaign Created!" + campaign ID + "View Campaign" button.

## Screen 6: Deal View

**Status tracker** (full width): 6-step horizontal stepper: Proposed → Accepted → Funded → Posted → Verifying → Completed. Same style as campaign stepper. Disputed state = red X icon.

**Left column (2/3)**:
Campaign details card: title + Twitter icon, content text in code block (#0A0A0F bg, mono, scrollable), metrics (views, duration, deadline, amount), advertiser/influencer addresses truncated mono with copy buttons.

**Action panel** (card with colored left border, changes per status):
- Proposed: green border, "New proposal" header, Accept (green gradient) / Decline (red outline) buttons
- Accepted: blue border, "Fund this campaign", "Deposit 500 USDC" gradient button
- Funded: purple border, "Submit your tweet", URL input placeholder "https://x.com/...", Submit button
- Posted: purple border, "Ready for verification", "Verify Post" large gradient button with Chainlink icon, helper text listing checks
- Verifying: purple border pulsing, spinner + "Verification in progress... usually 1-2 minutes"
- Completed: green border, checkmark + "Deal Completed!", tx hash Etherscan link
- Disputed: red border, "Deal Disputed", reason text, "Under review" message

Verification history table: columns #, Date, Views (actual/required, green if met red if not), Content (Match/Mismatch badge), Duration (actual/required), Edited (Yes red/No green), Action badge, Status badge.

**Right column (1/3)**:
Deal info card: large status badge, amount "500 USDC" in H2, created date, separator, advertiser section (Jazzicon + address + name), influencer section (same), separator, on-chain ID with Etherscan link, deposit TX hash, release TX hash (if completed).

Activity timeline card: vertical line left-aligned, colored dots + timestamps + descriptions.

## Screen 7: My Campaigns

Tabs: Active | Completed | Draft | All. "Create Campaign" gradient button top-right. Full-width table: Title, Budget, Deals count, Status badge, Deadline, "View" link. Alternating row shading, hover highlight. **Empty state**: megaphone line-art illustration (purple tint, 120px) + "No campaigns yet" H3 + "Create your first campaign" muted text + CTA button.

## Screen 8: My Channels

"Register Channel" gradient button top-right. Vertical stack of full-width horizontal cards: Twitter icon circle + @handle bold + display name muted, follower count, price/post, category tags, active/inactive toggle, edit/delete links. **Empty state**: broadcast tower illustration + "No channels registered" + "Add your Twitter channels" + CTA.

## Global States & Responsive

**Loading**: skeleton shimmer loaders matching card/table shapes (#1A1A2E to #2A2A3E pulse animation). **Empty**: unique line-art illustration per section (purple tint) + description + CTA. **Error**: red left-border alert box, red input borders + error text. **Success**: green alerts, checkmark animations. **Toasts** (top-right, 360px cards): colored left border per type, auto-dismiss 5s, stackable.

**Responsive**: Desktop >1280px: sidebar + 3-col grid + 2-col deal view. Tablet 768-1279px: hamburger menu + 2-col grid + 1-col deal. Mobile <768px: overlay nav + 1-col everything. Stats: 4→2x2→stack. Tables: horizontal scroll. Stepper: horizontal→vertical on mobile. Forms: always single-col full-width.
