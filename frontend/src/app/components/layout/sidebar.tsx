import { Link, useLocation } from 'react-router';
import { LayoutDashboard, Store, Megaphone, Radio, Handshake, Settings, FileText, X } from 'lucide-react';
import { cn } from '../ui/utils';

interface SidebarProps {
  isOpen: boolean;
  onClose: () => void;
}

const navItems = [
  { path: '/dashboard', label: 'Dashboard', icon: LayoutDashboard },
  { path: '/marketplace', label: 'Marketplace', icon: Store },
  { path: '/my-campaigns', label: 'My Campaigns', icon: Megaphone },
  { path: '/my-channels', label: 'My Channels', icon: Radio },
  { path: '/my-deals', label: 'My Deals', icon: Handshake },
];

const bottomNavItems = [
  { path: '/settings', label: 'Settings', icon: Settings },
  { path: '/docs', label: 'Docs', icon: FileText },
];

export function Sidebar({ isOpen, onClose }: SidebarProps) {
  const location = useLocation();

  return (
    <>
      {/* Overlay for mobile */}
      {isOpen && (
        <div
          className="fixed inset-0 bg-background/80 backdrop-blur-sm z-40 lg:hidden"
          onClick={onClose}
        />
      )}

      {/* Sidebar */}
      <aside
        className={cn(
          "fixed top-16 left-0 bottom-0 w-60 bg-card border-r border-border-color z-40 transition-transform duration-300 flex flex-col",
          isOpen ? "translate-x-0" : "-translate-x-full lg:translate-x-0"
        )}
      >
        {/* Close button for mobile */}
        <button
          onClick={onClose}
          className="lg:hidden absolute top-4 right-4 p-2 hover:bg-elevated rounded-lg transition-colors"
        >
          <X className="w-5 h-5" />
        </button>

        {/* Navigation */}
        <nav className="flex-1 px-3 py-6 space-y-1 overflow-y-auto">
          {navItems.map((item) => {
            const Icon = item.icon;
            const isActive = location.pathname === item.path;
            
            return (
              <Link
                key={item.path}
                to={item.path}
                onClick={onClose}
                className={cn(
                  "flex items-center gap-3 px-3 py-2.5 rounded-lg transition-all duration-200 relative group",
                  isActive
                    ? "bg-elevated text-text-primary"
                    : "text-text-secondary hover:text-text-primary hover:bg-elevated/50"
                )}
              >
                {isActive && (
                  <div className="absolute left-0 top-0 bottom-0 w-1 bg-accent-violet rounded-r-full" />
                )}
                <Icon className="w-5 h-5 flex-shrink-0" />
                <span className="font-medium">{item.label}</span>
              </Link>
            );
          })}

          {/* Separator */}
          <div className="my-4 h-px bg-border-color" />

          {bottomNavItems.map((item) => {
            const Icon = item.icon;
            const isActive = location.pathname === item.path;
            
            return (
              <Link
                key={item.path}
                to={item.path}
                onClick={onClose}
                className={cn(
                  "flex items-center gap-3 px-3 py-2.5 rounded-lg transition-all duration-200 relative group",
                  isActive
                    ? "bg-elevated text-text-primary"
                    : "text-text-secondary hover:text-text-primary hover:bg-elevated/50"
                )}
              >
                {isActive && (
                  <div className="absolute left-0 top-0 bottom-0 w-1 bg-accent-violet rounded-r-full" />
                )}
                <Icon className="w-5 h-5 flex-shrink-0" />
                <span className="font-medium">{item.label}</span>
              </Link>
            );
          })}
        </nav>

        {/* Tagline at bottom */}
        <div className="px-6 py-4 border-t border-border-color">
          <p className="text-xs text-text-muted text-center leading-relaxed">
            Trustless influencer marketing, powered by Chainlink
          </p>
        </div>
      </aside>
    </>
  );
}
