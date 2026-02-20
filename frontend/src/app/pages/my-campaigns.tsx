import { useState, useEffect } from 'react';
import { Link } from 'react-router';
import { Megaphone, Plus } from 'lucide-react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '../components/ui/tabs';
import { StatusBadge } from '../components/status-badge';
import { GradientButton } from '../components/gradient-button';
import { formatUnits } from 'viem';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface Campaign {
  id: number;
  title: string;
  amount: string;
  status: string;
  expiryDeadline: string;
  dealCount?: number;
}

export function MyCampaigns() {
  const [activeTab, setActiveTab] = useState('all');
  const [campaigns, setCampaigns] = useState<Campaign[]>([]);

  useEffect(() => {
    apiClient('/api/campaigns')
      .then((data) => setCampaigns(data.campaigns || []))
      .catch(() => {});
  }, []);

  const filteredCampaigns = campaigns.filter(c => {
    if (activeTab === 'all') return true;
    if (activeTab === 'active') return !['completed', 'refunded', 'cancelled', 'draft'].includes(c.status);
    return c.status === activeTab;
  });

  const EmptyState = () => (
    <div className="flex flex-col items-center justify-center py-16 px-4">
      <div className="w-32 h-32 mb-6 flex items-center justify-center">
        <Megaphone className="w-24 h-24 text-accent-violet/20" strokeWidth={1} />
      </div>
      <h3 className="mb-2">No campaigns yet</h3>
      <p className="text-text-secondary mb-6 text-center max-w-md">
        Create your first campaign to start working with influencers
      </p>
      <Link to="/create-campaign">
        <GradientButton>Create Campaign</GradientButton>
      </Link>
    </div>
  );

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      <div className="flex items-center justify-between">
        <h1>My Campaigns</h1>
        <Link to="/create-campaign">
          <GradientButton icon={<Plus className="w-4 h-4" />}>
            Create Campaign
          </GradientButton>
        </Link>
      </div>

      <Tabs value={activeTab} onValueChange={setActiveTab}>
        <TabsList className="bg-card border border-border-color">
          <TabsTrigger value="active">Active</TabsTrigger>
          <TabsTrigger value="completed">Completed</TabsTrigger>
          <TabsTrigger value="draft">Draft</TabsTrigger>
          <TabsTrigger value="all">All</TabsTrigger>
        </TabsList>

        <TabsContent value={activeTab} className="mt-6">
          {filteredCampaigns.length === 0 ? (
            <EmptyState />
          ) : (
            <div className="bg-card border border-border-color rounded-xl overflow-hidden">
              {/* Table header */}
              <div className="hidden md:grid md:grid-cols-12 gap-4 px-6 py-4 bg-elevated border-b border-border-color text-sm text-text-muted font-medium">
                <div className="col-span-4">Title</div>
                <div className="col-span-2">Budget</div>
                <div className="col-span-2">Deals</div>
                <div className="col-span-2">Status</div>
                <div className="col-span-1">Deadline</div>
                <div className="col-span-1"></div>
              </div>

              {/* Table rows */}
              <div className="divide-y divide-border-color">
                {filteredCampaigns.map((campaign, index) => (
                  <div
                    key={campaign.id}
                    className={cn(
                      'grid grid-cols-1 md:grid-cols-12 gap-4 px-6 py-4 hover:bg-elevated/50 transition-colors',
                      index % 2 === 0 ? 'bg-card' : 'bg-elevated/30'
                    )}
                  >
                    {/* Title */}
                    <div className="col-span-1 md:col-span-4">
                      <div className="text-sm text-text-muted md:hidden mb-1">Title</div>
                      <div className="font-semibold">{campaign.title}</div>
                    </div>

                    {/* Budget */}
                    <div className="col-span-1 md:col-span-2">
                      <div className="text-sm text-text-muted md:hidden mb-1">Budget</div>
                      <div className="font-medium">{parseFloat(formatUnits(BigInt(campaign.amount), MTK_TOKEN.decimals))} MTK</div>
                    </div>

                    {/* Deals */}
                    <div className="col-span-1 md:col-span-2">
                      <div className="text-sm text-text-muted md:hidden mb-1">Deals</div>
                      <div className="font-medium">{campaign.dealCount ?? 0}</div>
                    </div>

                    {/* Status */}
                    <div className="col-span-1 md:col-span-2">
                      <div className="text-sm text-text-muted md:hidden mb-1">Status</div>
                      <StatusBadge status={campaign.status} />
                    </div>

                    {/* Deadline */}
                    <div className="col-span-1 md:col-span-1">
                      <div className="text-sm text-text-muted md:hidden mb-1">Deadline</div>
                      <div className="text-sm">{campaign.expiryDeadline ? new Date(campaign.expiryDeadline).toLocaleDateString() : '-'}</div>
                    </div>

                    {/* Actions */}
                    <div className="col-span-1 md:col-span-1 flex justify-end">
                      <Link
                        to={`/campaign/${campaign.id}`}
                        className="text-accent-violet hover:text-accent-blue transition-colors text-sm font-medium"
                      >
                        View
                      </Link>
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}
        </TabsContent>
      </Tabs>
    </div>
  );
}

function cn(...classes: (string | boolean | undefined)[]) {
  return classes.filter(Boolean).join(' ');
}
