import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { useAccount, useConnect, useDisconnect } from 'wagmi';
import { useConnectModal } from '@rainbow-me/rainbowkit';
import { apiClient, setAuthToken, clearAuthToken } from '../lib/api';

interface WalletContextType {
  address: string | null;
  isConnected: boolean;
  role: 'advertiser' | 'influencer' | 'both' | null;
  displayName: string;
  connect: () => void;
  disconnect: () => void;
  setRole: (role: 'advertiser' | 'influencer' | 'both') => void;
  setDisplayName: (name: string) => void;
}

const WalletContext = createContext<WalletContextType | undefined>(undefined);

export function WalletProvider({ children }: { children: ReactNode }) {
  const { address: wagmiAddress, isConnected: wagmiConnected } = useAccount();
  const { disconnect: wagmiDisconnect } = useDisconnect();
  const { openConnectModal } = useConnectModal();

  const [role, setRoleState] = useState<'advertiser' | 'influencer' | 'both' | null>(null);
  const [displayName, setDisplayNameState] = useState<string>('');

  // Load user profile when wallet connects
  useEffect(() => {
    if (wagmiConnected && wagmiAddress) {
      apiClient('/api/auth/me')
        .then((data) => {
          if (data.user) {
            setRoleState(data.user.role);
            setDisplayNameState(data.user.displayName || '');
          }
        })
        .catch(() => {
          // Not authenticated yet — user needs to verify via SIWE on connect-wallet page
        });
    }
  }, [wagmiConnected, wagmiAddress]);

  const connect = () => {
    openConnectModal?.();
  };

  const disconnect = () => {
    wagmiDisconnect();
    clearAuthToken();
    setRoleState(null);
    setDisplayNameState('');
  };

  const setRole = async (newRole: 'advertiser' | 'influencer' | 'both') => {
    setRoleState(newRole);
    try {
      await apiClient('/api/auth/me', {
        method: 'PATCH',
        body: JSON.stringify({ role: newRole }),
      });
    } catch {
      // Silently fail — role is set locally even if API fails
    }
  };

  const setDisplayName = async (name: string) => {
    setDisplayNameState(name);
    try {
      await apiClient('/api/auth/me', {
        method: 'PATCH',
        body: JSON.stringify({ displayName: name }),
      });
    } catch {
      // Silently fail
    }
  };

  return (
    <WalletContext.Provider
      value={{
        address: wagmiAddress ?? null,
        isConnected: wagmiConnected,
        role,
        displayName,
        connect,
        disconnect,
        setRole,
        setDisplayName,
      }}
    >
      {children}
    </WalletContext.Provider>
  );
}

export function useWallet() {
  const context = useContext(WalletContext);
  if (!context) {
    throw new Error('useWallet must be used within WalletProvider');
  }
  return context;
}
