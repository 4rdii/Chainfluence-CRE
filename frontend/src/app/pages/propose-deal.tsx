import { useState, useEffect } from 'react';
import { useParams, Link, useNavigate } from 'react-router';
import { Twitter, Users, Loader2, ArrowLeft, Info } from 'lucide-react';
import { Input } from '../components/ui/input';
import { Textarea } from '../components/ui/textarea';
import { GradientButton } from '../components/gradient-button';
import { parseUnits } from 'viem';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface Channel {
  id: number;
  platformHandle: string;
  followerCount: number;
  categories: string[];
  pricePerPost: number;
}

function parseDuration(val: string): number {
  const map: Record<string, number> = {
    '1d': 86400,
    '3d': 259200,
    '7d': 604800,
    '14d': 1209600,
    '30d': 2592000,
  };
  return map[val] || 0;
}

export function ProposeDeal() {
  const { channelId } = useParams();
  const navigate = useNavigate();
  const [channel, setChannel] = useState<Channel | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState('');

  const [formData, setFormData] = useState({
    title: '',
    contentText: '',
    minViews: '',
    duration: '',
    deadline: '',
    amount: '',
  });

  useEffect(() => {
    if (!channelId) return;
    const id = parseInt(channelId, 10);
    if (Number.isNaN(id)) {
      setError('Invalid channel');
      setLoading(false);
      return;
    }
    apiClient(`/api/channels/${id}`)
      .then((d) => d.channel)
      .then(setChannel)
      .catch(() => setError('Failed to load'))
      .finally(() => setLoading(false));
  }, [channelId]);

  const defaultDeadline = () => {
    const d = new Date();
    d.setDate(d.getDate() + 14);
    return d.toISOString().slice(0, 16);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!channel || !channelId) return;
    setIsSubmitting(true);
    setSubmitError('');
    try {
      const amount = parseUnits(formData.amount || '0', MTK_TOKEN.decimals).toString();
      const { deal } = await apiClient(`/api/channels/${channelId}/propose`, {
        method: 'POST',
        body: JSON.stringify({
          title: formData.title,
          contentText: formData.contentText,
          tokenAddress: MTK_TOKEN.address,
          amount,
          minViews: parseInt(formData.minViews, 10) || 0,
          campaignDuration: parseDuration(formData.duration),
          expiryDeadline: formData.deadline || new Date(Date.now() + 14 * 86400000).toISOString(),
          categories: [],
        }),
      });
      navigate(`/deal/${deal.id}`);
    } catch (err: any) {
      console.error('Propose deal failed:', err);
      setSubmitError(err?.message || 'Failed to create proposal');
    } finally {
      setIsSubmitting(false);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-32">
        <Loader2 className="w-8 h-8 animate-spin text-accent-violet" />
      </div>
    );
  }

  if (!channel || error) {
    return (
      <div className="max-w-2xl mx-auto text-center py-20">
        <h2 className="mb-2">Channel not found</h2>
        <p className="text-text-secondary mb-6">{error || 'This channel does not exist.'}</p>
        <Link to="/marketplace">
          <GradientButton>Back to Marketplace</GradientButton>
        </Link>
      </div>
    );
  }

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <Link
        to="/marketplace"
        className="inline-flex items-center gap-1 text-sm text-text-secondary hover:text-text-primary transition-colors mb-2"
      >
        <ArrowLeft className="w-4 h-4" />
        Back to Marketplace
      </Link>

      <div className="bg-card border border-border-color rounded-xl p-6 mb-6">
        <div className="flex items-center gap-3 mb-2">
          <div className="w-12 h-12 rounded-full bg-sky-500/10 flex items-center justify-center">
            <Twitter className="w-6 h-6 text-sky-500" />
          </div>
          <div>
            <h2 className="font-semibold text-accent-violet">@{channel.platformHandle}</h2>
            <div className="flex items-center gap-2 text-sm text-text-muted">
              <Users className="w-4 h-4" />
              {channel.followerCount.toLocaleString()} followers · ${channel.pricePerPost}/post
            </div>
          </div>
        </div>
        <p className="text-text-secondary text-sm">
          Enter your ad content and criteria below. A deal proposal will be sent to @{channel.platformHandle}; they can accept and then you fund the escrow.
        </p>
      </div>

      <form onSubmit={handleSubmit} className="bg-card border border-border-color rounded-xl p-6 space-y-6">
        <h3 className="font-semibold text-lg">Proposal details</h3>

        <div>
          <label className="block mb-2 text-sm font-medium">Title *</label>
          <Input
            value={formData.title}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
            placeholder="e.g. Q1 Promo"
            className="bg-elevated border-border-color"
            required
          />
        </div>

        <div>
          <label className="block mb-2 text-sm font-medium">Tweet content *</label>
          <Textarea
            value={formData.contentText}
            onChange={(e) => setFormData({ ...formData, contentText: e.target.value })}
            placeholder="The exact text the influencer should post..."
            className="bg-background border-border-color font-mono text-sm"
            rows={4}
            required
          />
          <div className="text-xs text-text-muted mt-1 text-right">{formData.contentText.length}/280</div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label className="block mb-2 text-sm font-medium">Minimum views</label>
            <Input
              type="number"
              min={0}
              value={formData.minViews}
              onChange={(e) => setFormData({ ...formData, minViews: e.target.value })}
              placeholder="0"
              className="bg-elevated border-border-color"
            />
          </div>
          <div>
            <label className="block mb-2 text-sm font-medium">Post must stay up</label>
            <select
              value={formData.duration}
              onChange={(e) => setFormData({ ...formData, duration: e.target.value })}
              className="w-full px-4 h-11 rounded-lg bg-elevated border border-border-color text-text-primary"
            >
              <option value="">No requirement</option>
              <option value="1d">1 day</option>
              <option value="3d">3 days</option>
              <option value="7d">7 days</option>
              <option value="14d">14 days</option>
              <option value="30d">30 days</option>
            </select>
          </div>
        </div>

        <div>
          <label className="block mb-2 text-sm font-medium">Proposal deadline</label>
          <Input
            type="datetime-local"
            value={formData.deadline || defaultDeadline()}
            onChange={(e) => setFormData({ ...formData, deadline: e.target.value })}
            className="bg-elevated border-border-color"
          />
        </div>

        <div>
          <label className="block mb-2 text-sm font-medium">Amount ({MTK_TOKEN.symbol}) *</label>
          <Input
            type="number"
            min={0}
            step="any"
            value={formData.amount}
            onChange={(e) => setFormData({ ...formData, amount: e.target.value })}
            placeholder={channel.pricePerPost ? String(channel.pricePerPost / 100) : '0'}
            className="bg-elevated border-border-color"
            required
          />
        </div>

        <div className="p-4 bg-accent-blue/10 border border-accent-blue/20 rounded-lg flex gap-3">
          <Info className="w-5 h-5 text-accent-blue flex-shrink-0 mt-0.5" />
          <div className="text-sm text-text-secondary">
            After @{channel.platformHandle} accepts, you’ll deposit the amount into escrow on the deal page. Verification runs when they submit their post URL.
          </div>
        </div>

        {submitError && (
          <div className="p-4 bg-error/10 border border-error/20 rounded-lg text-sm text-error">
            {submitError}
          </div>
        )}

        <GradientButton type="submit" disabled={isSubmitting} loading={isSubmitting} className="w-full h-12">
          {isSubmitting ? 'Creating proposal...' : 'Send proposal'}
        </GradientButton>
      </form>
    </div>
  );
}
