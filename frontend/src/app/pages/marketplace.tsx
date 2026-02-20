import { useState, useEffect } from 'react';
import { Link } from 'react-router';
import { Twitter, Search, Eye, Clock, DollarSign, BadgeCheck, Users, Star } from 'lucide-react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../components/ui/tabs';
import { Input } from '../components/ui/input';
import { GradientButton } from '../components/gradient-button';
import { Jazzicon } from '../utils/jazzicon';
import { cn } from '../components/ui/utils';
import { formatUnits } from 'viem';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface MarketCampaign {
  id: number;
  title: string;
  contentText: string;
  categories: string[];
  amount: string;
  minViews: number;
  campaignDuration: number;
  expiryDeadline: string;
}

interface MarketChannel {
  channel: {
    id: number;
    platformHandle: string;
    followerCount: number;
    pricePerPost: number;
    categories: string[];
  };
  user: {
    walletAddress: string;
    displayName: string;
  };
}

export function Marketplace() {
  const [activeTab, setActiveTab] = useState('campaigns');
  const [campaigns, setCampaigns] = useState<MarketCampaign[]>([]);
  const [influencers, setInfluencers] = useState<MarketChannel[]>([]);
  const [search, setSearch] = useState('');

  useEffect(() => {
    apiClient(`/api/marketplace/campaigns${search ? `?search=${search}` : ''}`)
      .then((data) => setCampaigns(data.campaigns || []))
      .catch(() => {});
  }, [search]);

  useEffect(() => {
    apiClient('/api/marketplace/channels')
      .then((data) => setInfluencers(data.channels || []))
      .catch(() => {});
  }, []);

  const getDaysRemaining = (deadline: string) => {
    const diff = new Date(deadline).getTime() - Date.now();
    return Math.max(0, Math.ceil(diff / (1000 * 60 * 60 * 24)));
  };

  const getDeadlineBadgeColor = (days: number) => {
    if (days > 7) return 'bg-success/10 text-success border-success/20';
    if (days >= 1 && days <= 7) return 'bg-warning/10 text-warning border-warning/20';
    return 'bg-error/10 text-error border-error/20';
  };

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <h1>Marketplace</h1>
      </div>

      <Tabs value={activeTab} onValueChange={setActiveTab}>
        <TabsList className="bg-card border border-border-color">
          <TabsTrigger value="campaigns">Campaigns</TabsTrigger>
          <TabsTrigger value="influencers">Influencers</TabsTrigger>
        </TabsList>

        {/* Search and Filters */}
        <div className="flex flex-col sm:flex-row gap-4 mt-6">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-text-muted" />
            <Input
              placeholder={activeTab === 'campaigns' ? 'Search campaigns...' : 'Search influencers...'}
              className="pl-10 h-12 bg-card border-border-color"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
          </div>
          <select className="px-4 h-12 rounded-lg bg-card border border-border-color text-text-primary">
            <option>Newest First</option>
            <option>Highest Budget</option>
            <option>Ending Soon</option>
          </select>
        </div>

        <TabsContent value="campaigns" className="mt-6">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {campaigns.map((campaign) => (
              <div
                key={campaign.id}
                className="bg-card border border-border-color rounded-xl p-6 hover:-translate-y-0.5 hover:border-accent-violet/50 hover:shadow-lg hover:shadow-accent-violet/10 transition-all group"
              >
                {/* Header */}
                <div className="flex items-start justify-between mb-4">
                  <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center">
                    <Twitter className="w-5 h-5 text-sky-500" />
                  </div>
                  {(() => {
                    const days = getDaysRemaining(campaign.expiryDeadline);
                    return (
                      <span className={cn('px-2 py-1 rounded-lg text-xs font-medium border', getDeadlineBadgeColor(days))}>
                        {days > 7 ? `${days}d` : days < 1 ? '<1d' : `${days}d`}
                      </span>
                    );
                  })()}
                </div>

                {/* Title & Description */}
                <h3 className="font-semibold mb-2 line-clamp-2">{campaign.title}</h3>
                <p className="text-sm text-text-secondary mb-4 line-clamp-2">{campaign.contentText}</p>

                {/* Categories */}
                <div className="flex flex-wrap gap-2 mb-4">
                  {(campaign.categories || []).map((cat) => (
                    <span key={cat} className="px-2 py-1 bg-elevated rounded text-xs text-text-secondary">
                      {cat}
                    </span>
                  ))}
                </div>

                {/* Separator */}
                <div className="h-px bg-border-color mb-4" />

                {/* Metrics */}
                <div className="grid grid-cols-3 gap-2 mb-4">
                  <div>
                    <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
                      <DollarSign className="w-3 h-3" />
                      Budget
                    </div>
                    <div className="font-semibold text-sm">{parseFloat(formatUnits(BigInt(campaign.amount), MTK_TOKEN.decimals))} MTK</div>
                  </div>
                  <div>
                    <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
                      <Eye className="w-3 h-3" />
                      Min Views
                    </div>
                    <div className="font-semibold text-sm">{campaign.minViews ? `${(campaign.minViews / 1000).toFixed(0)}K` : '0'}</div>
                  </div>
                  <div>
                    <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
                      <Clock className="w-3 h-3" />
                      Duration
                    </div>
                    <div className="font-semibold text-sm">{campaign.campaignDuration ? `${Math.ceil(campaign.campaignDuration / 86400)} days` : '-'}</div>
                  </div>
                </div>

                {/* Actions */}
                <div className="flex gap-2">
                  <Link to={`/campaign/${campaign.id}`} className="flex-1">
                    <button className="w-full px-4 py-2 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all text-sm font-medium">
                      View Details
                    </button>
                  </Link>
                  <Link to={`/apply/${campaign.id}`} className="flex-1">
                    <button className="w-full px-4 py-2 rounded-lg bg-gradient-to-br from-accent-violet to-accent-blue text-white text-sm font-medium hover:shadow-lg hover:shadow-accent-violet/25 transition-all">
                      Apply
                    </button>
                  </Link>
                </div>
              </div>
            ))}
          </div>
        </TabsContent>

        <TabsContent value="influencers" className="mt-6">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {influencers.map((item) => (
              <div
                key={item.channel.id}
                className="bg-card border border-border-color rounded-xl p-6 hover:-translate-y-0.5 hover:border-accent-violet/50 hover:shadow-lg hover:shadow-accent-violet/10 transition-all"
              >
                {/* Header */}
                <div className="flex items-start gap-3 mb-4">
                  <Jazzicon address={item.user.walletAddress} size={40} />
                  <div className="flex-1 min-w-0">
                    <div className="flex items-center gap-1 mb-1">
                      <span className="text-accent-violet font-medium truncate">@{item.channel.platformHandle}</span>
                    </div>
                    <div className="font-semibold truncate">{item.user.displayName || item.channel.platformHandle}</div>
                  </div>
                </div>

                {/* Categories */}
                <div className="flex flex-wrap gap-2 mb-3">
                  {(item.channel.categories || []).map((cat) => (
                    <span key={cat} className="px-2 py-1 bg-elevated rounded text-xs text-text-secondary">
                      {cat}
                    </span>
                  ))}
                </div>

                {/* Separator */}
                <div className="h-px bg-border-color mb-4" />

                {/* Metrics */}
                <div className="grid grid-cols-2 gap-2 mb-4">
                  <div>
                    <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
                      <Users className="w-3 h-3" />
                      Followers
                    </div>
                    <div className="font-semibold text-sm">{item.channel.followerCount.toLocaleString()}</div>
                  </div>
                  <div>
                    <div className="text-xs text-text-muted mb-1 flex items-center gap-1">
                      <DollarSign className="w-3 h-3" />
                      Rate
                    </div>
                    <div className="font-semibold text-sm">${item.channel.pricePerPost}/post</div>
                  </div>
                </div>

                {/* Actions */}
                <div className="flex gap-2">
                  <Link to={`/channel/${item.channel.id}`} className="flex-1">
                    <button className="w-full px-4 py-2 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all text-sm font-medium">
                      View Profile
                    </button>
                  </Link>
                  <Link to={`/propose-deal/${item.channel.id}`} className="flex-1">
                    <button className="w-full px-4 py-2 rounded-lg bg-gradient-to-br from-accent-violet to-accent-blue text-white text-sm font-medium hover:shadow-lg hover:shadow-accent-violet/25 transition-all">
                      Propose Deal
                    </button>
                  </Link>
                </div>
              </div>
            ))}
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}
