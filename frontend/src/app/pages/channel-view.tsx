import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router';
import { Twitter, Users, DollarSign, Loader2, ArrowLeft, ShieldCheck } from 'lucide-react';
import { GradientButton } from '../components/gradient-button';
import { apiClient } from '../lib/api';

interface Channel {
  id: number;
  platformHandle: string;
  platform: string;
  followerCount: number;
  profileImageUrl?: string;
  categories: string[];
  pricePerPost: number;
  isActive: boolean;
  worldIdVerifiedAt?: string;
}

export function ChannelView() {
  const { id } = useParams();
  const [channel, setChannel] = useState<Channel | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    if (!id) return;
    apiClient(`/api/channels/${id}`)
      .then((data) => {
        setChannel(data.channel);
        setLoading(false);
      })
      .catch(() => {
        setError('Channel not found');
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

      <div className="bg-card border border-border-color rounded-xl p-8">
        <div className="flex flex-col sm:flex-row items-start gap-6">
          {channel.profileImageUrl ? (
            <img
              src={channel.profileImageUrl}
              alt={`@${channel.platformHandle}`}
              className="w-24 h-24 rounded-full object-cover border-2 border-border-color"
            />
          ) : (
            <div className="w-24 h-24 rounded-full bg-sky-500/10 flex items-center justify-center border-2 border-border-color">
              <Twitter className="w-12 h-12 text-sky-500" />
            </div>
          )}
          <div className="flex-1 min-w-0">
            <div className="flex items-center gap-2 mb-1">
              <span className="text-accent-violet font-semibold text-lg">@{channel.platformHandle}</span>
              {channel.worldIdVerifiedAt && (
                <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 text-xs font-medium">
                  <ShieldCheck className="w-3 h-3" />
                  Verified Human
                </span>
              )}
              {channel.platform && (
                <span className="text-xs text-text-muted capitalize">{channel.platform}</span>
              )}
            </div>
            {!channel.isActive && (
              <span className="inline-block px-2 py-0.5 rounded text-xs bg-text-muted/20 text-text-muted mb-2">
                Inactive
              </span>
            )}
            <div className="flex flex-wrap gap-4 mt-4">
              <div className="flex items-center gap-2">
                <Users className="w-5 h-5 text-text-muted" />
                <span className="font-semibold">{channel.followerCount.toLocaleString()} followers</span>
              </div>
              <div className="flex items-center gap-2">
                <DollarSign className="w-5 h-5 text-text-muted" />
                <span className="font-semibold">${channel.pricePerPost}/post</span>
              </div>
            </div>
          </div>
        </div>

        {channel.categories && channel.categories.length > 0 && (
          <>
            <div className="h-px bg-border-color my-6" />
            <div>
              <div className="text-sm text-text-muted mb-2">Categories</div>
              <div className="flex flex-wrap gap-2">
                {channel.categories.map((cat) => (
                  <span key={cat} className="px-3 py-1 bg-elevated rounded-full text-sm text-text-secondary">
                    {cat}
                  </span>
                ))}
              </div>
            </div>
          </>
        )}

        <div className="mt-8 flex flex-col sm:flex-row gap-3">
          <Link to={`/propose-deal/${channel.id}`} className="flex-1">
            <GradientButton className="w-full">Propose Deal</GradientButton>
          </Link>
          <Link to="/marketplace" className="flex-1">
            <button className="w-full px-4 py-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all font-medium text-sm">
              Back to Marketplace
            </button>
          </Link>
        </div>
      </div>
    </div>
  );
}
