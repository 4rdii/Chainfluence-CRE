import { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router';
import { Twitter, Copy, ExternalLink, Check, X, Loader2, AlertTriangle, Link as LinkIcon, Eye, Clock, Calendar, DollarSign } from 'lucide-react';
import { useWriteContract, useAccount } from 'wagmi';
import { readContract, waitForTransactionReceipt } from '@wagmi/core';
import { formatUnits, parseUnits, keccak256, toBytes } from 'viem';
import { StatusBadge } from '../components/status-badge';
import { Jazzicon } from '../utils/jazzicon';
import { GradientButton } from '../components/gradient-button';
import { Input } from '../components/ui/input';
import { cn } from '../components/ui/utils';
import { apiClient } from '../lib/api';
import { escrowAbi, erc20Abi } from '../lib/escrow-abi';
import { ESCROW_ADDRESS, MTK_TOKEN, wagmiConfig } from '../lib/wagmi';

type DealStatus = 'proposed' | 'accepted' | 'funded' | 'posted' | 'verifying' | 'completed' | 'disputed';

const statusSteps: DealStatus[] = ['proposed', 'accepted', 'funded', 'posted', 'verifying', 'completed'];

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
  advertiserId: number;
  advertiserWalletAddress?: string;
  txHash?: string;
  onchainCampaignId?: number;
  isDirectDeal?: boolean;
}

interface Channel {
  id: number;
  platformHandle: string;
  followerCount: number;
  profileImageUrl?: string;
}

interface Deal {
  id: number;
  campaignId: number;
  channelId?: number;
  influencerId: number;
  status: string;
  proposedBy: string;
  postUrl?: string;
  onchainCampaignId?: number;
  acceptedAt?: string;
  fundedAt?: string;
  postedAt?: string;
  completedAt?: string;
  createdAt: string;
  updatedAt: string;
}

interface Verification {
  id: number;
  dealId: number;
  status: string;
  viewsChecked?: number;
  contentMatched?: boolean;
  durationChecked?: boolean;
  isEdited?: boolean;
  action?: string;
  message?: string;
  createdAt: string;
}

function formatDuration(seconds: number): string {
  if (!seconds) return 'No requirement';
  const days = Math.floor(seconds / 86400);
  if (days >= 1) return `${days} day${days > 1 ? 's' : ''}`;
  const hours = Math.floor(seconds / 3600);
  return `${hours} hour${hours > 1 ? 's' : ''}`;
}

function shortenAddress(address: string): string {
  if (!address || address.length < 10) return address || '—';
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
}

