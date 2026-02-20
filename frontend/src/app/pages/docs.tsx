import { useState } from 'react';
import {
  Book,
  Megaphone,
  Radio,
  Code,
  ChevronDown,
  ChevronRight,
  Copy,
  Check,
  ExternalLink,
  Zap,
  Shield,
  ArrowRight,
} from 'lucide-react';

const ESCROW_ADDRESS = '0x09431B4603E4Aa7511dD9341f2852fD031eA8C71';
const MTK_ADDRESS = '0xeab216ca4381a5c19e751f1471c55b452db0758a';
const SEPOLIA_CHAIN_ID = 11155111;

function CopyInline({ value }: { value: string }) {
  const [copied, setCopied] = useState(false);
  const copy = () => {
    navigator.clipboard.writeText(value).then(() => {
      setCopied(true);
      setTimeout(() => setCopied(false), 1500);
    });
  };
  return (
    <button
      onClick={copy}
      className="inline-flex items-center gap-1 font-mono text-xs bg-elevated px-2 py-0.5 rounded text-accent-violet hover:text-accent-blue transition-colors ml-1"
    >
      {value.slice(0, 10)}…{value.slice(-6)}
      {copied ? <Check className="w-3 h-3" /> : <Copy className="w-3 h-3" />}
    </button>
  );
}

function ExternalA({ href, children }: { href: string; children: React.ReactNode }) {
  return (
    <a
      href={href}
      target="_blank"
      rel="noopener noreferrer"
      className="text-accent-violet hover:text-accent-blue inline-flex items-center gap-1 underline underline-offset-2"
    >
      {children}
      <ExternalLink className="w-3 h-3" />
    </a>
  );
}

interface AccordionProps {
  title: string;
  children: React.ReactNode;
  defaultOpen?: boolean;
}

function Accordion({ title, children, defaultOpen = false }: AccordionProps) {
  const [open, setOpen] = useState(defaultOpen);
  return (
    <div className="border border-border-color rounded-xl overflow-hidden">
      <button
        onClick={() => setOpen(!open)}
        className="w-full flex items-center justify-between px-5 py-4 bg-card hover:bg-elevated transition-colors text-left"
      >
        <span className="font-semibold">{title}</span>
        {open ? <ChevronDown className="w-4 h-4 text-text-secondary" /> : <ChevronRight className="w-4 h-4 text-text-secondary" />}
      </button>
      {open && (
        <div className="px-5 py-4 bg-card border-t border-border-color space-y-3 text-sm text-text-secondary leading-relaxed">
          {children}
        </div>
      )}
    </div>
  );
}

function SectionHeader({ icon: Icon, title, subtitle }: { icon: React.ElementType; title: string; subtitle: string }) {
  return (
    <div className="flex items-start gap-4 mb-6">
      <div className="w-12 h-12 rounded-xl bg-gradient-to-br from-accent-violet to-accent-blue flex items-center justify-center flex-shrink-0">
        <Icon className="w-6 h-6 text-white" />
      </div>
      <div>
        <h2 className="text-xl font-bold mb-1">{title}</h2>
        <p className="text-text-secondary text-sm">{subtitle}</p>
      </div>
    </div>
  );
}

function StepRow({ n, title, desc }: { n: number; title: string; desc: string }) {
  return (
    <div className="flex gap-4">
      <div className="w-7 h-7 rounded-full bg-gradient-to-br from-accent-violet to-accent-blue flex items-center justify-center text-white text-xs font-bold flex-shrink-0 mt-0.5">
        {n}
      </div>
      <div>
        <div className="font-semibold text-sm mb-0.5">{title}</div>
        <div className="text-sm text-text-secondary">{desc}</div>
      </div>
    </div>
  );
}

function DealStage({ label, color }: { label: string; color: string }) {
  return (
    <span className={`px-2.5 py-1 rounded-full text-xs font-semibold ${color}`}>
      {label}
    </span>
  );
}

