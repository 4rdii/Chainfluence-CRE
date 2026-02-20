import { useState, useEffect } from 'react';
import { useSearchParams } from 'react-router';
import { Twitter, Plus, Edit2, Trash2, Radio as RadioIcon, Loader2, CheckCircle, BadgeCheck } from 'lucide-react';
import { GradientButton } from '../components/gradient-button';
import { Switch } from '../components/ui/switch';
import { apiClient } from '../lib/api';

interface Channel {
  id: number;
  platformHandle: string;
  platformUserId?: string;
  profileImageUrl?: string;
  followerCount: number;
  pricePerPost: number;
  categories: string[];
  isActive: boolean;
  verifiedAt?: string;
}

export function MyChannels() {
  const [channelsList, setChannelsList] = useState<Channel[]>([]);
  const [connecting, setConnecting] = useState(false);
  const [searchParams, setSearchParams] = useSearchParams();

  const loadChannels = () => {
    apiClient('/api/channels')
      .then((data) => setChannelsList(data.channels || []))
      .catch(() => {});
  };

  useEffect(() => {
    loadChannels();
  }, []);

  // Handle OAuth callback params
  useEffect(() => {
    if (searchParams.get('connected') === 'true') {
      loadChannels();
      // Clear the query param
      setSearchParams({}, { replace: true });
    }
    if (searchParams.get('error')) {
      console.error('X OAuth error:', searchParams.get('error'));
      setSearchParams({}, { replace: true });
    }
  }, [searchParams]);

  const handleRegister = async () => {
    setConnecting(true);
    try {
      const { url } = await apiClient('/api/auth/x/connect');
      if (url) {
        window.location.href = url;
      }
    } catch (err) {
      console.error('Failed to start X OAuth:', err);
      setConnecting(false);
    }
  };

  const toggleActive = (id: number) => {
    const channel = channelsList.find(ch => ch.id === id);
    if (!channel) return;
    const newActive = !channel.isActive;
    setChannelsList(channelsList.map(ch =>
      ch.id === id ? { ...ch, isActive: newActive } : ch
    ));
    apiClient(`/api/channels/${id}`, {
      method: 'PATCH',
      body: JSON.stringify({ isActive: newActive }),
    }).catch(() => {});
  };

  const handleDelete = async (id: number) => {
    if (!confirm('Remove this channel?')) return;
    await apiClient(`/api/channels/${id}`, { method: 'DELETE' }).catch(() => {});
    setChannelsList(channelsList.filter(ch => ch.id !== id));
  };

  const EmptyState = () => (
    <div className="flex flex-col items-center justify-center py-16 px-4 bg-card border border-border-color rounded-xl">
      <div className="w-32 h-32 mb-6 flex items-center justify-center">
        <RadioIcon className="w-24 h-24 text-accent-violet/20" strokeWidth={1} />
      </div>
      <h3 className="mb-2">No channels registered</h3>
      <p className="text-text-secondary mb-6 text-center max-w-md">
        Connect your X (Twitter) account to start receiving campaign proposals
      </p>
      <GradientButton onClick={handleRegister} disabled={connecting}>
        {connecting ? (
          <span className="flex items-center gap-2">
            <Loader2 className="w-4 h-4 animate-spin" />
            Connecting...
          </span>
        ) : (
          <span className="flex items-center gap-2">
            <Twitter className="w-4 h-4" />
            Connect X Account
          </span>
        )}
      </GradientButton>
    </div>
  );

  if (channelsList.length === 0) {
    return (
      <div className="max-w-7xl mx-auto space-y-6">
        <div className="flex items-center justify-between">
          <h1>My Channels</h1>
        </div>
        <EmptyState />
      </div>
    );
  }

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <h1>My Channels</h1>
        <GradientButton onClick={handleRegister} disabled={connecting} icon={<Plus className="w-4 h-4" />}>
          {connecting ? 'Connecting...' : 'Add Channel'}
        </GradientButton>
      </div>

      <div className="space-y-4">
        {channelsList.map((channel) => (
          <div
            key={channel.id}
            className="bg-card border border-border-color rounded-xl p-6 hover:border-accent-violet/30 transition-all"
          >
            <div className="flex flex-col lg:flex-row lg:items-center gap-6">
              {/* Left: Channel Info */}
              <div className="flex items-center gap-4 flex-1">
                {/* Avatar or icon */}
                {channel.profileImageUrl ? (
                  <img
                    src={channel.profileImageUrl}
                    alt={`@${channel.platformHandle}`}
                    className="w-12 h-12 rounded-full flex-shrink-0"
                  />
                ) : (
                  <div className="w-12 h-12 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
                    <Twitter className="w-6 h-6 text-sky-500" />
                  </div>
                )}

                {/* Details */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 mb-1">
                    <span className="text-accent-violet font-semibold">@{channel.platformHandle}</span>
                    {channel.verifiedAt && (
                      <BadgeCheck className="w-4 h-4 text-sky-500" />
                    )}
                  </div>
                  <div className="flex items-center gap-4 text-sm text-text-secondary">
                    <span>{channel.followerCount.toLocaleString()} followers</span>
                    <span>•</span>
                    <span>${channel.pricePerPost}/post</span>
                  </div>
                </div>
              </div>

              {/* Center: Categories */}
              <div className="flex flex-wrap gap-2 lg:min-w-[200px]">
                {channel.categories.map((cat) => (
                  <span
                    key={cat}
                    className="px-3 py-1 bg-elevated rounded-full text-sm text-text-secondary"
                  >
                    {cat}
                  </span>
                ))}
              </div>

              {/* Right: Actions */}
              <div className="flex items-center gap-4 flex-shrink-0">
                {/* Active Toggle */}
                <div className="flex items-center gap-2">
                  <span className="text-sm text-text-secondary">
                    {channel.isActive ? 'Active' : 'Inactive'}
                  </span>
                  <Switch
                    checked={channel.isActive}
                    onCheckedChange={() => toggleActive(channel.id)}
                  />
                </div>

                {/* Edit */}
                <button className="p-2 hover:bg-elevated rounded-lg transition-colors text-text-secondary hover:text-text-primary">
                  <Edit2 className="w-4 h-4" />
                </button>

                {/* Delete */}
                <button
                  onClick={() => handleDelete(channel.id)}
                  className="p-2 hover:bg-error/10 rounded-lg transition-colors text-text-secondary hover:text-error"
                >
                  <Trash2 className="w-4 h-4" />
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
