import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router';
import { Twitter, Copy, ExternalLink, Check, Clock, Eye, Calendar, Tag, ChevronLeft, Loader2, FileText } from 'lucide-react';
import { formatUnits } from 'viem';
import { StatusBadge } from '../components/status-badge';
import { Jazzicon } from '../utils/jazzicon';
import { GradientButton } from '../components/gradient-button';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface Campaign {
  id: number;
  onchainCampaignId?: number;
  advertiserId: number;
  title: string;
  contentText: string;
  contentHash?: string;
  tokenAddress: string;
  amount: string;
  minViews: number;
  campaignDuration: number;
  expiryDeadline: string;
  chainId: number;
  status: string;
  txHash?: string;
  categories: string[];
  createdAt: string;
  updatedAt: string;
}

interface CampaignDeal {
  id: number;
  campaignId: number;
  channelId?: number;
  influencerId: number;
  status: string;
  proposedBy: string;
  postUrl?: string;
  onchainCampaignId?: number;
  createdAt: string;
}

function formatDuration(seconds: number): string {
  if (!seconds) return 'No requirement';
  const days = Math.floor(seconds / 86400);
  if (days >= 1) return `${days} day${days > 1 ? 's' : ''}`;
  const hours = Math.floor(seconds / 3600);
  return `${hours} hour${hours > 1 ? 's' : ''}`;
}

function copyToClipboard(text: string) {
  navigator.clipboard.writeText(text);
}

