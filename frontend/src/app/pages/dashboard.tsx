import { useState, useEffect } from 'react';
import { Link } from 'react-router';
import { Twitter, TrendingUp, AlertCircle, Target, CheckCircle2 } from 'lucide-react';
import { GradientButton } from '../components/gradient-button';
import { StatusBadge } from '../components/status-badge';
import { Jazzicon } from '../utils/jazzicon';
import { apiClient } from '../lib/api';

interface Deal {
  id: number;
  title?: string;
  campaignTitle?: string;
  counterparty?: string;
  proposedBy?: string;
  status: string;
  nextAction?: string;
  deadline?: string;
  expiryDeadline?: string;
}

function shortenAddress(addr: string | null | undefined): string {
  if (addr == null || typeof addr !== 'string') return '—';
  if (addr.length < 10) return addr;
  return `${addr.slice(0, 6)}...${addr.slice(-4)}`;
}

const defaultStats = [
  { label: 'Active Deals', value: '0', icon: Target, color: 'text-accent-violet' },
  { label: 'Pending Verifications', value: '0', icon: AlertCircle, color: 'text-warning' },
  { label: 'Total Earned', value: '$0', icon: TrendingUp, color: 'text-success' },
  { label: 'Success Rate', value: '0%', icon: CheckCircle2, color: 'text-accent-violet' },
];

export function Dashboard() {
  const [activeDeals, setActiveDeals] = useState<Deal[]>([]);
  const [stats, setStats] = useState(defaultStats);
  const [recentActivity] = useState([
    { type: 'info', text: 'Welcome to Chainfluence!', time: 'Just now', color: 'bg-accent-violet' },
  ]);

  useEffect(() => {
    apiClient('/api/deals')
      .then((data) => {
        const deals = data.deals || [];
        setActiveDeals(
          deals
            .filter((d: any) => !['completed', 'refunded', 'cancelled'].includes(d.status))
            .map((d: any) => ({
              ...d,
              title: d.campaignTitle ?? d.title,
              counterparty: d.proposedBy ?? d.counterparty,
              nextAction: d.nextAction ?? (d.status === 'verifying' ? 'Verification in progress' : d.status === 'posted' ? 'Ready to verify' : undefined),
              deadline: d.deadline ?? (d.expiryDeadline ? new Date(d.expiryDeadline).toLocaleDateString() : undefined),
            }))
        );
        const active = deals.filter((d: any) => !['completed', 'refunded', 'cancelled'].includes(d.status)).length;
        const verifying = deals.filter((d: any) => d.status === 'verifying').length;
        const completed = deals.filter((d: any) => d.status === 'completed').length;
        const total = deals.length;
        setStats([
          { label: 'Active Deals', value: String(active), icon: Target, color: 'text-accent-violet' },
          { label: 'Pending Verifications', value: String(verifying), icon: AlertCircle, color: 'text-warning' },
          { label: 'Total Earned', value: `${completed} deals`, icon: TrendingUp, color: 'text-success' },
          { label: 'Success Rate', value: total ? `${Math.round((completed / total) * 100)}%` : '0%', icon: CheckCircle2, color: 'text-accent-violet' },
        ]);
      })
      .catch(() => {});
  }, []);
  return (
    <div className="max-w-7xl mx-auto space-y-6">
      {/* Stats Row */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        {stats.map((stat) => {
          const Icon = stat.icon;
          return (
            <div
              key={stat.label}
              className="bg-card border border-border-color rounded-xl p-6 hover:-translate-y-0.5 transition-transform"
            >
              <div className="flex items-start justify-between mb-4">
                <div className={`p-2 rounded-lg bg-elevated ${stat.color}`}>
                  <Icon className="w-5 h-5" />
                </div>
              </div>
              <div className="text-3xl font-bold mb-1">{stat.value}</div>
              <div className="text-sm text-text-secondary">{stat.label}</div>
            </div>
          );
        })}
      </div>

      {/* Two column layout */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Active Deals - 2/3 width */}
        <div className="lg:col-span-2">
          <div className="bg-card border border-border-color rounded-xl p-6">
            <div className="flex items-center justify-between mb-6">
              <h2>Active Deals</h2>
              <Link to="/my-deals" className="text-sm text-accent-violet hover:text-accent-blue transition-colors">
                View All
              </Link>
            </div>

            <div className="space-y-3">
              {activeDeals.map((deal) => (
                <Link
                  key={deal.id}
                  to={`/deal/${deal.id}`}
                  className="block p-4 bg-elevated border border-border-color rounded-lg hover:border-accent-violet/50 hover:-translate-y-0.5 transition-all group"
                >
                  <div className="flex items-start gap-4">
                    {/* Twitter Icon */}
                    <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
                      <Twitter className="w-5 h-5 text-sky-500" />
                    </div>

                    {/* Content */}
                    <div className="flex-1 min-w-0">
                      <h3 className="font-semibold mb-1 truncate">{deal.title ?? 'Deal'}</h3>
                      <div className="flex items-center gap-2 mb-2">
                        {deal.counterparty && <Jazzicon address={deal.counterparty} size={16} />}
                        <span className="text-sm text-text-secondary font-mono truncate">
                          {shortenAddress(deal.counterparty ?? deal.proposedBy)}
                        </span>
                      </div>
                      <p className="text-sm text-text-muted">{deal.nextAction ?? ''}</p>
                    </div>

                    {/* Status & Deadline */}
                    <div className="flex flex-col items-end gap-2 flex-shrink-0">
                      <StatusBadge status={deal.status} />
                      {deal.deadline && <span className="text-xs text-warning">{deal.deadline}</span>}
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          </div>
        </div>

        {/* Recent Activity - 1/3 width */}
        <div className="lg:col-span-1">
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h2 className="mb-6">Recent Activity</h2>

            <div className="space-y-4">
              {recentActivity.map((activity, index) => (
                <div key={index} className="flex gap-3">
                  {/* Timeline dot */}
                  <div className="relative flex flex-col items-center">
                    <div className={`w-2 h-2 rounded-full ${activity.color} flex-shrink-0 mt-2`} />
                    {index < recentActivity.length - 1 && (
                      <div className="w-px h-full bg-border-color mt-1" />
                    )}
                  </div>

                  {/* Content */}
                  <div className="flex-1 pb-2">
                    <p className="text-sm mb-1">{activity.text}</p>
                    <p className="text-xs text-text-muted">{activity.time}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Quick Actions */}
      <div className="flex flex-wrap gap-4">
        <Link to="/create-campaign">
          <GradientButton>Create Campaign</GradientButton>
        </Link>
        <Link to="/marketplace">
          <button className="px-6 h-11 rounded-lg font-medium border border-border-color bg-transparent hover:border-accent-violet/50 transition-all">
            Browse Marketplace
          </button>
        </Link>
      </div>
    </div>
  );
}
