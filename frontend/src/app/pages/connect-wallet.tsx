import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router';
import { Shield, Check, Megaphone, Radio as RadioIcon, Handshake, Loader2 } from 'lucide-react';
import { useAccount, useSignMessage } from 'wagmi';
import { useWallet } from '../context/wallet-context';
import { GradientButton } from '../components/gradient-button';
import { cn } from '../components/ui/utils';
import { Input } from '../components/ui/input';
import { apiClient, setAuthToken } from '../lib/api';

export function ConnectWallet() {
  const navigate = useNavigate();
  const { connect, address, isConnected, setRole, setDisplayName } = useWallet();
  const { signMessageAsync } = useSignMessage();
  const { chain } = useAccount();
  const [step, setStep] = useState<'connect' | 'sign' | 'role'>('connect');
  const [selectedRole, setSelectedRole] = useState<'advertiser' | 'influencer' | 'both' | null>(null);
  const [name, setName] = useState('');
  const [signing, setSigning] = useState(false);

  // Auto-advance to sign step when wallet connects
  useEffect(() => {
    if (isConnected && address && step === 'connect') {
      setStep('sign');
    }
  }, [isConnected, address, step]);

  const handleConnect = () => {
    connect();
  };

  const handleSign = async () => {
    if (!address) return;
    setSigning(true);
    try {
      // Get nonce from backend
      const { nonce } = await apiClient('/api/auth/nonce');

      // Build SIWE message string (EIP-4361 format)
      const domain = window.location.host;
      const uri = window.location.origin;
      const chainId = chain?.id || 11155111;
      const issuedAt = new Date().toISOString();
      const messageStr = `${domain} wants you to sign in with your Ethereum account:\n${address}\n\nSign in to Chainfluence\n\nURI: ${uri}\nVersion: 1\nChain ID: ${chainId}\nNonce: ${nonce}\nIssued At: ${issuedAt}`;
      const signature = await signMessageAsync({ message: messageStr });

      // Verify with backend
      const { token } = await apiClient('/api/auth/verify', {
        method: 'POST',
        body: JSON.stringify({ message: messageStr, signature }),
      });

      setAuthToken(token);
      setStep('role');
    } catch (err) {
      console.error('SIWE failed:', err);
    } finally {
      setSigning(false);
    }
  };

  const handleGetStarted = () => {
    if (selectedRole) {
      setRole(selectedRole);
      setDisplayName(name);
      navigate('/dashboard');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-6 relative overflow-hidden">
      {/* Background gradient mesh */}
      <div className="absolute inset-0 bg-gradient-to-br from-accent-violet/5 via-background to-accent-blue/5" />
      <div className="absolute top-1/4 left-1/4 w-96 h-96 bg-accent-violet/10 rounded-full blur-3xl" />
      <div className="absolute bottom-1/4 right-1/4 w-96 h-96 bg-accent-blue/10 rounded-full blur-3xl" />

      {/* Content card */}
      <div className="relative w-full max-w-md">
        <div className="bg-card border border-border-color rounded-xl p-8 shadow-2xl">
          {/* Logo */}
          <div className="flex justify-center mb-6">
            <div className="w-12 h-12 bg-gradient-to-br from-accent-violet to-accent-blue rounded-xl flex items-center justify-center">
              <Shield className="w-7 h-7 text-white" />
            </div>
          </div>

          {/* Title */}
          <h1 className="text-center mb-2">Welcome to Chainfluence</h1>
          <p className="text-center text-text-secondary mb-8">
            Connect your wallet to start creating or completing ad campaigns with trustless escrow.
          </p>

          {/* Steps */}
          {step === 'connect' && (
            <>
              <GradientButton
                onClick={handleConnect}
                className="w-full h-13 mb-6"
              >
                Connect Wallet
              </GradientButton>

              {/* Feature pills */}
              <div className="flex flex-wrap justify-center gap-2">
                <div className="px-3 py-1.5 bg-elevated border border-border-color rounded-full text-sm text-text-secondary">
                  Trustless Escrow
                </div>
                <div className="px-3 py-1.5 bg-elevated border border-border-color rounded-full text-sm text-text-secondary">
                  Twitter Verified
                </div>
                <div className="px-3 py-1.5 bg-elevated border border-border-color rounded-full text-sm text-text-secondary">
                  Chainlink Powered
                </div>
              </div>
            </>
          )}

          {step === 'sign' && address && (
            <>
              <div className="flex items-center justify-center gap-3 mb-6 p-4 bg-success/10 border border-success/20 rounded-lg">
                <Check className="w-5 h-5 text-success" />
                <span className="font-mono text-sm">{address.slice(0, 6)}...{address.slice(-4)}</span>
              </div>

              <p className="text-center text-text-secondary mb-6">
                Sign a message to verify ownership of this wallet
              </p>

              <GradientButton
                onClick={handleSign}
                disabled={signing}
                className="w-full h-13"
              >
                {signing ? (
                  <span className="flex items-center gap-2">
                    <Loader2 className="w-4 h-4 animate-spin" />
                    Signing...
                  </span>
                ) : (
                  'Sign Message'
                )}
              </GradientButton>
            </>
          )}

          {step === 'role' && (
            <>
              <h2 className="text-center mb-6">Choose your role</h2>
              
              <div className="space-y-3 mb-6">
                {/* Advertiser */}
                <button
                  onClick={() => setSelectedRole('advertiser')}
                  className={cn(
                    "w-full p-4 rounded-lg border-2 transition-all hover:border-accent-violet/50",
                    selectedRole === 'advertiser'
                      ? "border-accent-violet bg-accent-violet/10"
                      : "border-border-color bg-elevated"
                  )}
                >
                  <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-accent-violet to-accent-blue flex items-center justify-center flex-shrink-0">
                      <Megaphone className="w-5 h-5 text-white" />
                    </div>
                    <div className="text-left">
                      <div className="font-semibold">Advertiser</div>
                      <div className="text-sm text-text-secondary">Create campaigns and hire influencers</div>
                    </div>
                  </div>
                </button>

                {/* Influencer */}
                <button
                  onClick={() => setSelectedRole('influencer')}
                  className={cn(
                    "w-full p-4 rounded-lg border-2 transition-all hover:border-accent-violet/50",
                    selectedRole === 'influencer'
                      ? "border-accent-violet bg-accent-violet/10"
                      : "border-border-color bg-elevated"
                  )}
                >
                  <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-accent-violet to-accent-blue flex items-center justify-center flex-shrink-0">
                      <RadioIcon className="w-5 h-5 text-white" />
                    </div>
                    <div className="text-left">
                      <div className="font-semibold">Influencer</div>
                      <div className="text-sm text-text-secondary">Apply to campaigns and earn rewards</div>
                    </div>
                  </div>
                </button>

                {/* Both */}
                <button
                  onClick={() => setSelectedRole('both')}
                  className={cn(
                    "w-full p-4 rounded-lg border-2 transition-all hover:border-accent-violet/50",
                    selectedRole === 'both'
                      ? "border-accent-violet bg-accent-violet/10"
                      : "border-border-color bg-elevated"
                  )}
                >
                  <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-accent-violet to-accent-blue flex items-center justify-center flex-shrink-0">
                      <Handshake className="w-5 h-5 text-white" />
                    </div>
                    <div className="text-left">
                      <div className="font-semibold">Both</div>
                      <div className="text-sm text-text-secondary">Create campaigns and apply to others</div>
                    </div>
                  </div>
                </button>
              </div>

              {/* Optional display name */}
              <div className="mb-6">
                <label className="block text-sm text-text-secondary mb-2">
                  Display Name (optional)
                </label>
                <Input
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  placeholder="Enter your name"
                  className="bg-elevated border-border-color"
                />
              </div>

              <GradientButton
                onClick={handleGetStarted}
                disabled={!selectedRole}
                className="w-full h-13"
              >
                Get Started
              </GradientButton>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
