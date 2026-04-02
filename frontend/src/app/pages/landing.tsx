import { useState } from 'react';
import { useNavigate } from 'react-router';
import { Shield, ArrowRight, CheckCircle, Zap, Eye, UserCheck, Lock, BarChart3, Mail } from 'lucide-react';
import { GradientButton } from '../components/gradient-button';
import { apiClient } from '../lib/api';

const steps = [
  {
    number: '01',
    title: 'Advertiser Funds Escrow',
    description: 'Define your campaign requirements and lock funds in a smart contract. Your money is safe — not held by any platform.',
  },
  {
    number: '02',
    title: 'Influencer Posts Tweet',
    description: 'A verified influencer accepts the deal and posts the agreed content on X/Twitter.',
  },
  {
    number: '03',
    title: 'Chainlink Verifies',
    description: 'Decentralized oracle nodes independently verify content match, view count, duration, and edit status.',
  },
  {
    number: '04',
    title: 'Funds Release Automatically',
    description: 'Verification passes → influencer gets paid. Fails → advertiser gets refunded. No middleman decides.',
  },
];

const features = [
  {
    icon: Lock,
    title: 'Non-Custodial Escrow',
    description: 'Funds locked in auditable smart contracts — not held by us or any third party.',
  },
  {
    icon: Eye,
    title: 'Oracle-Verified Views',
    description: 'Chainlink CRE nodes independently fetch and verify tweet metrics via consensus.',
  },
  {
    icon: UserCheck,
    title: 'Verified Influencers',
    description: 'Every influencer profile is verified — no bots, no fake accounts.',
  },
  {
    icon: Zap,
    title: 'Automatic Settlement',
    description: 'No disputes, no support tickets. Verification is on-chain and deterministic.',
  },
  {
    icon: CheckCircle,
    title: 'Content Hash Matching',
    description: 'Tweet text is hashed and matched on-chain — edits and modifications are detected.',
  },
  {
    icon: BarChart3,
    title: 'Transparent & Auditable',
    description: 'Every deal, verification, and payment is recorded on-chain for full transparency.',
  },
];