function timeAgo(dateStr: string): string {
  const now = new Date();
  const date = new Date(dateStr);
  const diffMs = now.getTime() - date.getTime();
  const minutes = Math.floor(diffMs / 60000);
  if (minutes < 60) return `${minutes}m ago`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}h ago`;
  const days = Math.floor(hours / 24);
  return `${days}d ago`;
}

function buildTimeline(
  deal: Deal,
  isAdvertiser: boolean,
  isInfluencer: boolean
): { time: string; text: string; color: string }[] {
  const events: { time: string; text: string; color: string }[] = [];
  if (deal.completedAt) {
    events.push({ time: timeAgo(deal.completedAt), text: 'Campaign completed — funds released', color: 'bg-success' });
  }
  if (deal.status === 'verifying') {
    events.push({ time: timeAgo(deal.updatedAt), text: 'Chainlink verification initiated', color: 'bg-accent-violet' });
  }
  if (deal.postedAt && isAdvertiser) {
    events.push({ time: timeAgo(deal.postedAt), text: 'Tweet URL submitted by influencer', color: 'bg-warning' });
  }
  if (deal.fundedAt && isInfluencer) {
    events.push({ time: timeAgo(deal.fundedAt), text: 'Campaign funded by advertiser', color: 'bg-accent-blue' });
  }
  if (deal.acceptedAt) {
    events.push({ time: timeAgo(deal.acceptedAt), text: 'Proposal accepted', color: 'bg-success' });
  }
  if (deal.status === 'disputed') {
    events.push({ time: timeAgo(deal.updatedAt), text: 'Deal disputed', color: 'bg-error' });
  }
  events.push({ time: timeAgo(deal.createdAt), text: 'Deal proposed', color: 'bg-text-muted' });
  return events;
}

export function DealView() {
  const { id } = useParams();
  const [currentStatus, setCurrentStatus] = useState<DealStatus>('proposed');
  const [tweetUrl, setTweetUrl] = useState('');
  const [isVerifying, setIsVerifying] = useState(false);
  const [deal, setDeal] = useState<Deal | null>(null);
  const [campaign, setCampaign] = useState<Campaign | null>(null);
  const [channel, setChannel] = useState<Channel | null>(null);
  const [verifications, setVerifications] = useState<Verification[]>([]);
  const [loading, setLoading] = useState(true);
  const [fundingError, setFundingError] = useState('');
  const [isFunding, setIsFunding] = useState(false);
  const [submitError, setSubmitError] = useState('');
  const [isSubmitting, setIsSubmitting] = useState(false);
  const { address } = useAccount();
  const { writeContractAsync } = useWriteContract();

  const [isAdvertiserFromApi, setIsAdvertiserFromApi] = useState(false);
  const [isInfluencerFromApi, setIsInfluencerFromApi] = useState(false);
  const [advertiserWalletFromApi, setAdvertiserWalletFromApi] = useState<string | undefined>(undefined);
  const [influencerWalletFromApi, setInfluencerWalletFromApi] = useState<string | undefined>(undefined);
  const [chainTxHashes, setChainTxHashes] = useState<{ acceptDealTx?: string; settlementTx?: string }>({});

  useEffect(() => {
    if (!id) return;
    apiClient(`/api/deals/${id}`)
      .then(async (data) => {
        const d = data.deal as Deal;
        setDeal(d);
        setCurrentStatus((d.status || 'proposed') as DealStatus);
        setTweetUrl(d.postUrl || '');
        // Server tells us role so Accept/Deposit show correctly for the advertiser
        setIsAdvertiserFromApi(Boolean((data as { isAdvertiser?: boolean }).isAdvertiser));
        setIsInfluencerFromApi(Boolean((data as { isInfluencer?: boolean }).isInfluencer));
        setAdvertiserWalletFromApi((data as { advertiserWalletAddress?: string }).advertiserWalletAddress);
        setInfluencerWalletFromApi((data as { influencerWalletAddress?: string }).influencerWalletAddress);
        setChainTxHashes((data as { chainTxHashes?: { acceptDealTx?: string; settlementTx?: string } }).chainTxHashes || {});

        const promises: Promise<any>[] = [
          apiClient(`/api/campaigns/${d.campaignId}`),
        ];
        if (d.channelId) {
          promises.push(apiClient(`/api/channels/${d.channelId}`));
        }
        // Fetch verifications
        promises.push(apiClient(`/api/deals/${id}/verifications`));

        const results = await Promise.allSettled(promises);

        if (results[0].status === 'fulfilled') {
          setCampaign(results[0].value.campaign);
        }
        if (d.channelId && results[1]?.status === 'fulfilled') {
          setChannel(results[1].value.channel);
        }
        const vIdx = d.channelId ? 2 : 1;
        if (results[vIdx]?.status === 'fulfilled') {
          setVerifications(results[vIdx].value.verifications || []);
        }
        setLoading(false);
      })
      .catch(() => setLoading(false));
  }, [id]);

  // Poll deal status while verifying so UI updates when webhook sets completed/refunded
  useEffect(() => {
    if (!id || currentStatus !== 'verifying') return;
    const interval = setInterval(async () => {
      try {
        const data = await apiClient(`/api/deals/${id}`);
        const d = data.deal as Deal;
        setDeal(d);
        const status = (d?.status || 'verifying') as DealStatus;
        setCurrentStatus(status);
        if (status !== 'verifying') {
          const v = await apiClient(`/api/deals/${id}/verifications`).catch(() => ({}));
          setVerifications((v as { verifications?: Verification[] }).verifications || []);
        }
      } catch (_) {}
    }, 15000);
    return () => clearInterval(interval);
  }, [id, currentStatus]);

  const getCurrentStepIndex = () => statusSteps.indexOf(currentStatus);

  const handleRefreshStatus = async () => {
    if (!id) return;
    try {
      const data = await apiClient(`/api/deals/${id}`);
      const d = data.deal as Deal;
      setDeal(d);
      setCurrentStatus((d?.status || currentStatus) as DealStatus);
      const v = await apiClient(`/api/deals/${id}/verifications`).catch(() => ({}));
      setVerifications((v as { verifications?: Verification[] }).verifications || []);
    } catch (_) {}
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
  };

  const handleAccept = async () => {
    try {
      await apiClient(`/api/deals/${id}/accept`, { method: 'POST' });
      setCurrentStatus('accepted');
      setDeal((prev) => prev ? { ...prev, status: 'accepted', acceptedAt: new Date().toISOString() } : prev);
    } catch (err) {
      console.error('Accept failed:', err);
    }
  };

  const handleFund = async () => {
    if (!campaign || !address) return;
    setIsFunding(true);
    setFundingError('');
    try {
      const amount = BigInt(campaign.amount);

      // Check allowance, approve if needed
      const currentAllowance = await readContract(wagmiConfig, {
        address: MTK_TOKEN.address,
        abi: erc20Abi,
        functionName: 'allowance',
        args: [address, ESCROW_ADDRESS],
      });
      if ((currentAllowance as bigint) < amount) {
        await writeContractAsync({
          address: MTK_TOKEN.address,
          abi: erc20Abi,
          functionName: 'approve',
          args: [ESCROW_ADDRESS, amount],
        });
      }

      // Compute content hash
      const normalized = campaign.contentText.trim().toLowerCase().replace(/\s+/g, ' ');
      const contentHash = keccak256(toBytes(normalized));

      // dealId = backend deal id so contract and backend stay in sync
      const dealId = BigInt(deal.id);
      const deadlineUnix = BigInt(Math.floor(new Date(campaign.expiryDeadline).getTime() / 1000));
      const durationSeconds = BigInt(campaign.campaignDuration || 0);

      const influencerAddress = (influencerWalletFromApi || '') as `0x${string}`
      if (!influencerAddress) throw new Error('Deal has no influencer address')

      const hash = await writeContractAsync({
        address: ESCROW_ADDRESS,
        abi: escrowAbi,
        functionName: 'deposit',
        args: [{
          dealId,
          token: MTK_TOKEN.address,
          influencer: influencerAddress,
          amount,
          contentHash,
          minViews: BigInt(campaign.minViews || 0),
          expiryDeadline: deadlineUnix,
          campaignDuration: durationSeconds,
          channelName: channel?.platformHandle || '',
        }],
      });

      // Wait for receipt
      const receipt = await waitForTransactionReceipt(wagmiConfig, { hash });

      if (receipt.status === 'reverted') {
        await apiClient(`/api/campaigns/${campaign.id}`, {
          method: 'PATCH',
          body: JSON.stringify({ txHash: hash, status: 'failed' }),
        });
        throw new Error('Transaction reverted on-chain.');
      }

      await apiClient(`/api/campaigns/${campaign.id}`, {
        method: 'PATCH',
        body: JSON.stringify({ txHash: hash, status: 'funded' }),
      });

      // Store deal id as on-chain id (contract uses same dealId = deal.id)
      const onchainDealId = deal.id;

      await apiClient(`/api/deals/${id}/fund`, {
        method: 'POST',
        body: JSON.stringify({ onchainCampaignId: onchainDealId }),
      });

      setCampaign((prev) => prev ? { ...prev, txHash: hash, status: 'funded' } : prev);
      setCurrentStatus('funded');
      setDeal((prev) => prev ? { ...prev, status: 'funded', fundedAt: new Date().toISOString(), onchainCampaignId: onchainDealId } : prev);
    } catch (err: any) {
      console.error('Fund failed:', err);
      setFundingError(err?.shortMessage || err?.message || 'Transaction failed');
    } finally {
      setIsFunding(false);
    }
  };

  const handleSubmitTweet = async () => {
    if (!tweetUrl || !deal?.onchainCampaignId) return;
    setIsSubmitting(true);
    setSubmitError('');
    try {
      // 1. Call acceptDeal(dealId, tweetUrl) on-chain — stores tweet URL in contract and accepts the deal
      const txHash = await writeContractAsync({
        address: ESCROW_ADDRESS,
        abi: escrowAbi,
        functionName: 'acceptDeal',
        args: [BigInt(deal.onchainCampaignId), tweetUrl],
      });
      await waitForTransactionReceipt(wagmiConfig, { hash: txHash });

      // 2. Notify backend
      await apiClient(`/api/deals/${id}/post`, {
        method: 'POST',
        body: JSON.stringify({ postUrl: tweetUrl }),
      });
      setCurrentStatus('posted');
      setDeal((prev) => prev ? { ...prev, status: 'posted', postUrl: tweetUrl, postedAt: new Date().toISOString() } : prev);
    } catch (err: any) {
      console.error('Submit tweet failed:', err);
      setSubmitError(err?.shortMessage || err?.message || 'Submission failed');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleVerify = async () => {
    setIsVerifying(true);
    try {
      await apiClient(`/api/deals/${id}/verify`, { method: 'POST' });
      setCurrentStatus('verifying');
      setTimeout(() => {
        apiClient(`/api/deals/${id}`).then((data) => {
          setCurrentStatus((data.deal?.status || 'verifying') as DealStatus);
          setDeal(data.deal);
          setIsVerifying(false);
        });
        apiClient(`/api/deals/${id}/verifications`).then((data) => {
          setVerifications(data.verifications || []);
        });
      }, 5000);
    } catch (err) {
      console.error('Verify failed:', err);
      setIsVerifying(false);
    }
  };

  const handleDispute = async () => {
    try {
      if (deal?.onchainCampaignId) {
        await writeContractAsync({
          address: ESCROW_ADDRESS,
          abi: escrowAbi,
          functionName: 'raiseDispute',
          args: [BigInt(deal.onchainCampaignId)],
        });
      }
      await apiClient(`/api/deals/${id}/dispute`, { method: 'POST' });
      setCurrentStatus('disputed');
    } catch (err) {
      console.error('Dispute failed:', err);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center py-32">
        <Loader2 className="w-8 h-8 animate-spin text-accent-violet" />
      </div>
    );
  }

  if (!deal) {
    return (
      <div className="max-w-2xl mx-auto text-center py-20">
        <h2 className="mb-2">Deal not found</h2>
        <p className="text-text-secondary mb-6">This deal does not exist.</p>
        <Link to="/my-deals">
          <GradientButton>Back to My Deals</GradientButton>
        </Link>
      </div>
    );
  }

  const amount = campaign
    ? parseFloat(formatUnits(BigInt(campaign.amount), MTK_TOKEN.decimals)).toString()
    : '—';
  const deadline = campaign ? new Date(campaign.expiryDeadline) : null;
  // Use server role flags so advertiser always sees Accept/Deposit; fallback to wallet match if API didn't return flags
  const isAdvertiser = isAdvertiserFromApi || Boolean(
    address && campaign?.advertiserWalletAddress && address.toLowerCase() === campaign.advertiserWalletAddress.toLowerCase()
  );
  const isInfluencer = isInfluencerFromApi || Boolean(
    address && deal.proposedBy && address.toLowerCase() === deal.proposedBy.toLowerCase()
  );
  // Who sent the proposal (proposedBy wallet). The other party can Accept/Decline.
  const isProposer = Boolean(
    address && deal.proposedBy && address.toLowerCase() === deal.proposedBy.toLowerCase()
  );
  const timeline = buildTimeline(deal, isAdvertiser, isInfluencer);
  // Advertiser = campaign owner; never use deal.proposedBy (that's the proposer's wallet)
  const advertiserAddress = advertiserWalletFromApi ?? campaign?.advertiserWalletAddress ?? undefined;

  return (
    <div className="max-w-7xl mx-auto space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <Link
            to="/my-deals"
            className="inline-flex items-center gap-1 text-sm text-text-secondary hover:text-text-primary transition-colors mb-2"
          >
            ← Back to My Deals
          </Link>
          <h1>Deal #{id}</h1>
          <p className="text-text-secondary mt-1">{campaign?.title || 'Loading...'}</p>
        </div>
        <StatusBadge status={currentStatus} className="text-base px-4 py-2" />
      </div>

      {/* Status Tracker */}
      <div className="bg-card border border-border-color rounded-xl p-6">
        <div className="flex items-center justify-between">
          {statusSteps.map((step, index) => {
            const isCurrent = step === currentStatus;
            const isPast = index < getCurrentStepIndex();
            const isDisputed = currentStatus === 'disputed';

            return (
              <div key={step} className="flex items-center flex-1">
                <div className="flex flex-col items-center flex-1">
                  <div
                    className={cn(
                      'w-12 h-12 rounded-full flex items-center justify-center border-2 transition-all',
                      isDisputed
                        ? 'bg-error/10 border-error'
                        : isPast
                        ? 'bg-success border-success'
                        : isCurrent
                        ? 'bg-accent-violet border-accent-violet'
                        : 'bg-elevated border-border-color'
                    )}
                  >
                    {isDisputed ? (
                      <X className="w-6 h-6 text-error" />
                    ) : isPast ? (
                      <Check className="w-6 h-6 text-white" />
                    ) : isCurrent ? (
                      <div className="w-3 h-3 rounded-full bg-white animate-pulse" />
                    ) : (
                      <div className="w-3 h-3 rounded-full bg-text-muted" />
                    )}
                  </div>
                  <span
                    className={cn(
                      'text-sm mt-3 text-center capitalize',
                      isCurrent ? 'text-text-primary font-semibold' : 'text-text-muted'
                    )}
                  >
                    {step}
                  </span>
                </div>
                {index < statusSteps.length - 1 && (
                  <div
                    className={cn(
                      'h-0.5 flex-1 -mt-8',
                      isPast ? 'bg-success' : 'bg-border-color'
                    )}
                  />
                )}
              </div>
            );
          })}
        </div>
      </div>

      {/* Main Content */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        {/* Left Column - 2/3 */}
        <div className="lg:col-span-2 space-y-6">
          {/* Campaign Details */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <div className="flex items-center gap-3 mb-6">
              <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center">
                <Twitter className="w-5 h-5 text-sky-500" />
              </div>
              <h2>Campaign Details</h2>
            </div>

            <div className="space-y-4">
              <div>
                <label className="text-sm text-text-muted">Tweet Content</label>
                <div className="mt-2 bg-background border border-border-color rounded-lg p-4 font-mono text-sm whitespace-pre-wrap break-words">
                  {campaign?.contentText || '—'}
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div>
                  <label className="text-sm text-text-muted flex items-center gap-1">
                    <Eye className="w-3 h-3" /> Minimum Views
                  </label>
                  <div className="font-semibold mt-1">
                    {campaign?.minViews ? campaign.minViews.toLocaleString() : 'None'}
                  </div>
                </div>
                <div>
                  <label className="text-sm text-text-muted flex items-center gap-1">
                    <Clock className="w-3 h-3" /> Duration
                  </label>
                  <div className="font-semibold mt-1">
                    {campaign ? formatDuration(campaign.campaignDuration) : '—'}
                  </div>
                </div>
                <div>
                  <label className="text-sm text-text-muted flex items-center gap-1">
                    <Calendar className="w-3 h-3" /> Deadline
                  </label>
                  <div className="font-semibold mt-1">
                    {deadline
                      ? deadline.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })
                      : '—'}
                  </div>
                </div>
                <div>
                  <label className="text-sm text-text-muted flex items-center gap-1">
                    <DollarSign className="w-3 h-3" /> Budget
                  </label>
                  <div className="font-semibold mt-1">{amount} {MTK_TOKEN.symbol}</div>
                </div>
              </div>

              {campaign?.categories && campaign.categories.length > 0 && (
                <>
                  <div className="h-px bg-border-color" />
                  <div className="flex flex-wrap gap-2">
                    {campaign.categories.map((cat) => (
                      <span key={cat} className="px-3 py-1 bg-elevated rounded-full text-xs text-text-secondary">
                        {cat}
                      </span>
                    ))}
                  </div>
                </>
              )}

              <div className="h-px bg-border-color" />

              <div className="grid grid-cols-2 gap-4">
                {advertiserAddress && (
                  <div>
                    <label className="text-sm text-text-muted flex items-center gap-2">
                      Advertiser
                      <button onClick={() => copyToClipboard(advertiserAddress)}>
                        <Copy className="w-3 h-3" />
                      </button>
                    </label>
                    <div className="flex items-center gap-2 mt-2">
                      <Jazzicon address={advertiserAddress} size={24} />
                      <span className="font-mono text-sm">{shortenAddress(advertiserAddress)}</span>
                    </div>
                  </div>
                )}
                {channel && (
                  <div>
                    <label className="text-sm text-text-muted">Influencer Channel</label>
                    <div className="flex items-center gap-2 mt-2">
                      {channel.profileImageUrl ? (
                        <img src={channel.profileImageUrl} alt={`@${channel.platformHandle}`} className="w-6 h-6 rounded-full" />
                      ) : (
                        <div className="w-6 h-6 rounded-full bg-sky-500/10 flex items-center justify-center">
                          <Twitter className="w-3 h-3 text-sky-500" />
                        </div>
                      )}
                      <span className="text-sm font-semibold text-accent-violet">@{channel.platformHandle}</span>
                      <span className="text-xs text-text-muted">({channel.followerCount.toLocaleString()} followers)</span>
                    </div>
                  </div>
                )}
              </div>
            </div>
          </div>

          {/* Action Panel — proposed: Accept/Decline. For channel-originated (isDirectDeal) only influencer can accept; for campaign-originated the non-proposer (advertiser or influencer) can accept. */}
          {currentStatus === 'proposed' && !isProposer && (campaign?.isDirectDeal === true ? isInfluencer : (isAdvertiser || isInfluencer)) && (
            <div className="bg-card border-l-4 border-l-success border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-success mb-4">New Proposal</h3>
              <p className="text-text-secondary mb-6">
                {isAdvertiser
                  ? 'An influencer wants to complete this campaign. Review the details and decide whether to accept or decline.'
                  : 'An advertiser sent you a deal proposal. Review the details and decide whether to accept or decline.'}
              </p>
              <div className="flex gap-3">
                <GradientButton onClick={handleAccept} className="flex-1">
                  Accept Proposal
                </GradientButton>
                <button className="flex-1 px-6 py-3 rounded-lg border-2 border-error text-error hover:bg-error/10 transition-all font-medium">
                  Decline
                </button>
              </div>
            </div>
          )}

          {currentStatus === 'proposed' && isProposer && (
            <div className="bg-card border-l-4 border-l-success border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-success mb-4">Proposal sent</h3>
              <p className="text-text-secondary">
                {campaign?.isDirectDeal === true
                  ? 'Waiting for the channel owner to accept or decline your proposal.'
                  : isAdvertiser
                    ? 'Waiting for the influencer to accept or decline your proposal.'
                    : 'Waiting for the advertiser to accept or decline your proposal.'}
              </p>
            </div>
          )}

          {/* Fund only for advertiser */}
          {currentStatus === 'accepted' && isAdvertiser && (
            <div className="bg-card border-l-4 border-l-accent-blue border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-accent-blue mb-4">Fund this Campaign</h3>
              <p className="text-text-secondary mb-4">
                The proposal has been accepted. Deposit {amount} {MTK_TOKEN.symbol} into escrow to activate the campaign.
              </p>
              <p className="text-xs text-text-muted mb-4">
                This will approve the escrow contract and deposit tokens. You'll confirm two transactions in your wallet.
              </p>
              <GradientButton onClick={handleFund} className="w-full" disabled={isFunding}>
                {isFunding ? (
                  <>
                    <Loader2 className="w-5 h-5 mr-2 animate-spin" />
                    Processing Deposit...
                  </>
                ) : (
                  `Deposit ${amount} ${MTK_TOKEN.symbol}`
                )}
              </GradientButton>
              {fundingError && (
                <div className="mt-3 p-3 bg-error/10 border border-error/20 rounded-lg text-sm text-error">
                  {fundingError}
                </div>
              )}
            </div>
          )}

          {currentStatus === 'accepted' && isInfluencer && (
            <div className="bg-card border-l-4 border-l-accent-blue border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-accent-blue mb-4">Waiting for funding</h3>
              <p className="text-text-secondary">
                The advertiser will deposit {amount} {MTK_TOKEN.symbol} into escrow to activate the campaign.
              </p>
            </div>
          )}

          {/* Submit tweet only for influencer */}
          {currentStatus === 'funded' && isInfluencer && (
            <div className="bg-card border-l-4 border-l-accent-violet border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-accent-violet mb-4">Submit Your Tweet</h3>
              <p className="text-text-secondary mb-4">
                Campaign is funded! Post the tweet with the required content and submit the URL here.
              </p>
              <div className="flex gap-3">
                <div className="flex-1 relative">
                  <LinkIcon className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-text-muted" />
                  <Input
                    value={tweetUrl}
                    onChange={(e) => setTweetUrl(e.target.value)}
                    placeholder="https://x.com/..."
                    className="pl-10 bg-elevated border-border-color"
                  />
                </div>
                <GradientButton onClick={handleSubmitTweet} disabled={!tweetUrl || isSubmitting}>
                  {isSubmitting ? (
                    <>
                      <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                      Submitting...
                    </>
                  ) : 'Submit'}
                </GradientButton>
              </div>
              {submitError && (
                <div className="mt-3 p-3 bg-error/10 border border-error/20 rounded-lg text-sm text-error">
                  {submitError}
                </div>
              )}
            </div>
          )}

          {currentStatus === 'funded' && isAdvertiser && (
            <div className="bg-card border-l-4 border-l-accent-violet border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-accent-violet mb-4">Waiting for post</h3>
              <p className="text-text-secondary">
                The influencer will post the tweet and submit the URL here.
              </p>
            </div>
          )}

          {currentStatus === 'posted' && (
            <div className="bg-card border-l-4 border-l-accent-violet border-y border-r border-border-color rounded-xl p-6">
              <h3 className="text-accent-violet mb-4">Ready for Verification</h3>
              <p className="text-text-secondary mb-4">
                Tweet has been submitted. Click verify to start Chainlink oracle verification.
              </p>
              {deal.postUrl && (
                <div className="mb-4 p-3 bg-elevated rounded-lg flex items-center gap-2">
                  <LinkIcon className="w-4 h-4 text-text-muted flex-shrink-0" />
                  <a
                    href={deal.postUrl}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-sm text-accent-violet hover:text-accent-blue truncate"
                  >
                    {deal.postUrl}
                  </a>
                  <ExternalLink className="w-3 h-3 text-text-muted flex-shrink-0" />
                </div>
              )}
              <div className="mb-4 p-3 bg-elevated rounded-lg">
                <p className="text-sm text-text-muted mb-2">Verification checks:</p>
                <ul className="text-sm space-y-1">
                  <li className="flex items-center gap-2">
                    <div className="w-1.5 h-1.5 rounded-full bg-accent-violet" />
                    View count meets minimum requirement
                  </li>
                  <li className="flex items-center gap-2">
                    <div className="w-1.5 h-1.5 rounded-full bg-accent-violet" />
                    Tweet content matches campaign
                  </li>
                  <li className="flex items-center gap-2">
                    <div className="w-1.5 h-1.5 rounded-full bg-accent-violet" />
                    Tweet has not been edited
                  </li>
                  <li className="flex items-center gap-2">
                    <div className="w-1.5 h-1.5 rounded-full bg-accent-violet" />
                    Tweet posted for required duration
                  </li>
                </ul>
              </div>
              <GradientButton onClick={handleVerify} className="w-full" disabled={isVerifying}>
                {isVerifying ? (
                  <>
                    <Loader2 className="w-5 h-5 mr-2 animate-spin" />
                    Verifying...
                  </>
                ) : (
                  <>
                    <svg className="w-5 h-5 mr-2" viewBox="0 0 20 20" fill="currentColor">
                      <path d="M10 2L4 6V10C4 13.866 6.686 17.166 10 18C13.314 17.166 16 13.866 16 10V6L10 2Z" />
                    </svg>
                    Verify Post
                  </>
                )}
              </GradientButton>
            </div>
          )}

          {currentStatus === 'verifying' && (
            <div className="bg-card border-l-4 border-l-accent-violet border-y border-r border-border-color rounded-xl p-6">
              <div className="flex items-center gap-3 mb-4">
                <Loader2 className="w-6 h-6 text-accent-violet animate-spin" />
                <h3 className="text-accent-violet">Verification in Progress</h3>
              </div>
              <p className="text-text-secondary mb-4">
                Chainlink oracle is verifying the tweet... This usually takes 1-2 minutes. The page will update automatically when the result is in.
              </p>
              <div className="flex flex-wrap gap-2">
                <button
                  type="button"
                  onClick={handleRefreshStatus}
                  className="px-4 py-2 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all text-sm font-medium"
                >
                  Refresh status
                </button>
                <button
                  type="button"
                  onClick={handleVerify}
                  disabled={isVerifying}
                  className="px-4 py-2 rounded-lg border border-accent-violet/50 bg-accent-violet/10 hover:bg-accent-violet/20 transition-all text-sm font-medium text-accent-violet disabled:opacity-50"
                >
                  {isVerifying ? (
                    <>
                      <Loader2 className="w-4 h-4 animate-spin inline mr-2" />
                      Triggering...
                    </>
                  ) : (
                    'Trigger verify'
                  )}
                </button>
              </div>
            </div>
          )}

          {currentStatus === 'completed' && (
            <div className="bg-card border-l-4 border-l-success border-y border-r border-border-color rounded-xl p-6">
              <div className="flex items-center gap-3 mb-4">
                <Check className="w-6 h-6 text-success" />
                <h3 className="text-success">Deal Completed!</h3>
              </div>
              <p className="text-text-secondary mb-4">
                All verifications passed. Funds have been released to the influencer.
              </p>
              {campaign?.txHash && (
                <a
                  href={`https://sepolia.etherscan.io/tx/${campaign.txHash}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="inline-flex items-center gap-2 text-accent-violet hover:text-accent-blue transition-colors"
                >
                  View transaction on Etherscan
                  <ExternalLink className="w-4 h-4" />
                </a>
              )}
            </div>
          )}

          {currentStatus === 'disputed' && (
            <div className="bg-card border-l-4 border-l-error border-y border-r border-border-color rounded-xl p-6">
              <div className="flex items-center gap-3 mb-4">
                <AlertTriangle className="w-6 h-6 text-error" />
                <h3 className="text-error">Deal Disputed</h3>
              </div>
              <p className="text-text-secondary mb-2">
                This deal has been disputed and is under review.
              </p>
              <p className="text-sm text-text-muted">
                Both parties will be contacted. The dispute will be resolved through the on-chain arbitration process.
              </p>
            </div>
          )}

          {/* Verification History */}
          {verifications.length > 0 && (
            <div className="bg-card border border-border-color rounded-xl p-6">
              <h2 className="mb-4">Verification History</h2>
              <div className="overflow-x-auto">
                <table className="w-full text-sm">
                  <thead>
                    <tr className="border-b border-border-color">
                      <th className="text-left py-3 px-2 text-text-muted font-medium">#</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Date</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Views</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Content</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Duration</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Edited</th>
                      <th className="text-left py-3 px-2 text-text-muted font-medium">Status</th>
                    </tr>
                  </thead>
                  <tbody>
                    {verifications.map((v, idx) => (
                      <tr key={v.id} className="border-b border-border-color last:border-0">
                        <td className="py-3 px-2">{idx + 1}</td>
                        <td className="py-3 px-2 font-mono text-xs">
                          {new Date(v.createdAt).toLocaleString('en-US', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })}
                        </td>
                        <td className="py-3 px-2">
                          {v.viewsChecked != null ? (
                            <>
                              <span className={v.viewsChecked >= (campaign?.minViews || 0) ? 'text-success' : 'text-error'}>
                                {v.viewsChecked.toLocaleString()}
                              </span>
                              <span className="text-text-muted"> / {(campaign?.minViews || 0).toLocaleString()}</span>
                            </>
                          ) : '—'}
                        </td>
                        <td className="py-3 px-2">
                          {v.contentMatched != null ? (
                            <span className={`px-2 py-0.5 rounded text-xs ${v.contentMatched ? 'bg-success/10 text-success' : 'bg-error/10 text-error'}`}>
                              {v.contentMatched ? 'Match' : 'Mismatch'}
                            </span>
                          ) : '—'}
                        </td>
                        <td className="py-3 px-2">
                          {v.durationChecked != null ? (
                            <span className={v.durationChecked ? 'text-success' : 'text-error'}>
                              {v.durationChecked ? 'Met' : 'Not met'}
                            </span>
                          ) : '—'}
                        </td>
                        <td className="py-3 px-2">
                          {v.isEdited != null ? (
                            v.isEdited ? <span className="text-error">Yes</span> : <span className="text-success">No</span>
                          ) : '—'}
                        </td>
                        <td className="py-3 px-2">
                          <StatusBadge status={v.status === 'passed' ? 'completed' : v.status === 'failed' ? 'disputed' : 'verifying'} />
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          )}
        </div>

        {/* Right Column - 1/3 */}
        <div className="lg:col-span-1 space-y-6">
          {/* Deal Info */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h3 className="mb-4">Deal Information</h3>

            <div className="flex justify-center mb-6">
              <StatusBadge status={currentStatus} className="text-lg px-4 py-2" />
            </div>

            <div className="text-center mb-6">
              <div className="text-3xl font-bold mb-1">{amount} {MTK_TOKEN.symbol}</div>
              <div className="text-sm text-text-muted">Campaign Budget</div>
            </div>

            <div className="space-y-3 text-sm">
              <div className="flex justify-between">
                <span className="text-text-muted">Created:</span>
                <span>{new Date(deal.createdAt).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}</span>
              </div>

              <div className="h-px bg-border-color" />

              <div>
                <div className="text-text-muted mb-2">Proposed by:</div>
                <div className="flex items-center gap-2">
                  <Jazzicon address={deal.proposedBy} size={20} />
                  <span className="font-mono text-xs">{shortenAddress(deal.proposedBy)}</span>
                  <button onClick={() => copyToClipboard(deal.proposedBy)} className="text-text-muted hover:text-text-primary">
                    <Copy className="w-3 h-3" />
                  </button>
                </div>
              </div>
              {advertiserAddress && (
                <div>
                  <div className="text-text-muted mb-2">Advertiser:</div>
                  <div className="flex items-center gap-2">
                    <Jazzicon address={advertiserAddress} size={20} />
                    <span className="font-mono text-xs">{shortenAddress(advertiserAddress)}</span>
                    <button onClick={() => copyToClipboard(advertiserAddress)} className="text-text-muted hover:text-text-primary">
                      <Copy className="w-3 h-3" />
                    </button>
                  </div>
                </div>
              )}

              {channel && (
                <div>
                  <div className="text-text-muted mb-2">Channel:</div>
                  <div className="flex items-center gap-2">
                    <Twitter className="w-4 h-4 text-sky-500" />
                    <span className="text-accent-violet font-semibold">@{channel.platformHandle}</span>
                  </div>
                </div>
              )}

              <div className="h-px bg-border-color" />

              {deal.onchainCampaignId && (
                <div className="flex justify-between">
                  <span className="text-text-muted">On-chain ID:</span>
                  <span className="font-mono text-xs">{deal.onchainCampaignId}</span>
                </div>
              )}

              {campaign?.txHash && deal.onchainCampaignId != null && (
                <div>
                  <div className="text-text-muted mb-1">Deposit TX:</div>
                  <a
                    href={`https://sepolia.etherscan.io/tx/${campaign.txHash}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="font-mono text-xs text-accent-violet hover:text-accent-blue flex items-center gap-1"
                  >
                    {shortenAddress(campaign.txHash)}
                    <ExternalLink className="w-3 h-3" />
                  </a>
                </div>
              )}
              {chainTxHashes.acceptDealTx && (
                <div>
                  <div className="text-text-muted mb-1">Accept TX:</div>
                  <a
                    href={`https://sepolia.etherscan.io/tx/${chainTxHashes.acceptDealTx}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="font-mono text-xs text-accent-violet hover:text-accent-blue flex items-center gap-1"
                  >
                    {shortenAddress(chainTxHashes.acceptDealTx)}
                    <ExternalLink className="w-3 h-3" />
                  </a>
                </div>
              )}
              {chainTxHashes.settlementTx && (
                <div>
                  <div className="text-text-muted mb-1">Settlement TX:</div>
                  <a
                    href={`https://sepolia.etherscan.io/tx/${chainTxHashes.settlementTx}`}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="font-mono text-xs text-accent-violet hover:text-accent-blue flex items-center gap-1"
                  >
                    {shortenAddress(chainTxHashes.settlementTx)}
                    <ExternalLink className="w-3 h-3" />
                  </a>
                </div>
              )}

              {deal.postUrl && (
                <div>
                  <div className="text-text-muted mb-1">Tweet:</div>
                  <a
                    href={deal.postUrl}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-xs text-accent-violet hover:text-accent-blue flex items-center gap-1"
                  >
                    View tweet
                    <ExternalLink className="w-3 h-3" />
                  </a>
                </div>
              )}
            </div>
          </div>

          {/* Activity Timeline */}
          <div className="bg-card border border-border-color rounded-xl p-6">
            <h3 className="mb-4">Activity</h3>
            <div className="space-y-4">
              {timeline.map((activity, index) => (
                <div key={index} className="flex gap-3">
                  <div className="relative flex flex-col items-center">
                    <div className={`w-2 h-2 rounded-full ${activity.color} flex-shrink-0 mt-2`} />
                    {index < timeline.length - 1 && (
                      <div className="w-px h-full bg-border-color mt-1" />
                    )}
                  </div>
                  <div className="flex-1 pb-2">
                    <p className="text-sm mb-1">{activity.text}</p>
                    <p className="text-xs text-text-muted">{activity.time}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Actions */}
          {(currentStatus === 'posted' || currentStatus === 'completed') && (
            <div className="bg-card border border-border-color rounded-xl p-6">
              <h3 className="mb-4">Actions</h3>
              <div className="space-y-3">
                {currentStatus === 'posted' && (
                  <button
                    onClick={handleDispute}
                    className="w-full px-4 py-3 rounded-lg border-2 border-error text-error hover:bg-error/10 transition-all font-medium text-sm"
                  >
                    Dispute Deal
                  </button>
                )}
                {!campaign?.isDirectDeal && (
                  <Link to={`/campaign/${deal.campaignId}`} className="block">
                    <button className="w-full px-4 py-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all font-medium text-sm">
                      View Campaign
                    </button>
                  </Link>
                )}
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
