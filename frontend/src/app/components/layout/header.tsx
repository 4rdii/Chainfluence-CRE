import { Bell, Menu } from 'lucide-react';
import { Jazzicon } from '../../utils/jazzicon';
import { useWallet } from '../../context/wallet-context';

interface HeaderProps {
  onMenuClick: () => void;
}

export function Header({ onMenuClick }: HeaderProps) {
  const { address } = useWallet();

  const truncateAddress = (addr: string) => {
    return `${addr.slice(0, 6)}...${addr.slice(-4)}`;
  };

  return (
    <header className="fixed top-0 left-0 right-0 h-16 bg-card border-b border-border-color z-50 flex items-center px-6">
      <div className="flex items-center gap-4 flex-1">
        {/* Mobile menu button */}
        <button
          onClick={onMenuClick}
          className="lg:hidden p-2 hover:bg-elevated rounded-lg transition-colors"
        >
          <Menu className="w-5 h-5" />
        </button>

        {/* Logo */}
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 bg-gradient-to-br from-accent-violet to-accent-blue rounded-lg flex items-center justify-center">
            <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M10 2L4 6V10C4 13.866 6.686 17.166 10 18C13.314 17.166 16 13.866 16 10V6L10 2Z" fill="white" fillOpacity="0.9"/>
              <path d="M7 10L9 12L13 8" stroke="#7C3AED" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
            </svg>
          </div>
          <span className="font-semibold text-lg hidden sm:block">Chainfluence</span>
        </div>
      </div>

      {/* Right side */}
      <div className="flex items-center gap-3">
        {/* Network pill */}
        <div className="hidden sm:flex items-center gap-2 px-3 py-1.5 bg-elevated border border-border-color rounded-full">
          <div className="w-2 h-2 bg-success rounded-full animate-pulse" />
          <span className="text-sm text-text-secondary">Sepolia</span>
        </div>

        {/* Wallet address */}
        {address && (
          <div className="flex items-center gap-2 px-3 py-1.5 bg-elevated border border-border-color rounded-full">
            <Jazzicon address={address} size={24} />
            <span className="text-sm font-mono hidden sm:block">{truncateAddress(address)}</span>
          </div>
        )}

        {/* Notifications */}
        <button className="relative p-2 hover:bg-elevated rounded-lg transition-colors">
          <Bell className="w-5 h-5" />
          <span className="absolute top-1 right-1 w-2 h-2 bg-error rounded-full" />
        </button>
      </div>
    </header>
  );
}