export function Landing() {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [waitlistStatus, setWaitlistStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle');
  const [waitlistMsg, setWaitlistMsg] = useState('');

  async function handleWaitlist(e: React.FormEvent) {
    e.preventDefault();
    if (!email) return;
    setWaitlistStatus('loading');
    try {
      await apiClient('/api/waitlist', {
        method: 'POST',
        body: JSON.stringify({ email }),
      });
      setWaitlistStatus('success');
      setWaitlistMsg("You're on the list! We'll be in touch.");
      setEmail('');
    } catch {
      setWaitlistStatus('error');
      setWaitlistMsg('Something went wrong. Try again.');
    }
  }

  return (
    <div className="min-h-screen bg-background text-text-primary">
      {/* Nav */}
      <nav className="fixed top-0 w-full z-50 bg-background/80 backdrop-blur-lg border-b border-border-color">
        <div className="max-w-6xl mx-auto px-6 h-16 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="w-9 h-9 bg-gradient-to-br from-accent-violet to-accent-blue rounded-lg flex items-center justify-center">
              <Shield className="w-5 h-5 text-white" />
            </div>
            <span className="text-lg font-bold">Chainfluence</span>
          </div>
          <div className="flex items-center gap-4">
            <a href="#how-it-works" className="text-sm text-text-secondary hover:text-text-primary transition hidden sm:block">
              How It Works
            </a>
            <a href="#features" className="text-sm text-text-secondary hover:text-text-primary transition hidden sm:block">
              Features
            </a>
            <GradientButton onClick={() => navigate('/connect')} className="h-9 px-4 text-sm">
              Launch App
            </GradientButton>
          </div>
        </div>
      </nav>

      {/* Hero */}
      <section className="pt-32 pb-20 px-6 relative overflow-hidden">
        <div className="absolute inset-0 bg-gradient-to-br from-accent-violet/5 via-transparent to-accent-blue/5" />
        <div className="absolute top-1/3 left-1/4 w-[500px] h-[500px] bg-accent-violet/8 rounded-full blur-3xl" />
        <div className="absolute bottom-1/4 right-1/4 w-[400px] h-[400px] bg-accent-blue/8 rounded-full blur-3xl" />

        <div className="relative max-w-4xl mx-auto text-center">
          <div className="flex flex-wrap items-center justify-center gap-3 mb-8">
            <div className="inline-flex items-center gap-2 px-4 py-2 bg-accent-violet/10 border border-accent-violet/20 rounded-full">
              <Zap className="w-4 h-4 text-accent-violet" />
              <span className="text-sm text-accent-violet font-medium">Powered by Chainlink CRE</span>
            </div>
            <div className="inline-flex items-center gap-2 px-3 py-2 bg-yellow-500/10 border border-yellow-500/30 rounded-full">
              <span className="text-sm text-yellow-500 font-medium">Live on Sepolia Testnet</span>
            </div>
          </div>

          <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold leading-tight mb-6">
            Influencer Marketing,{' '}
            <span className="bg-gradient-to-r from-accent-violet to-accent-blue bg-clip-text text-transparent">
              Without the Trust
            </span>
          </h1>

          <p className="text-lg sm:text-xl text-text-secondary max-w-2xl mx-auto mb-10">
            Advertisers fund escrow. Influencers post. Chainlink oracles verify delivery.
            Funds release automatically. No middleman. No disputes.
          </p>

          <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mb-12">
            <GradientButton onClick={() => navigate('/connect')} className="h-13 px-8 text-base">
              Get Started <ArrowRight className="w-5 h-5 ml-1" />
            </GradientButton>
            <a
              href="#how-it-works"
              className="h-13 px-8 inline-flex items-center justify-center rounded-lg border border-border-color text-text-secondary hover:text-text-primary hover:border-accent-violet/50 transition"
            >
              See How It Works
            </a>
          </div>

          {/* Social proof pills */}
          <div className="flex flex-wrap justify-center gap-3">
            <div className="px-4 py-2 bg-card border border-border-color rounded-full text-sm text-text-secondary">
              On-Chain Escrow
            </div>
            <div className="px-4 py-2 bg-card border border-border-color rounded-full text-sm text-text-secondary">
              Oracle-Verified Views
            </div>
            <div className="px-4 py-2 bg-card border border-border-color rounded-full text-sm text-text-secondary">
              Verified Influencers
            </div>
            <div className="px-4 py-2 bg-card border border-border-color rounded-full text-sm text-text-secondary">
              Zero Fees for Early Users
            </div>
          </div>
        </div>
      </section>

      {/* Pain Points — two columns */}
      <section className="py-20 px-6 border-t border-border-color">
        <div className="max-w-5xl mx-auto grid md:grid-cols-2 gap-12">
          <div className="bg-card border border-border-color rounded-xl p-8">
            <h3 className="text-xl font-bold mb-4">For Influencers</h3>
            <p className="text-text-secondary mb-6">Tired of chasing payments after you deliver?</p>
            <ul className="space-y-3">
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Funds locked in escrow <strong>before</strong> you post</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Automatic payment when verification passes</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>No invoicing, no payment delays, no ghosting</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Verified profile = higher trust with advertisers</span>
              </li>
            </ul>
          </div>

          <div className="bg-card border border-border-color rounded-xl p-8">
            <h3 className="text-xl font-bold mb-4">For Advertisers</h3>
            <p className="text-text-secondary mb-6">Tired of paying for fake engagement?</p>
            <ul className="space-y-3">
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Only pay when verified views + content match</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Auto-refund if requirements aren't met</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Verified influencer profiles — no bots or fakes</span>
              </li>
              <li className="flex items-start gap-3">
                <CheckCircle className="w-5 h-5 text-success mt-0.5 flex-shrink-0" />
                <span>Edit detection — no bait-and-switch content</span>
              </li>
            </ul>
          </div>
        </div>
      </section>

      {/* How It Works */}
      <section id="how-it-works" className="py-20 px-6 border-t border-border-color">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl sm:text-4xl font-bold text-center mb-4">How It Works</h2>
          <p className="text-text-secondary text-center mb-16 max-w-2xl mx-auto">
            Four steps. Fully automated. No trust required.
          </p>

          <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {steps.map((step) => (
              <div key={step.number} className="bg-card border border-border-color rounded-xl p-6 relative">
                <div className="text-4xl font-bold text-accent-violet/20 mb-3">{step.number}</div>
                <h3 className="text-lg font-semibold mb-2">{step.title}</h3>
                <p className="text-sm text-text-secondary">{step.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Features Grid */}
      <section id="features" className="py-20 px-6 border-t border-border-color">
        <div className="max-w-5xl mx-auto">
          <h2 className="text-3xl sm:text-4xl font-bold text-center mb-4">Why Chainfluence</h2>
          <p className="text-text-secondary text-center mb-16 max-w-2xl mx-auto">
            Built on Chainlink's decentralized oracle network for trustless influencer marketing.
          </p>

          <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {features.map((feature) => (
              <div key={feature.title} className="bg-card border border-border-color rounded-xl p-6 hover:border-accent-violet/30 transition">
                <div className="w-10 h-10 bg-accent-violet/10 rounded-lg flex items-center justify-center mb-4">
                  <feature.icon className="w-5 h-5 text-accent-violet" />
                </div>
                <h3 className="font-semibold mb-2">{feature.title}</h3>
                <p className="text-sm text-text-secondary">{feature.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CTA + Waitlist */}
      <section className="py-20 px-6 border-t border-border-color">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-3xl sm:text-4xl font-bold mb-4">Get Early Access</h2>
          <p className="text-text-secondary mb-8 text-lg">
            We're live on Sepolia testnet and onboarding early users. Join the waitlist to be first in line for mainnet launch — zero platform fees for early adopters.
          </p>

          {waitlistStatus === 'success' ? (
            <div className="inline-flex items-center gap-2 px-6 py-3 bg-success/10 border border-success/30 rounded-lg">
              <CheckCircle className="w-5 h-5 text-success" />
              <span className="text-success font-medium">{waitlistMsg}</span>
            </div>
          ) : (
            <form onSubmit={handleWaitlist} className="flex flex-col sm:flex-row items-center justify-center gap-3 max-w-md mx-auto mb-6">
              <div className="relative flex-1 w-full">
                <Mail className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-text-secondary" />
                <input
                  type="email"
                  placeholder="your@email.com"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                  className="w-full h-12 pl-10 pr-4 bg-card border border-border-color rounded-lg text-text-primary placeholder:text-text-secondary/50 focus:outline-none focus:border-accent-violet/50 transition"
                />
              </div>
              <GradientButton type="submit" loading={waitlistStatus === 'loading'} className="h-12 px-6 whitespace-nowrap">
                Join Waitlist
              </GradientButton>
            </form>
          )}

          {waitlistStatus === 'error' && (
            <p className="text-red-400 text-sm mt-2">{waitlistMsg}</p>
          )}

          <div className="mt-6">
            <button
              onClick={() => navigate('/connect')}
              className="text-sm text-text-secondary hover:text-accent-violet transition underline underline-offset-4"
            >
              Or try the testnet app now →
            </button>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-8 px-6 border-t border-border-color">
        <div className="max-w-6xl mx-auto flex flex-col sm:flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-2">
            <div className="w-7 h-7 bg-gradient-to-br from-accent-violet to-accent-blue rounded-md flex items-center justify-center">
              <Shield className="w-4 h-4 text-white" />
            </div>
            <span className="text-sm font-semibold">Chainfluence</span>
          </div>
          <p className="text-sm text-text-secondary">
            Trustless influencer marketing, powered by Chainlink CRE
          </p>
        </div>
      </footer>
    </div>
  );
}