export function Docs() {
  return (
    <div className="max-w-4xl mx-auto space-y-10 pb-16">
      {/* Header */}
      <div>
        <h1 className="mb-2">Documentation</h1>
        <p className="text-text-secondary">
          Everything you need to know about Chainfluence — the decentralized creator economy platform powered by Chainlink.
        </p>
      </div>

      {/* ── OVERVIEW ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-4">
        <SectionHeader icon={Book} title="Platform Overview" subtitle="What Chainfluence is and how it works" />

        <p className="text-sm text-text-secondary leading-relaxed">
          Chainfluence connects <strong className="text-text-primary">advertisers</strong> who want to promote products with
          <strong className="text-text-primary"> influencers</strong> who own social media channels. Every deal is secured
          by an on-chain escrow smart contract — funds are locked when a deal is funded and released automatically
          once a Chainlink Functions job verifies the post met the agreed criteria (views, content, duration).
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4 mt-2">
          {[
            { icon: Shield, label: 'Trustless Escrow', desc: 'Funds locked on-chain, released by smart contract — no middleman.' },
            { icon: Zap, label: 'Automated Verification', desc: 'Chainlink Functions checks view counts, content match, and post duration.' },
            { icon: Radio, label: 'Verified Channels', desc: 'X OAuth ensures influencers own the accounts they claim.' },
          ].map(({ icon: Icon, label, desc }) => (
            <div key={label} className="bg-elevated rounded-xl p-4 space-y-2">
              <Icon className="w-5 h-5 text-accent-violet" />
              <div className="font-semibold text-sm">{label}</div>
              <div className="text-xs text-text-secondary">{desc}</div>
            </div>
          ))}
        </div>

        <div className="mt-4">
          <div className="text-xs font-semibold text-text-muted uppercase tracking-wider mb-3">Network &amp; Contracts</div>
          <div className="space-y-2 text-sm">
            <div className="flex items-center justify-between bg-elevated rounded-lg px-4 py-2.5">
              <span className="text-text-secondary">Network</span>
              <span className="font-medium">Ethereum Sepolia (chain ID {SEPOLIA_CHAIN_ID})</span>
            </div>
            <div className="flex items-center justify-between bg-elevated rounded-lg px-4 py-2.5">
              <span className="text-text-secondary">Escrow Contract</span>
              <span className="flex items-center gap-1">
                <CopyInline value={ESCROW_ADDRESS} />
                <ExternalA href={`https://sepolia.etherscan.io/address/${ESCROW_ADDRESS}`}>Etherscan</ExternalA>
              </span>
            </div>
            <div className="flex items-center justify-between bg-elevated rounded-lg px-4 py-2.5">
              <span className="text-text-secondary">MTK Token</span>
              <span className="flex items-center gap-1">
                <CopyInline value={MTK_ADDRESS} />
                <ExternalA href={`https://sepolia.etherscan.io/address/${MTK_ADDRESS}`}>Etherscan</ExternalA>
              </span>
            </div>
          </div>
        </div>
      </section>

      {/* ── GETTING STARTED ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={Book} title="Getting Started" subtitle="Set up your wallet and get testnet tokens" />

        <div className="space-y-4">
          <StepRow
            n={1}
            title="Install a wallet"
            desc="Chainfluence works with any EVM wallet. MetaMask is recommended. Install it from metamask.io and create an account."
          />
          <StepRow
            n={2}
            title="Switch to Sepolia testnet"
            desc="Open MetaMask → click the network dropdown → Add Network → Sepolia (or search 'Sepolia' in the list). Chain ID 11155111."
          />
          <StepRow
            n={3}
            title="Get Sepolia ETH"
            desc="You need ETH to pay for gas. Use a faucet — Google 'Sepolia faucet' or use the Alchemy or Chainlink faucets. Paste your wallet address and request 0.5 ETH."
          />
          <StepRow
            n={4}
            title="Get MTK test tokens"
            desc="MTK is the payment token used for all deals on Chainfluence. Ask in the project Discord or request from the deployer address. MTK has 18 decimals."
          />
          <StepRow
            n={5}
            title="Connect wallet on Chainfluence"
            desc="Click 'Connect Wallet' in the top-right corner. Sign the authentication message (SIWE — Sign-In With Ethereum) — this creates your account and issues a JWT session token. No gas required for sign-in."
          />
        </div>
      </section>

      {/* ── FOR ADVERTISERS ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={Megaphone} title="For Advertisers" subtitle="Create campaigns and manage deals with influencers" />

        <p className="text-sm text-text-secondary">There are two ways to start a deal — pick either one:</p>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div className="bg-elevated rounded-xl p-4 border border-border-color space-y-2">
            <div className="font-semibold text-sm text-accent-violet">Option A — Post a public campaign</div>
            <div className="text-xs text-text-secondary leading-relaxed">
              Go to <strong className="text-text-primary">My Campaigns → Create Campaign</strong>. Set the content, budget, min views, duration, deadline, and categories. Your campaign is listed publicly in the Marketplace — influencers find it and apply. No wallet transaction at creation.
            </div>
          </div>
          <div className="bg-elevated rounded-xl p-4 border border-border-color space-y-2">
            <div className="font-semibold text-sm text-accent-violet">Option B — Propose directly to an influencer</div>
            <div className="text-xs text-text-secondary leading-relaxed">
              Go to <strong className="text-text-primary">Marketplace</strong>, find a channel you like, and click <strong className="text-text-primary">Propose Deal</strong>. Select an existing campaign or create one on the spot. The deal goes straight to that influencer — the campaign is hidden from the public listing.
            </div>
          </div>
        </div>

        <div className="space-y-3 mt-2">
          <Accordion title="Fund an Accepted Deal">
            <p>When an influencer accepts your proposal, the deal moves to <strong className="text-text-primary">accepted</strong> status.</p>
            <p>Open the deal from <strong className="text-text-primary">My Campaigns → (campaign) → Deals</strong> or from the deal link.</p>
            <p>Click <strong className="text-text-primary">Fund Deal</strong>. This triggers two wallet transactions:</p>
            <ol className="list-decimal pl-5 space-y-1">
              <li><strong className="text-text-primary">ERC-20 Approve</strong> — authorise the escrow contract to spend your MTK tokens.</li>
              <li><strong className="text-text-primary">Deposit</strong> — locks the tokens into the escrow contract on-chain. The campaign is registered on-chain at this point.</li>
            </ol>
            <p>After both transactions confirm, the deal status becomes <strong className="text-text-primary">funded</strong>.</p>
          </Accordion>

          <Accordion title="Wait for Verification">
            <p>Once the influencer submits a tweet URL, the deal enters <strong className="text-text-primary">verifying</strong> status. The CRE workflow checks:</p>
            <ul className="list-disc pl-5 space-y-1">
              <li><strong className="text-text-primary">Content match</strong> — the tweet text must hash to the agreed content hash (keccak256). Fails immediately if mismatched.</li>
              <li><strong className="text-text-primary">Edit detection</strong> — the tweet must not have been edited after posting. Fails if edited.</li>
              <li><strong className="text-text-primary">Min views</strong> — impression count must meet or exceed the campaign threshold. Funds release only when this passes.</li>
            </ul>
            <p className="text-text-muted text-xs mt-1">Note: post duration (uptime) is recorded and passed to the contract but is not currently enforced as a blocking condition.</p>
            <p>If all checks pass → <strong className="text-text-primary">completed</strong>, funds released to influencer.<br />
            If checks fail → <strong className="text-text-primary">refunded</strong>, your MTK is returned to your wallet.</p>
          </Accordion>

          <Accordion title="Refunds &amp; Disputes">
            <p>If the expiry deadline passes before a deal is completed, you can trigger a refund from the deal view page.</p>
            <p>If verification fails, the smart contract automatically refunds the escrow balance to your wallet.</p>
            <p>If a deal is in dispute, contact support — manual resolution is currently handled off-chain.</p>
          </Accordion>
        </div>
      </section>

      {/* ── FOR INFLUENCERS ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={Radio} title="For Influencers" subtitle="Register channels, apply to campaigns, and get paid" />

        <div className="space-y-3">
          <Accordion title="Step 1 — Register Your X Channel" defaultOpen>
            <p>Go to <strong className="text-text-primary">My Channels → Connect X Account</strong>.</p>
            <p>You'll be redirected to X (Twitter) to authorize Chainfluence. After you approve, X redirects you back and we automatically pull your:</p>
            <ul className="list-disc pl-5 space-y-1">
              <li>Verified handle (@username)</li>
              <li>Follower count</li>
              <li>Profile image</li>
              <li>X user ID (used to verify tweet ownership)</li>
            </ul>
            <p>Your channel appears with a blue verification badge. You cannot register a channel you don't own — the OAuth proves ownership.</p>
          </Accordion>

          <Accordion title="Step 2 — Set Your Price &amp; Categories">
            <p>After registration, edit your channel from <strong className="text-text-primary">My Channels</strong>:</p>
            <ul className="list-disc pl-5 space-y-1">
              <li><strong className="text-text-primary">Price per post</strong> — how much MTK you charge per promoted tweet.</li>
              <li><strong className="text-text-primary">Categories</strong> — topics you cover (DeFi, NFT, Gaming, Tech, etc.). This helps advertisers find you in the marketplace.</li>
              <li><strong className="text-text-primary">Active / Inactive toggle</strong> — hide your channel from the marketplace without deleting it.</li>
            </ul>
          </Accordion>

          <Accordion title="Step 3 — Find &amp; Apply to Campaigns">
            <p>Browse the <strong className="text-text-primary">Marketplace</strong> to see active campaigns looking for influencers.</p>
            <p>Each campaign shows the required content, budget, minimum views, categories, and deadline. Click <strong className="text-text-primary">View Details</strong> to see the full brief and apply.</p>
            <p>Alternatively, advertisers may send you a direct deal proposal. You'll see incoming proposals in <strong className="text-text-primary">My Deals</strong>.</p>
          </Accordion>

          <Accordion title="Step 4 — Accept and Post">
            <p>When you receive a proposal (or apply to a campaign and the advertiser responds), the deal shows in <strong className="text-text-primary">My Deals</strong>.</p>
            <p>Review the terms: content requirements, budget, minimum views, deadline. Click <strong className="text-text-primary">Accept Deal</strong>.</p>
            <p>After the advertiser funds the deal (locks tokens in escrow), post the agreed content to your verified X account and submit the tweet URL in the deal view.</p>
            <p>The deal moves to <strong className="text-text-primary">verifying</strong> — wait for the Chainlink verification to complete.</p>
          </Accordion>

          <Accordion title="Step 5 — Get Paid">
            <p>Once Chainlink Functions confirms your tweet meets all criteria, the escrow contract automatically releases the MTK tokens to your wallet.</p>
            <p>No manual action needed — payment is trustless and on-chain. The deal status becomes <strong className="text-text-primary">completed</strong>.</p>
            <p>MTK tokens appear in your wallet on Sepolia. You can view them in MetaMask by adding the token address: <CopyInline value={MTK_ADDRESS} /></p>
          </Accordion>
        </div>
      </section>

      {/* ── DEAL LIFECYCLE ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={ArrowRight} title="Deal Lifecycle" subtitle="All possible deal states and transitions" />

        <div className="overflow-x-auto">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-border-color">
                <th className="text-left py-3 px-2 text-text-muted font-medium">Status</th>
                <th className="text-left py-3 px-2 text-text-muted font-medium">Meaning</th>
                <th className="text-left py-3 px-2 text-text-muted font-medium">Who can act</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-border-color">
              {[
                { status: 'proposed', color: 'bg-blue-500/10 text-blue-400', meaning: 'Deal sent, awaiting influencer response', actor: 'Influencer (accept/decline)' },
                { status: 'accepted', color: 'bg-yellow-500/10 text-yellow-400', meaning: 'Influencer accepted, awaiting advertiser funding', actor: 'Advertiser (fund)' },
                { status: 'funded', color: 'bg-violet-500/10 text-violet-400', meaning: 'Tokens locked in escrow, influencer should post', actor: 'Influencer (submit tweet URL)' },
                { status: 'posted', color: 'bg-cyan-500/10 text-cyan-400', meaning: 'Tweet URL submitted, queued for verification', actor: 'System (Chainlink)' },
                { status: 'verifying', color: 'bg-orange-500/10 text-orange-400', meaning: 'Chainlink Functions job running on-chain', actor: 'System (automated)' },
                { status: 'completed', color: 'bg-green-500/10 text-green-400', meaning: 'Verification passed, payment released', actor: 'Terminal state' },
                { status: 'refunded', color: 'bg-red-500/10 text-red-400', meaning: 'Verification failed or expired, funds returned', actor: 'Terminal state' },
                { status: 'disputed', color: 'bg-gray-500/10 text-gray-400', meaning: 'Manual review required', actor: 'Support team' },
              ].map(({ status, color, meaning, actor }) => (
                <tr key={status}>
                  <td className="py-3 px-2"><DealStage label={status} color={color} /></td>
                  <td className="py-3 px-2 text-text-secondary">{meaning}</td>
                  <td className="py-3 px-2 text-text-secondary">{actor}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </section>

      {/* ── SMART CONTRACTS ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={Code} title="Smart Contracts" subtitle="On-chain architecture and Chainlink integration" />

        <div className="space-y-3">
          <Accordion title="Escrow Contract" defaultOpen>
            <p>The escrow contract (<CopyInline value={ESCROW_ADDRESS} />) holds all campaign funds until verification completes.</p>
            <p>Key functions:</p>
            <ul className="list-disc pl-5 space-y-1 font-mono text-xs">
              <li><span className="text-accent-violet">deposit(campaignId, contentHash, tokenAddress, amount, minViews, duration)</span> — advertiser locks tokens. Creates on-chain campaign record.</li>
              <li><span className="text-accent-violet">submitPostUrl(campaignId, postUrl)</span> — influencer submits their tweet. Triggers Chainlink Functions request.</li>
              <li><span className="text-accent-violet">refund(campaignId)</span> — callable after expiry deadline if deal not completed.</li>
            </ul>
            <p className="font-normal text-text-secondary">Funds flow: Advertiser wallet → ERC-20 approve → <code className="bg-elevated px-1 rounded">deposit()</code> → escrow holds tokens → Chainlink fulfills → tokens sent to influencer or refunded.</p>
          </Accordion>

          <Accordion title="Chainlink Functions Integration">
            <p>Chainfluence uses <strong className="text-text-primary">Chainlink Functions</strong> to fetch off-chain Twitter/X data and deliver it to the smart contract trustlessly.</p>
            <p>When <code className="bg-elevated px-1 rounded text-accent-violet">submitPostUrl</code> is called, the contract sends a Chainlink Functions request with:</p>
            <ul className="list-disc pl-5 space-y-1">
              <li>The tweet URL to verify</li>
              <li>The expected content hash (keccak256 of the campaign content)</li>
              <li>The minimum view count threshold</li>
              <li>The required post duration in seconds</li>
            </ul>
            <p>The DON (Decentralized Oracle Network) runs the JavaScript source, fetches tweet data from the X API, and returns a pass/fail result. The contract's <code className="bg-elevated px-1 rounded text-accent-violet">fulfillRequest</code> callback then releases or refunds the escrow.</p>
          </Accordion>

          <Accordion title="MTK Token">
            <p>MTK (<CopyInline value={MTK_ADDRESS} />) is a standard ERC-20 token used for all payments on Chainfluence on Sepolia testnet.</p>
            <p><strong className="text-text-primary">Decimals:</strong> 18 — same as ETH. 1 MTK = <code className="bg-elevated px-1 rounded">1000000000000000000</code> (10^18) raw units.</p>
            <p>To add MTK to MetaMask: Assets → Import tokens → paste the address above.</p>
            <p>On mainnet, MTK would be replaced by a real ERC-20 stablecoin or governance token.</p>
          </Accordion>

          <Accordion title="Content Hash">
            <p>When funding a deal, the advertiser's campaign content text is hashed using <code className="bg-elevated px-1 rounded text-accent-violet">keccak256</code> and stored on-chain.</p>
            <p>During verification, Chainlink Functions checks that the submitted tweet's text matches this hash. This prevents influencers from posting different content than agreed.</p>
            <p>The hash is computed client-side in the frontend as: <code className="bg-elevated px-1 rounded text-accent-violet">keccak256(toBytes(contentText.toLowerCase().trim()))</code></p>
          </Accordion>
        </div>
      </section>

      {/* ── FAQ ── */}
      <section className="bg-card border border-border-color rounded-xl p-6 space-y-5">
        <SectionHeader icon={Book} title="FAQ" subtitle="Frequently asked questions" />

        <div className="space-y-3">
          <Accordion title="What happens if the influencer edits or deletes the tweet?">
            <p>The Chainlink verification job checks for tweet edits and deletions. If the tweet is edited after posting, the <code className="bg-elevated px-1 rounded">isEdited</code> flag is flagged. If it's deleted, verification fails and the deal is refunded to the advertiser.</p>
          </Accordion>

          <Accordion title="Can I register more than one X account?">
            <p>Yes. You can connect multiple X accounts to your Chainfluence wallet. Each OAuth connection creates a separate channel. Manage them all from <strong className="text-text-primary">My Channels</strong>.</p>
          </Accordion>

          <Accordion title="What if I post the tweet but verification never completes?">
            <p>Chainlink Functions jobs can occasionally time out or encounter network issues. If verification is stuck, the deal will remain in <code className="bg-elevated px-1 rounded">verifying</code> status. Use the <strong className="text-text-primary">Refresh Status</strong> button on the deal page to manually re-sync from chain. If the issue persists, contact support.</p>
          </Accordion>

          <Accordion title="Can I cancel a deal after funding?">
            <p>Once funds are locked in escrow (status: funded), the deal cannot be unilaterally cancelled by the advertiser — this protects influencers who are preparing content. After the <strong className="text-text-primary">expiry deadline</strong> passes without completion, you can trigger a refund.</p>
          </Accordion>

          <Accordion title="Is my wallet address public?">
            <p>Wallet addresses are public on-chain by nature. Your address is visible in the escrow contract events. However, your display name and X handle are stored off-chain in our database and only shared with counterparties on active deals.</p>
          </Accordion>

          <Accordion title="What is the role of 'both' in settings?">
            <p>Your account role determines which features are highlighted in the UI. <strong className="text-text-primary">Advertiser</strong> — see My Campaigns primarily. <strong className="text-text-primary">Influencer</strong> — see My Channels and My Deals primarily. <strong className="text-text-primary">Both</strong> — all features equally visible. Role is cosmetic; you can always access all pages regardless of role.</p>
          </Accordion>
        </div>
      </section>

      {/* ── RESOURCES ── */}
      <section className="bg-card border border-border-color rounded-xl p-6">
        <h3 className="mb-4">External Resources</h3>
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
          {[
            { label: 'Sepolia Testnet Explorer', url: 'https://sepolia.etherscan.io', desc: 'View transactions and contracts' },
            { label: 'Sepolia Faucet (Alchemy)', url: 'https://sepoliafaucet.com', desc: 'Get free Sepolia ETH' },
            { label: 'Chainlink Functions Docs', url: 'https://docs.chain.link/chainlink-functions', desc: 'Learn about the oracle network' },
            { label: 'RainbowKit', url: 'https://rainbowkit.com', desc: 'Wallet connection library used in this app' },
            { label: 'MetaMask', url: 'https://metamask.io', desc: 'Recommended wallet for Chainfluence' },
            { label: 'X Developer Portal', url: 'https://developer.x.com', desc: 'X OAuth and API documentation' },
          ].map(({ label, url, desc }) => (
            <a
              key={label}
              href={url}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-start gap-3 p-3 bg-elevated rounded-xl hover:border-accent-violet/40 border border-transparent transition-all group"
            >
              <ExternalLink className="w-4 h-4 text-accent-violet mt-0.5 flex-shrink-0" />
              <div>
                <div className="font-semibold group-hover:text-accent-violet transition-colors">{label}</div>
                <div className="text-xs text-text-secondary">{desc}</div>
              </div>
            </a>
          ))}
        </div>
      </section>
    </div>
  );
}
