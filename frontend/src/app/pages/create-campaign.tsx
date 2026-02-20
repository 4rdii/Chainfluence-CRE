import { useState } from 'react';
import { useNavigate } from 'react-router';
import { Twitter, Check, Info, ChevronLeft } from 'lucide-react';
import { useAccount, useReadContract } from 'wagmi';
import { parseUnits, formatUnits } from 'viem';
import { Input } from '../components/ui/input';
import { Textarea } from '../components/ui/textarea';
import { GradientButton } from '../components/gradient-button';
import { cn } from '../components/ui/utils';
import { erc20Abi } from '../lib/escrow-abi';
import { MTK_TOKEN } from '../lib/wagmi';
import { apiClient } from '../lib/api';

const steps = [
  { id: 1, label: 'Platform' },
  { id: 2, label: 'Content' },
  { id: 3, label: 'Requirements' },
  { id: 4, label: 'Payment' },
  { id: 5, label: 'Confirm' },
];

export function CreateCampaign() {
  const navigate = useNavigate();
  const { address } = useAccount();
  const [currentStep, setCurrentStep] = useState(1);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState('');
  const [txHash, setTxHash] = useState('');
  const [campaignId, setCampaignId] = useState('');

  // Form state
  const [formData, setFormData] = useState({
    platform: 'twitter',
    title: '',
    description: '',
    contentText: '',
    categories: [] as string[],
    minViews: '',
    duration: '',
    deadline: '',
    amount: '',
  });

  const [newCategory, setNewCategory] = useState('');

  // Read MTK balance (for display in payment step)
  const { data: mtkBalance } = useReadContract({
    address: MTK_TOKEN.address,
    abi: erc20Abi,
    functionName: 'balanceOf',
    args: address ? [address] : undefined,
  });

  const handleNext = () => {
    if (currentStep < 5) {
      setCurrentStep(currentStep + 1);
    }
  };

  const handleBack = () => {
    if (currentStep > 1) {
      setCurrentStep(currentStep - 1);
    }
  };

  const addCategory = () => {
    if (newCategory && !formData.categories.includes(newCategory)) {
      setFormData({ ...formData, categories: [...formData.categories, newCategory] });
      setNewCategory('');
    }
  };

  const removeCategory = (cat: string) => {
    setFormData({ ...formData, categories: formData.categories.filter(c => c !== cat) });
  };

  const handleSubmit = async () => {
    setIsSubmitting(true);
    setSubmitError('');
    try {
      const amount = parseUnits(formData.amount || '0', MTK_TOKEN.decimals);

      // Create draft campaign in backend (no on-chain payment yet)
      const { campaign } = await apiClient('/api/campaigns', {
        method: 'POST',
        body: JSON.stringify({
          title: formData.title,
          contentText: formData.contentText,
          tokenAddress: MTK_TOKEN.address,
          amount: amount.toString(),
          minViews: parseInt(formData.minViews) || 0,
          campaignDuration: parseDuration(formData.duration),
          expiryDeadline: formData.deadline || new Date(Date.now() + 30 * 86400000).toISOString(),
          categories: formData.categories,
        }),
      });

      setCampaignId(campaign.id.toString());
      setTxHash('draft'); // Signal success without actual tx
    } catch (err: any) {
      console.error('Create campaign failed:', err);
      const msg = err?.shortMessage || err?.message || 'Failed to create campaign';
      setSubmitError(msg);
    } finally {
      setIsSubmitting(false);
    }
  };

  function parseDuration(val: string): number {
    const map: Record<string, number> = {
      '1d': 86400, '3d': 259200, '7d': 604800, '14d': 1209600, '30d': 2592000,
    };
    return map[val] || 0;
  }

  const platformFee = formData.amount ? (parseFloat(formData.amount) * 0.025).toFixed(2) : '0.00';
  const totalDeposit = formData.amount ? (parseFloat(formData.amount) + parseFloat(platformFee)).toFixed(2) : '0.00';
  const influencerReceives = formData.amount || '0.00';
  const formattedBalance = mtkBalance ? parseFloat(formatUnits(mtkBalance as bigint, MTK_TOKEN.decimals)).toFixed(2) : '—';

  return (
    <div className="max-w-2xl mx-auto">
      <div className="mb-8">
        <h1 className="mb-2">Create Campaign</h1>
        <p className="text-text-secondary">Set up a new advertising campaign with trustless escrow</p>
      </div>

      {/* Stepper */}
      <div className="mb-8">
        <div className="flex items-center justify-between">
          {steps.map((step, index) => (
            <div key={step.id} className="flex items-center flex-1">
              <div className="flex flex-col items-center flex-1">
                <div
                  className={cn(
                    'w-10 h-10 rounded-full flex items-center justify-center border-2 transition-all',
                    step.id < currentStep
                      ? 'bg-success border-success'
                      : step.id === currentStep
                      ? 'bg-accent-violet border-accent-violet animate-pulse'
                      : 'bg-elevated border-border-color'
                  )}
                >
                  {step.id < currentStep ? (
                    <Check className="w-5 h-5 text-white" />
                  ) : (
                    <span className={cn(
                      'text-sm font-semibold',
                      step.id === currentStep ? 'text-white' : 'text-text-muted'
                    )}>
                      {step.id}
                    </span>
                  )}
                </div>
                <span
                  className={cn(
                    'text-xs mt-2 text-center hidden sm:block',
                    step.id === currentStep ? 'text-text-primary font-semibold' : 'text-text-muted'
                  )}
                >
                  {step.label}
                </span>
              </div>
              {index < steps.length - 1 && (
                <div
                  className={cn(
                    'h-0.5 flex-1 -mt-5',
                    step.id < currentStep ? 'bg-success' : 'bg-border-color'
                  )}
                />
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Form Content */}
      <div className="bg-card border border-border-color rounded-xl p-6 sm:p-8 min-h-[400px]">
        {/* Step 1: Platform */}
        {currentStep === 1 && (
          <div className="space-y-6">
            <h2>Choose Platform</h2>
            <button
              className="w-full p-6 rounded-xl border-2 border-accent-violet bg-accent-violet/10 transition-all"
            >
              <div className="flex items-center gap-4">
                <div className="w-12 h-12 rounded-xl bg-sky-500/20 flex items-center justify-center">
                  <Twitter className="w-6 h-6 text-sky-500" />
                </div>
                <div className="text-left flex-1">
                  <div className="font-semibold text-lg mb-1">Twitter / X</div>
                  <div className="text-sm text-text-secondary">
                    Verify tweets, view counts, posting duration, and edit detection
                  </div>
                </div>
                <Check className="w-6 h-6 text-accent-violet" />
              </div>
            </button>
          </div>
        )}

        {/* Step 2: Content */}
        {currentStep === 2 && (
          <div className="space-y-6">
            <h2>Campaign Content</h2>

            <div>
              <label className="block mb-2">Campaign Title *</label>
              <Input
                value={formData.title}
                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                placeholder="Enter campaign title"
                className="bg-elevated border-border-color"
              />
            </div>

            <div>
              <label className="block mb-2">Description (optional)</label>
              <Textarea
                value={formData.description}
                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                placeholder="Describe your campaign..."
                className="bg-elevated border-border-color"
                rows={3}
              />
            </div>

            <div>
              <label className="block mb-2">Tweet Content *</label>
              <Textarea
                value={formData.contentText}
                onChange={(e) => setFormData({ ...formData, contentText: e.target.value })}
                placeholder="Enter the tweet text you want influencers to post..."
                className="bg-background border-border-color font-mono text-sm"
                rows={4}
              />
              <div className="text-xs text-text-muted mt-1 text-right">
                {formData.contentText.length}/280
              </div>
            </div>

            <div>
              <label className="block mb-2">Categories</label>
              <div className="flex gap-2 mb-3">
                <Input
                  value={newCategory}
                  onChange={(e) => setNewCategory(e.target.value)}
                  onKeyPress={(e) => e.key === 'Enter' && addCategory()}
                  placeholder="Add category (e.g., DeFi, NFT)"
                  className="bg-elevated border-border-color"
                />
                <button
                  onClick={addCategory}
                  className="px-4 rounded-lg bg-elevated border border-border-color hover:border-accent-violet/50 transition-all"
                >
                  Add
                </button>
              </div>
              <div className="flex flex-wrap gap-2">
                {formData.categories.map((cat) => (
                  <span
                    key={cat}
                    className="px-3 py-1.5 bg-accent-violet/10 border border-accent-violet/20 rounded-full text-sm flex items-center gap-2"
                  >
                    {cat}
                    <button onClick={() => removeCategory(cat)} className="hover:text-error">
                      ×
                    </button>
                  </span>
                ))}
              </div>
            </div>
          </div>
        )}

        {/* Step 3: Requirements */}
        {currentStep === 3 && (
          <div className="space-y-6">
            <h2>Campaign Requirements</h2>

            <div>
              <label className="block mb-2">Minimum Views</label>
              <Input
                type="number"
                value={formData.minViews}
                onChange={(e) => setFormData({ ...formData, minViews: e.target.value })}
                placeholder="10000"
                className="bg-elevated border-border-color"
              />
              <p className="text-xs text-text-muted mt-1">Enter 0 for no requirement</p>
            </div>

            <div>
              <label className="block mb-2">Campaign Duration</label>
              <select
                value={formData.duration}
                onChange={(e) => setFormData({ ...formData, duration: e.target.value })}
                className="w-full px-4 h-11 rounded-lg bg-elevated border border-border-color text-text-primary"
              >
                <option value="">Select duration...</option>
                <option value="no-requirement">No requirement</option>
                <option value="1d">1 day</option>
                <option value="3d">3 days</option>
                <option value="7d">7 days</option>
                <option value="14d">14 days</option>
                <option value="30d">30 days</option>
              </select>
            </div>

            <div>
              <label className="block mb-2">Application Deadline</label>
              <Input
                type="datetime-local"
                value={formData.deadline}
                onChange={(e) => setFormData({ ...formData, deadline: e.target.value })}
                className="bg-elevated border-border-color"
              />
            </div>

            <div className="p-4 bg-accent-blue/10 border border-accent-blue/20 rounded-lg flex gap-3">
              <Info className="w-5 h-5 text-accent-blue flex-shrink-0 mt-0.5" />
              <div className="text-sm text-text-secondary">
                The influencer must keep the tweet posted for the full campaign duration. Chainlink verification checks it hasn't been deleted or edited.
              </div>
            </div>
          </div>
        )}

        {/* Step 4: Payment */}
        {currentStep === 4 && (
          <div className="space-y-6">
            <h2>Payment Details</h2>

            <div>
              <label className="block mb-2">Token</label>
              <div className="w-full px-4 h-11 rounded-lg bg-elevated border border-border-color text-text-primary flex items-center gap-2">
                <span className="font-semibold">{MTK_TOKEN.symbol}</span>
                <span className="text-text-muted text-sm">({MTK_TOKEN.name})</span>
              </div>
            </div>

            <div>
              <label className="block mb-2">Campaign Budget</label>
              <div className="relative">
                <Input
                  type="number"
                  value={formData.amount}
                  onChange={(e) => setFormData({ ...formData, amount: e.target.value })}
                  placeholder="500"
                  className="bg-elevated border-border-color pr-16 text-xl h-14"
                />
                <span className="absolute right-4 top-1/2 -translate-y-1/2 text-text-muted font-medium">
                  {MTK_TOKEN.symbol}
                </span>
              </div>
            </div>

            <div className="bg-elevated border border-border-color rounded-lg p-4 space-y-2">
              <div className="flex justify-between text-sm">
                <span className="text-text-secondary">Influencer receives:</span>
                <span className="font-semibold">{influencerReceives} {MTK_TOKEN.symbol}</span>
              </div>
              <div className="flex justify-between text-sm">
                <span className="text-text-secondary">Platform fee (2.5%):</span>
                <span className="font-semibold">{platformFee} {MTK_TOKEN.symbol}</span>
              </div>
              <div className="h-px bg-border-color my-2" />
              <div className="flex justify-between">
                <span className="font-semibold">Total deposit:</span>
                <span className="font-semibold text-lg">{totalDeposit} {MTK_TOKEN.symbol}</span>
              </div>
            </div>

            <div className="text-sm text-text-secondary">
              Wallet balance: <span className="font-semibold text-text-primary">{formattedBalance} {MTK_TOKEN.symbol}</span>
            </div>
          </div>
        )}

        {/* Step 5: Confirm */}
        {currentStep === 5 && !txHash && (
          <div className="space-y-6">
            <h2>Confirm Campaign</h2>

            <div className="grid grid-cols-2 gap-x-6 gap-y-4 text-sm">
              <div className="text-text-secondary">Platform:</div>
              <div className="font-semibold">Twitter / X</div>

              <div className="text-text-secondary">Title:</div>
              <div className="font-semibold">{formData.title || '—'}</div>

              <div className="text-text-secondary">Content:</div>
              <div className="font-mono text-xs bg-background px-3 py-2 rounded border border-border-color break-words">
                {formData.contentText || '—'}
              </div>

              <div className="text-text-secondary">Min Views:</div>
              <div className="font-semibold">{formData.minViews || 'No requirement'}</div>

              <div className="text-text-secondary">Duration:</div>
              <div className="font-semibold">{formData.duration || 'No requirement'}</div>

              <div className="text-text-secondary">Deadline:</div>
              <div className="font-semibold">{formData.deadline || '—'}</div>

              <div className="text-text-secondary">Token:</div>
              <div className="font-semibold">{MTK_TOKEN.symbol}</div>

              <div className="text-text-secondary">Total Amount:</div>
              <div className="font-semibold text-lg">{totalDeposit} {MTK_TOKEN.symbol}</div>
            </div>

            <div className="p-4 bg-accent-blue/10 border border-accent-blue/20 rounded-lg flex gap-3">
              <Info className="w-5 h-5 text-accent-blue flex-shrink-0 mt-0.5" />
              <div className="text-sm text-text-secondary">
                Your campaign will be listed on the marketplace. Once an influencer applies and you accept the deal, you'll deposit {MTK_TOKEN.symbol} into escrow.
              </div>
            </div>

            <GradientButton
              onClick={handleSubmit}
              disabled={isSubmitting}
              loading={isSubmitting}
              className="w-full h-13"
            >
              {isSubmitting ? 'Creating Campaign...' : 'Create Campaign'}
            </GradientButton>

            {submitError && (
              <div className="p-4 bg-error/10 border border-error/20 rounded-lg flex gap-3">
                <Info className="w-5 h-5 text-error flex-shrink-0 mt-0.5" />
                <div className="text-sm text-error">{submitError}</div>
              </div>
            )}
          </div>
        )}

        {/* Success State */}
        {currentStep === 5 && txHash && (
          <div className="text-center py-8 space-y-6">
            <div className="w-16 h-16 rounded-full bg-success/10 flex items-center justify-center mx-auto">
              <Check className="w-8 h-8 text-success" />
            </div>

            <div>
              <h2 className="text-success mb-2">Campaign Created!</h2>
              <p className="text-text-secondary">Your campaign is now listed on the marketplace. Influencers can apply, and you'll fund the escrow after accepting a deal.</p>
            </div>

            <div className="bg-elevated border border-border-color rounded-lg p-4 space-y-2 text-left">
              <div className="flex justify-between text-sm">
                <span className="text-text-secondary">Campaign ID:</span>
                <span className="font-mono font-semibold">{campaignId}</span>
              </div>
              <div className="flex justify-between text-sm">
                <span className="text-text-secondary">Status:</span>
                <span className="text-accent-violet font-semibold">Listed on Marketplace</span>
              </div>
            </div>

            <div className="flex gap-3">
              <button
                onClick={() => navigate('/marketplace')}
                className="flex-1 px-6 py-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all font-medium"
              >
                Back to Marketplace
              </button>
              <GradientButton
                onClick={() => navigate(`/campaign/${campaignId}`)}
                className="flex-1"
              >
                View Campaign
              </GradientButton>
            </div>
          </div>
        )}

        {/* Navigation Buttons */}
        {currentStep < 5 && (
          <div className="flex gap-3 mt-8">
            {currentStep > 1 && (
              <button
                onClick={handleBack}
                className="flex items-center gap-2 px-6 py-3 rounded-lg border border-border-color hover:border-accent-violet/50 transition-all font-medium"
              >
                <ChevronLeft className="w-4 h-4" />
                Back
              </button>
            )}
            <GradientButton
              onClick={handleNext}
              className="flex-1"
            >
              {currentStep === 4 ? 'Review & Confirm' : 'Next'}
            </GradientButton>
          </div>
        )}
      </div>
    </div>
  );
}
