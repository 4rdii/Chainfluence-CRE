import { useState, useEffect, type ReactNode } from 'react';
import { useDisconnect, useAccount, useBalance, useReadContract } from 'wagmi';
import { formatUnits } from 'viem';
import {
  User, Twitter, Shield, Coins, Copy, Check, ExternalLink,
  Loader2, BadgeCheck, Trash2, Plus, LogOut, Save,
} from 'lucide-react';
import { Switch } from '../components/ui/switch';
import { Input } from '../components/ui/input';
import { GradientButton } from '../components/gradient-button';
import { apiClient, clearAuthToken } from '../lib/api';
import { erc20Abi } from '../lib/escrow-abi';
import { ESCROW_ADDRESS, MTK_TOKEN } from '../lib/wagmi';

interface User {
  id: number;
  walletAddress: string;
  displayName: string | null;
  role: string;
  createdAt: string;
}

interface Channel {
  id: number;
  platformHandle: string;
  profileImageUrl?: string;
  followerCount: number;
  pricePerPost: number;
  categories: string[];
  isActive: boolean;
  verifiedAt?: string;
}

function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);
  const copy = () => {
    navigator.clipboard.writeText(text);
    setCopied(true);
    setTimeout(() => setCopied(false), 1500);
  };
  return (
    <button onClick={copy} className="text-text-muted hover:text-text-primary transition-colors p-1">
      {copied ? <Check className="w-3.5 h-3.5 text-success" /> : <Copy className="w-3.5 h-3.5" />}
    </button>
  );
}

function SectionHeader({ icon, title, subtitle }: { icon: ReactNode; title: string; subtitle: string }) {
  return (
    <div className="flex items-center gap-3 mb-6">
      <div className="w-10 h-10 rounded-lg bg-accent-violet/10 flex items-center justify-center flex-shrink-0">
        {icon}
      </div>
      <div>
        <h2 className="text-base font-semibold">{title}</h2>
        <p className="text-sm text-text-muted">{subtitle}</p>
      </div>
    </div>
  );
}

const ROLES = [
  { value: 'both', label: 'Both', description: 'Create campaigns and apply as influencer' },
  { value: 'advertiser', label: 'Advertiser', description: 'Run campaigns and hire influencers' },
  { value: 'influencer', label: 'Influencer', description: 'Apply to campaigns and earn rewards' },
];

