import { useState, useEffect } from 'react';
import { Link } from 'react-router';
import { Twitter, Clock, DollarSign } from 'lucide-react';
import { formatUnits } from 'viem';
import { StatusBadge } from '../components/status-badge';
import { Jazzicon } from '../utils/jazzicon';
import { apiClient } from '../lib/api';
import { MTK_TOKEN } from '../lib/wagmi';

interface Deal {
  id: number;
  status: string;
  proposedBy: string;
  postUrl?: string;
  createdAt: string;
  campaignTitle?: string;
  campaignAmount?: string;
}

export function MyDeals() {
  const [deals, setDeals] = useState<Deal[]>([]);

  useEffect(() => {
    apiClient('/api/deals')
      .then((data) => setDeals(data.deals || []))
      .catch(() => {});
  }, []);

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      <h1>My Deals</h1>

      {deals.length === 0 && (
        <div className="text-center py-16 text-text-secondary">
          <p className="mb-2">No deals yet</p>
          <p className="text-sm text-text-muted">Browse the marketplace to find campaigns to apply to.</p>
        </div>
      )}

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {deals.map((deal) => {
          const amount = deal.campaignAmount
            ? parseFloat(formatUnits(BigInt(deal.campaignAmount), MTK_TOKEN.decimals))
            : null;

          return (
            <Link
              key={deal.id}
              to={`/deal/${deal.id}`}
              className="block bg-card border border-border-color rounded-xl p-6 hover:border-accent-violet/50 hover:-translate-y-0.5 transition-all"
            >
              <div className="flex items-start gap-4">
                <div className="w-12 h-12 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
                  <Twitter className="w-6 h-6 text-sky-500" />
                </div>

                <div className="flex-1 min-w-0">
                  <h3 className="font-semibold mb-1 truncate">
                    {deal.campaignTitle || `Deal #${deal.id}`}
                  </h3>
                  <p className="text-xs text-text-muted mb-2">Deal #{deal.id}</p>

                  <div className="flex items-center gap-2 mb-3">
                    <Jazzicon address={deal.proposedBy} size={20} />
                    <span className="text-sm text-text-secondary font-mono truncate">
                      {deal.proposedBy.slice(0, 6)}...{deal.proposedBy.slice(-4)}
                    </span>
                  </div>

                  <div className="flex items-center justify-between">
                    {amount != null && (
                      <div className="flex items-center gap-1 text-sm font-semibold">
                        <DollarSign className="w-3 h-3 text-text-muted" />
                        {amount} {MTK_TOKEN.symbol}
                      </div>
                    )}
                    <StatusBadge status={deal.status as any} />
                  </div>
                </div>
              </div>

              <div className="mt-4 pt-4 border-t border-border-color flex items-center justify-between text-sm">
                <span className="text-text-muted flex items-center gap-2">
                  <Clock className="w-4 h-4" />
                  Created
                </span>
                <span className="text-text-secondary">{new Date(deal.createdAt).toLocaleDateString()}</span>
              </div>
            </Link>
          );
        })}
      </div>
    </div>
  );
}