export function CampaignView() {
  const { id } = useParams();
  const [campaign, setCampaign] = useState<Campaign | null>(null);
  const [deals, setDeals] = useState<CampaignDeal[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!id) return;
    Promise.all([
      apiClient(`/api/campaigns/${id}`),
      apiClient(`/api/campaigns/${id}/deals`).catch(() => ({ deals: [] })),
    ])
      .then(([campData, dealsData]) => {
        setCampaign(campData.campaign);
        setDeals((dealsData as { deals?: CampaignDeal[] }).deals ?? []);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message || 'Failed to load campaign');
        setLoading(false);
      });
  }, [id]);

  if (loading) {
    return (
      <div className="flex items-center justify-center py-32">
        <Loader2 className="w-8 h-8 animate-spin text-accent-violet" />
      </div>
    );
  }

  if (error || !campaign) {
    return (
      <div className="max-w-2xl mx-auto text-center py-20">
        <h2 className="mb-2">Campaign not found</h2>
        <p className="text-text-secondary mb-6">{error || 'This campaign does not exist.'}</p>
        <Link to="/my-campaigns">
          <GradientButton>Back to My Campaigns</GradientButton>
        </Link>
      </div>
    );
  }

  const rawAmount = campaign.amount;
  const amount = parseFloat(formatUnits(BigInt(rawAmount), MTK_TOKEN.decimals)).toString();
  const deadline = new Date(campaign.expiryDeadline);
  const created = new Date(campaign.createdAt);

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      {/* Back link + Header */}
      <div>
        <Link
          to="/my-campaigns"
          className="inline-flex items-center gap-1 text-sm text-text-secondary hover:text-text-primary transition-colors mb-4"
        >
          <ChevronLeft className="w-4 h-4" />
          Back to My Campaigns
        </Link>

        <div className="flex items-center justify-between">
          <div>
            <h1 className="mb-1">{campaign.title}</h1>
            <p className="text-text-secondary">Campaign #{campaign.id}</p>
          </div>
          <StatusBadge status={campaign.status} className="text-base px-4 py-2" />
        </div>
      </div>

      {/* Main Content */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Left Column */}
        <div className="lg:col-span-2 space-y-6">
          {/* Tweet Content */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <div className="flex items-center gap-3 mb-6">
              <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center">
                <Twitter className="w-5 h-5 text-sky-500" />
              </div>
              <h2>Tweet Content</h2>
            </div>

            <div className="bg-background border border-border-color rounded-lg p-4 font-mono text-sm whitespace-pre-wrap break-words">
              {campaign.contentText}
            </div>

            {campaign.contentHash && (
              <div className="mt-4 flex items-center gap-2 text-xs text-text-muted">
                <span>Content Hash:</span>
                <code className="font-mono bg-elevated px-2 py-1 rounded">
                  {campaign.contentHash.slice(0, 10)}...{campaign.contentHash.slice(-8)}
                </code>
                <button onClick={() => copyToClipboard(campaign.contentHash!)}>
                  <Copy className="w-3 h-3 hover:text-text-primary transition-colors" />
                </button>
              </div>
            )}
          </div>

          {/* Requirements */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h2 className="mb-6">Requirements</h2>
            <div className="grid grid-cols-1 sm:grid-cols-3 gap-6">
              <div className="flex items-start gap-3">
                <div className="w-10 h-10 rounded-lg bg-accent-violet/10 flex items-center justify-center flex-shrink-0">
                  <Eye className="w-5 h-5 text-accent-violet" />
                </div>
                <div>
                  <div className="text-sm text-text-muted mb-1">Minimum Views</div>
                  <div className="font-semibold text-lg">
                    {campaign.minViews > 0 ? campaign.minViews.toLocaleString() : 'None'}
                  </div>
                </div>
              </div>

              <div className="flex items-start gap-3">
                <div className="w-10 h-10 rounded-lg bg-accent-blue/10 flex items-center justify-center flex-shrink-0">
                  <Clock className="w-5 h-5 text-accent-blue" />
                </div>
                <div>
                  <div className="text-sm text-text-muted mb-1">Post Duration</div>
                  <div className="font-semibold text-lg">
                    {formatDuration(campaign.campaignDuration)}
                  </div>
                </div>
              </div>

              <div className="flex items-start gap-3">
                <div className="w-10 h-10 rounded-lg bg-warning/10 flex items-center justify-center flex-shrink-0">
                  <Calendar className="w-5 h-5 text-warning" />
                </div>
                <div>
                  <div className="text-sm text-text-muted mb-1">Deadline</div>
                  <div className="font-semibold text-lg">
                    {deadline.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}
                  </div>
                </div>
              </div>
            </div>
          </div>

          {/* Categories */}
          {campaign.categories.length > 0 && (
            <div className="bg-card border border-border-color rounded-xl p-6">
              <div className="flex items-center gap-3 mb-4">
                <Tag className="w-5 h-5 text-text-muted" />
                <h2>Categories</h2>
              </div>
              <div className="flex flex-wrap gap-2">
                {campaign.categories.map((cat) => (
                  <span
                    key={cat}
                    className="px-4 py-1.5 bg-accent-violet/10 border border-accent-violet/20 rounded-full text-sm"
                  >
                    {cat}
                  </span>
                ))}
              </div>
            </div>
          )}
        </div>

        {/* Right Column */}
        <div className="lg:col-span-1 space-y-6">
          {/* Payment Info */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h3 className="mb-4">Payment</h3>

            <div className="text-center mb-6">
              <div className="text-3xl font-bold mb-1">{amount} {MTK_TOKEN.symbol}</div>
              <div className="text-sm text-text-muted">Campaign Budget</div>
            </div>

            <div className="bg-elevated rounded-lg p-4 space-y-2 text-sm">
              <div className="flex justify-between">
                <span className="text-text-muted">Influencer receives:</span>
                <span className="font-medium">{amount} {MTK_TOKEN.symbol}</span>
              </div>
              <div className="flex justify-between">
                <span className="text-text-muted">Platform fee (2.5%):</span>
                <span className="font-medium">
                  {(parseFloat(amount) * 0.025).toFixed(2)} {MTK_TOKEN.symbol}
                </span>
              </div>
              <div className="h-px bg-border-color my-1" />
              <div className="flex justify-between font-semibold">
                <span>Total deposit:</span>
                <span>
                  {(parseFloat(amount) * 1.025).toFixed(2)} {MTK_TOKEN.symbol}
                </span>
              </div>
            </div>
          </div>

          {/* Campaign Info */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h3 className="mb-4">Details</h3>

            <div className="space-y-3 text-sm">
              <div className="flex justify-between">
                <span className="text-text-muted">Status:</span>
                <StatusBadge status={campaign.status} />
              </div>

              <div className="flex justify-between">
                <span className="text-text-muted">Created:</span>
                <span>{created.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}</span>
              </div>

              <div className="flex justify-between">
                <span className="text-text-muted">Chain:</span>
                <span>Sepolia</span>
              </div>

              <div className="flex justify-between">
                <span className="text-text-muted">Token:</span>
                <span className="font-mono text-xs">
                  {campaign.tokenAddress.slice(0, 6)}...{campaign.tokenAddress.slice(-4)}
                </span>
              </div>

              {campaign.onchainCampaignId != null && (
                <>
                  <div className="h-px bg-border-color" />
                  <div className="flex justify-between">
                    <span className="text-text-muted">On-chain ID:</span>
                    <span className="font-semibold">{campaign.onchainCampaignId}</span>
                  </div>
                </>
              )}

              {campaign.txHash && (
                <div>
                  <div className="text-text-muted mb-1">Transaction:</div>
                  <a
                    href={`https://sepolia.etherscan.io/tx/${campaign.txHash}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="font-mono text-xs text-accent-violet hover:text-accent-blue flex items-center gap-1"
                  >
                    {campaign.txHash.slice(0, 10)}...{campaign.txHash.slice(-8)}
                    <ExternalLink className="w-3 h-3" />
                  </a>
                </div>
              )}
            </div>
          </div>

          {/* Status / Actions — no campaign-level funding; advertiser pays when they accept a deal */}
          {campaign.status === 'draft' && (
            <div className="bg-card border border-border-color rounded-xl p-6">
              <h3 className="mb-4">Status</h3>
              <div className="flex items-center gap-3 p-3 bg-accent-violet/10 border border-accent-violet/20 rounded-lg">
                <Check className="w-5 h-5 text-accent-violet" />
                <div className="text-sm">
                  <div className="font-semibold text-accent-violet">Listed on marketplace</div>
                  <div className="text-text-secondary">Influencers can propose deals. When you accept a deal, you’ll deposit funds in the deal view.</div>
                </div>
              </div>
            </div>
          )}

          {campaign.status === 'funded' && (
            <div className="bg-card border border-border-color rounded-xl p-6">
              <h3 className="mb-4">Status</h3>
              <div className="flex items-center gap-3 p-3 bg-success/10 border border-success/20 rounded-lg">
                <Check className="w-5 h-5 text-success" />
                <div className="text-sm">
                  <div className="font-semibold text-success">Funded & Active</div>
                  <div className="text-text-secondary">Waiting for influencer proposals</div>
                </div>
              </div>
            </div>
          )}

          {/* Deals for this campaign (one campaign can have multiple deals) */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h3 className="mb-4 flex items-center gap-2">
              <FileText className="w-5 h-5 text-text-muted" />
              Deals ({deals.length})
            </h3>
            {deals.length === 0 ? (
              <p className="text-sm text-text-secondary">No proposals yet. Influencers can apply from the marketplace.</p>
            ) : (
              <ul className="space-y-2">
                {deals.map((deal) => (
                  <li key={deal.id}>
                    <Link
                      to={`/deal/${deal.id}`}
                      className="flex items-center justify-between p-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all group"
                    >
                      <div className="flex items-center gap-2 min-w-0">
                        <Jazzicon address={deal.proposedBy} size={24} />
                        <span className="font-mono text-sm truncate">{deal.proposedBy.slice(0, 8)}...{deal.proposedBy.slice(-6)}</span>
                      </div>
                      <div className="flex items-center gap-2 flex-shrink-0">
                        <StatusBadge status={deal.status as any} />
                        <span className="text-text-muted text-xs">Deal #{deal.id}</span>
                      </div>
                    </Link>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