export function Settings() {
  const { address } = useAccount();
  const { disconnect } = useDisconnect();

  // User profile
  const [user, setUser] = useState<User | null>(null);
  const [displayName, setDisplayName] = useState('');
  const [role, setRole] = useState('both');
  const [savingProfile, setSavingProfile] = useState(false);
  const [profileSaved, setProfileSaved] = useState(false);
  const [profileError, setProfileError] = useState('');

  // Channels
  const [channels, setChannels] = useState<Channel[]>([]);
  const [connectingX, setConnectingX] = useState(false);

  // ETH balance
  const { data: ethBalance } = useBalance({ address });

  // MTK balance
  const { data: mtkBalance } = useReadContract({
    address: MTK_TOKEN.address,
    abi: erc20Abi,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
  });

  useEffect(() => {
    apiClient('/api/auth/me')
      .then((data) => {
        setUser(data.user);
        setDisplayName(data.user.displayName || '');
        setRole(data.user.role || 'both');
      })
      .catch(() => {});
    apiClient('/api/channels')
      .then((data) => setChannels(data.channels || []))
      .catch(() => {});
  }, []);

  const saveProfile = async () => {
    setSavingProfile(true);
    setProfileError('');
    try {
      const { user: updated } = await apiClient('/api/auth/me', {
        method: 'PATCH',
        body: JSON.stringify({ displayName: displayName.trim() || null, role }),
      });
      setUser(updated);
      setProfileSaved(true);
      setTimeout(() => setProfileSaved(false), 2000);
    } catch {
      setProfileError('Failed to save changes');
    } finally {
      setSavingProfile(false);
    }
  };

  const connectX = async () => {
    setConnectingX(true);
    try {
      const { url } = await apiClient('/api/auth/x/connect');
      if (url) window.location.href = url;
    } catch {
      setConnectingX(false);
    }
  };

  const toggleChannelActive = (id: number, current: boolean) => {
    setChannels((prev) => prev.map((ch) => ch.id === id ? { ...ch, isActive: !current } : ch));
    apiClient(`/api/channels/${id}`, {
      method: 'PATCH',
      body: JSON.stringify({ isActive: !current }),
    }).catch(() => {
      // revert on failure
      setChannels((prev) => prev.map((ch) => ch.id === id ? { ...ch, isActive: current } : ch));
    });
  };

  const deleteChannel = async (id: number) => {
    if (!confirm('Remove this channel from your account?')) return;
    await apiClient(`/api/channels/${id}`, { method: 'DELETE' }).catch(() => {});
    setChannels((prev) => prev.filter((ch) => ch.id !== id));
  };

  const handleDisconnect = () => {
    clearAuthToken();
    disconnect();
  };

  const formattedMtk = mtkBalance
    ? parseFloat(formatUnits(mtkBalance as bigint, MTK_TOKEN.decimals)).toFixed(4)
    : '—';
  const formattedEth = ethBalance
    ? parseFloat(ethBalance.formatted).toFixed(6)
    : '—';
  const memberSince = user?.createdAt
    ? new Date(user.createdAt).toLocaleDateString('en-US', { month: 'long', year: 'numeric' })
    : '—';

  return (
    <div className="max-w-3xl mx-auto space-y-6">
      <div>
        <h1>Settings</h1>
        <p className="text-text-secondary mt-1">Manage your profile, channels, and account</p>
      </div>

      {/* ── Profile ─────────────────────────────────── */}
      <div className="bg-card border border-border-color rounded-xl p-6">
        <SectionHeader
          icon={<User className="w-5 h-5 text-accent-violet" />}
          title="Profile"
          subtitle="Your display name and account role"
        />

        <div className="space-y-5">
          {/* Wallet address — read only */}
          <div>
            <label className="block text-sm text-text-muted mb-1.5">Wallet Address</label>
            <div className="flex items-center gap-2 px-3 py-2.5 bg-elevated border border-border-color rounded-lg">
              <span className="font-mono text-sm flex-1 truncate">{address || '—'}</span>
              {address && <CopyButton text={address} />}
              {address && (
                <a
                  href={`https://sepolia.etherscan.io/address/${address}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-text-muted hover:text-accent-violet transition-colors p-1"
                >
                  <ExternalLink className="w-3.5 h-3.5" />
                </a>
              )}
            </div>
          </div>

          {/* Display name */}
          <div>
            <label className="block text-sm text-text-muted mb-1.5">Display Name</label>
            <Input
              value={displayName}
              onChange={(e) => setDisplayName(e.target.value)}
              placeholder="e.g. CryptoWhale, DeFi.eth"
              className="bg-elevated border-border-color"
              maxLength={100}
            />
          </div>

          {/* Role */}
          <div>
            <label className="block text-sm text-text-muted mb-2">Account Role</label>
            <div className="grid grid-cols-3 gap-3">
              {ROLES.map((r) => (
                <button
                  key={r.value}
                  onClick={() => setRole(r.value)}
                  className={`p-3 rounded-lg border-2 text-left transition-all ${
                    role === r.value
                      ? 'border-accent-violet bg-accent-violet/10'
                      : 'border-border-color bg-elevated hover:border-accent-violet/40'
                  }`}
                >
                  <div className="font-semibold text-sm mb-0.5">{r.label}</div>
                  <div className="text-xs text-text-muted leading-tight">{r.description}</div>
                </button>
              ))}
            </div>
          </div>

          {/* Member since */}
          <div className="flex justify-between text-sm">
            <span className="text-text-muted">Member since</span>
            <span className="text-text-secondary">{memberSince}</span>
          </div>

          {/* Save button */}
          <div className="flex items-center gap-3 pt-1">
            <GradientButton onClick={saveProfile} disabled={savingProfile} className="gap-2">
              {savingProfile ? (
                <Loader2 className="w-4 h-4 animate-spin" />
              ) : profileSaved ? (
                <Check className="w-4 h-4" />
              ) : (
                <Save className="w-4 h-4" />
              )}
              {savingProfile ? 'Saving...' : profileSaved ? 'Saved!' : 'Save Changes'}
            </GradientButton>
            {profileError && <p className="text-sm text-error">{profileError}</p>}
          </div>
        </div>
      </div>

      {/* ── Connected Channels ───────────────────────── */}
      <div className="bg-card border border-border-color rounded-xl p-6">
        <div className="flex items-start justify-between mb-6">
          <SectionHeader
            icon={<Twitter className="w-5 h-5 text-sky-500" />}
            title="Connected Channels"
            subtitle="Your X (Twitter) accounts for influencer campaigns"
          />
          <button
            onClick={connectX}
            disabled={connectingX}
            className="flex items-center gap-2 px-3 py-2 rounded-lg border border-border-color hover:border-accent-violet/50 text-sm font-medium transition-all flex-shrink-0 mt-0.5"
          >
            {connectingX ? (
              <Loader2 className="w-4 h-4 animate-spin" />
            ) : (
              <Plus className="w-4 h-4" />
            )}
            Add Account
          </button>
        </div>

        {channels.length === 0 ? (
          <div className="text-center py-8 border border-dashed border-border-color rounded-lg">
            <Twitter className="w-8 h-8 text-text-muted mx-auto mb-3" />
            <p className="text-sm text-text-secondary mb-4">No X accounts connected yet</p>
            <button
              onClick={connectX}
              disabled={connectingX}
              className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-sky-500/10 border border-sky-500/20 text-sky-500 hover:bg-sky-500/20 transition-all text-sm font-medium"
            >
              {connectingX ? <Loader2 className="w-4 h-4 animate-spin" /> : <Twitter className="w-4 h-4" />}
              Connect X Account
            </button>
          </div>
        ) : (
          <div className="space-y-3">
            {channels.map((ch) => (
              <div
                key={ch.id}
                className="flex items-center gap-4 p-4 bg-elevated border border-border-color rounded-lg"
              >
                {/* Avatar */}
                {ch.profileImageUrl ? (
                  <img src={ch.profileImageUrl} alt={`@${ch.platformHandle}`} className="w-10 h-10 rounded-full flex-shrink-0" />
                ) : (
                  <div className="w-10 h-10 rounded-full bg-sky-500/10 flex items-center justify-center flex-shrink-0">
                    <Twitter className="w-5 h-5 text-sky-500" />
                  </div>
                )}

                {/* Info */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-1.5 mb-0.5">
                    <span className="font-semibold text-accent-violet truncate">@{ch.platformHandle}</span>
                    {ch.verifiedAt && <BadgeCheck className="w-4 h-4 text-sky-500 flex-shrink-0" />}
                  </div>
                  <div className="text-xs text-text-muted">
                    {ch.followerCount.toLocaleString()} followers
                    {ch.pricePerPost > 0 && <span> · ${ch.pricePerPost}/post</span>}
                  </div>
                  {ch.categories.length > 0 && (
                    <div className="flex gap-1 mt-1.5 flex-wrap">
                      {ch.categories.map((cat) => (
                        <span key={cat} className="px-2 py-0.5 bg-background rounded text-xs text-text-secondary">{cat}</span>
                      ))}
                    </div>
                  )}
                </div>

                {/* Controls */}
                <div className="flex items-center gap-3 flex-shrink-0">
                  <div className="flex items-center gap-1.5">
                    <span className="text-xs text-text-muted">{ch.isActive ? 'Active' : 'Paused'}</span>
                    <Switch
                      checked={ch.isActive}
                      onCheckedChange={() => toggleChannelActive(ch.id, ch.isActive)}
                    />
                  </div>
                  <button
                    onClick={() => deleteChannel(ch.id)}
                    className="p-1.5 hover:bg-error/10 rounded-lg transition-colors text-text-muted hover:text-error"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* ── Wallet & Balances ────────────────────────── */}
      <div className="bg-card border border-border-color rounded-xl p-6">
        <SectionHeader
          icon={<Shield className="w-5 h-5 text-success" />}
          title="Wallet & Security"
          subtitle="Your connected wallet and on-chain balances"
        />

        <div className="space-y-4">
          {/* Network */}
          <div className="flex items-center justify-between py-3 border-b border-border-color">
            <span className="text-sm text-text-muted">Network</span>
            <div className="flex items-center gap-2">
              <div className="w-2 h-2 rounded-full bg-success" />
              <span className="text-sm font-medium">Sepolia Testnet</span>
            </div>
          </div>

          {/* ETH Balance */}
          <div className="flex items-center justify-between py-3 border-b border-border-color">
            <span className="text-sm text-text-muted">ETH Balance</span>
            <span className="text-sm font-semibold font-mono">{formattedEth} ETH</span>
          </div>

          {/* MTK Balance */}
          <div className="flex items-center justify-between py-3 border-b border-border-color">
            <span className="text-sm text-text-muted">{MTK_TOKEN.symbol} Balance</span>
            <span className="text-sm font-semibold font-mono">{formattedMtk} {MTK_TOKEN.symbol}</span>
          </div>

          {/* Disconnect */}
          <div className="pt-2">
            <button
              onClick={handleDisconnect}
              className="flex items-center gap-2 px-4 py-2.5 rounded-lg border-2 border-error/30 text-error hover:bg-error/10 hover:border-error/60 transition-all text-sm font-medium"
            >
              <LogOut className="w-4 h-4" />
              Disconnect Wallet
            </button>
          </div>
        </div>
      </div>

      {/* ── Token & Contract Info ────────────────────── */}
      <div className="bg-card border border-border-color rounded-xl p-6">
        <SectionHeader
          icon={<Coins className="w-5 h-5 text-warning" />}
          title="Token & Contract Info"
          subtitle="Addresses for the escrow and reward token on Sepolia"
        />

        <div className="space-y-3">
          {[
            { label: 'MTK Token', address: MTK_TOKEN.address },
            { label: 'Escrow Contract', address: ESCROW_ADDRESS },
          ].map(({ label, address: addr }) => (
            <div key={label} className="flex items-center justify-between gap-3 p-3 bg-elevated border border-border-color rounded-lg">
              <div className="min-w-0">
                <div className="text-xs text-text-muted mb-1">{label}</div>
                <span className="font-mono text-xs truncate block">{addr}</span>
              </div>
              <div className="flex items-center gap-1 flex-shrink-0">
                <CopyButton text={addr} />
                <a
                  href={`https://sepolia.etherscan.io/address/${addr}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-text-muted hover:text-accent-violet transition-colors p-1"
                >
                  <ExternalLink className="w-3.5 h-3.5" />
                </a>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
