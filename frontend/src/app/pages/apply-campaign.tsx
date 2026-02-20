import { useState, useEffect } from 'react';
import { useParams, useNavigate, Link } from 'react-router';
import { Twitter, Eye, Clock, Calendar, DollarSign, ChevronLeft, Loader2, Check, Info } from 'lucide-react';
import { formatUnits } from 'viem';
import { GradientButton } from '../components/gradient-button';
import { StatusBadge } from '../components/status-badge';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface Campaign {
  id: number;
  title: string;
  contentText: string;
  categories: string[];
  amount: string;
  minViews: number;
  campaignDuration: number;
  expiryDeadline: string;
  status: string;
}

interface Channel {
  id: number;
  platformHandle: string;
  followerCount: number;
  profileImageUrl?: string;
  isActive: boolean;
}

function formatDuration(seconds: number): string {
  if (!seconds) return 'No requirement';
  const days = Math.floor(seconds / 86400);
  if (days >= 1) return `${days} day${days > 1 ? 's' : ''}`;
  const hours = Math.floor(seconds / 3600);
  return `${hours} hour${hours > 1 ? 's' : ''}`;
}

export function ApplyCampaign() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [campaign, setCampaign] = useState<Campaign | null>(null);
  const [channels, setChannels] = useState<Channel[]>([]);
  const [selectedChannel, setSelectedChannel] = useState<number | null>(null);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState(false);
  const [dealId, setDealId] = useState<number | null>(null);

  useEffect(() => {
    Promise.all([
      apiClient(`/api/campaigns/${id}`),
      apiClient('/api/channels'),
    ])
      .then(([campData, chanData]) => {
        setCampaign(campData.campaign);
        const active = (chanData.channels || []).filter((c: Channel) => c.isActive);
        setChannels(active);
        if (active.length === 1) setSelectedChannel(active[0].id);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message || 'Failed to load data');
        setLoading(false);
      });
  }, [id]);

  const handleApply = async () => {
    if (!selectedChannel) {
      setError('Please select a channel');
      return;
    }
    setSubmitting(true);
    setError('');
    try {
      const { deal } = await apiClient('/api/deals', {
        method: 'POST',
        body: JSON.stringify({
          campaignId: parseInt(id!),
          channelId: selectedChannel,
        }),
      });
      setDealId(deal.id);
      setSuccess(true);
    } catch (err: any) {
      setError(err.message || 'Failed to submit application');
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-32">
        <Loader2 className="w-8 h-8 animate-spin text-accent-violet" />
      </div>
    );
  }

  if (!campaign) {
    return (
      <div className="max-w-2xl mx-auto text-center py-20">
        <h2 className="mb-2">Campaign not found</h2>
        <p className="text-text-secondary mb-6">{error || 'This campaign does not exist.'}</p>
        <Link to="/marketplace">
          <GradientButton>Back to Marketplace</GradientButton>
        </Link>
      </div>
    );
  }

  if (success) {
    return (
      <div className="max-w-2xl mx-auto">
        <div className="bg-card border border-border-color rounded-xl p-8 text-center space-y-6">
          <div className="w-16 h-16 rounded-full bg-success/10 flex items-center justify-center mx-auto">
            <Check className="w-8 h-8 text-success" />
          </div>
          <div>
            <h2 className="text-success mb-2">Application Submitted!</h2>
            <p className="text-text-secondary">
              Your proposal for "{campaign.title}" has been sent to the advertiser.
            </p>
          </div>
          <div className="flex gap-3 justify-center">
            <button
              onClick={() => navigate('/marketplace')}
              className="px-6 py-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all font-medium"
            >
              Back to Marketplace
            </button>
            {dealId && (
              <GradientButton onClick={() => navigate(`/deal/${dealId}`)}>
                View Deal
              </GradientButton>
            )}
          </div>
        </div>
      </div>
    );
  }

  const amount = parseFloat(formatUnits(BigInt(campaign.amount), MTK_TOKEN.decimals)).toString();
  const deadline = new Date(campaign.expiryDeadline);

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <div>
        <Link
          to="/marketplace"
          className="inline-flex items-center gap-1 text-sm text-text-secondary hover:text-text-primary transition-colors mb-4"
        >
          <ChevronLeft className="w-4 h-4" />
          Back to Marketplace
        </Link>
        <h1 className="mb-1">Apply to Campaign</h1>
        <p className="text-text-secondary">Review the campaign details and submit your proposal</p>
      </div>

      {/* Campaign Summary */}
      <div className="bg-card border border-border-color rounded-xl p-6 space-y-4">
        <div className="flex items-start gap-3">
          <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
            <Twitter className="w-5 h-5 text-sky-500" />
          </div>
          <div className="flex-1">
            <h2 className="mb-1">{campaign.title}</h2>
            <StatusBadge status={campaign.status} />
          </div>
        </div>

        <div className="bg-background border border-border-color rounded-lg p-4 font-mono text-sm whitespace-pre-wrap break-words">
          {campaign.contentText}
        </div>

        {campaign.categories.length > 0 && (
          <div className="flex flex-wrap gap-2">
            {campaign.categories.map((cat) => (
              <span key={cat} className="px-3 py-1 bg-elevated rounded-full text-xs text-text-secondary">
                {cat}
              </span>
            ))}
          </div>
        )}

        <div className="grid grid-cols-2 sm:grid-cols-4 gap-4 pt-2">
          <div>
            <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
              <DollarSign className="w-3 h-3" /> Budget
            </div>
            <div className="font-semibold">{amount} {MTK_TOKEN.symbol}</div>
          </div>
          <div>
            <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
              <Eye className="w-3 h-3" /> Min Views
            </div>
            <div className="font-semibold">{campaign.minViews > 0 ? campaign.minViews.toLocaleString() : 'None'}</div>
          </div>
          <div>
            <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
              <Clock className="w-3 h-3" /> Duration
            </div>
            <div className="font-semibold">{formatDuration(campaign.campaignDuration)}</div>
          </div>
          <div>
            <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
              <Calendar className="w-3 h-3" /> Deadline
            </div>
            <div className="font-semibold">{deadline.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })}</div>
          </div>
        </div>
      </div>

      {/* Select Channel */}
      <div className="bg-card border border-border-color rounded-xl p-6 space-y-4">
        <h2>Select Your Channel</h2>
        <p className="text-sm text-text-secondary">
          Choose which X account you'll use to post the campaign content.
        </p>

        {channels.length === 0 ? (
          <div className="text-center py-8">
            <p className="text-text-secondary mb-4">You don't have any registered channels yet.</p>
            <Link to="/my-channels">
              <GradientButton>Register a Channel</GradientButton>
            </Link>
          </div>
        ) : (
          <div className="space-y-3">
            {channels.map((channel) => (
              <button
                key={channel.id}
                onClick={() => setSelectedChannel(channel.id)}
                className={`w-full p-4 rounded-xl border-2 transition-all text-left flex items-center gap-4 ${
                  selectedChannel === channel.id
                    ? 'border-accent-violet bg-accent-violet/5'
                    : 'border-border-color hover:border-accent-violet/30'
                }`}
              >
                {channel.profileImageUrl ? (
                  <img
                    src={channel.profileImageUrl}
                    alt={`@${channel.platformHandle}`}
                    className="w-10 h-10 rounded-full flex-shrink-0"
                  />
                ) : (
                  <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
                    <Twitter className="w-5 h-5 text-sky-500" />
                  </div>
                )}
                <div className="flex-1">
                  <div className="font-semibold text-accent-violet">@{channel.platformHandle}</div>
                  <div className="text-sm text-text-secondary">{channel.followerCount.toLocaleString()} followers</div>
                </div>
                {selectedChannel === channel.id && (
                  <Check className="w-5 h-5 text-accent-violet flex-shrink-0" />
                )}
              </button>
            ))}
          </div>
        )}
      </div>

      {/* What you're agreeing to */}
      <div className="p-4 bg-accent-blue/10 border border-accent-blue/20 rounded-lg flex gap-3">
        <Info className="w-5 h-5 text-accent-blue flex-shrink-0 mt-0.5" />
        <div className="text-sm text-text-secondary">
          By applying, you agree to post the campaign content on your X account and keep it up for the required duration.
          The advertiser will review your proposal and may accept or decline it.
        </div>
      </div>

      {/* Error */}
      {error && (
        <div className="p-4 bg-error/10 border border-error/20 rounded-lg flex gap-3">
          <Info className="w-5 h-5 text-error flex-shrink-0 mt-0.5" />
          <div className="text-sm text-error">{error}</div>
        </div>
      )}

      {/* Submit */}
      {channels.length > 0 && (
        <GradientButton
          onClick={handleApply}
          disabled={submitting || !selectedChannel}
          loading={submitting}
          className="w-full h-13"
        >
          {submitting ? 'Submitting...' : 'Submit Application'}
        </GradientButton>
      )}
    </div>
  );
}
