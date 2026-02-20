import { ButtonHTMLAttributes, forwardRef } from 'react';
import { cn } from './ui/utils';
import { Loader2 } from 'lucide-react';

interface GradientButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  loading?: boolean;
  icon?: React.ReactNode;
}

export const GradientButton = forwardRef<HTMLButtonElement, GradientButtonProps>(
  ({ children, loading, icon, className, disabled, ...props }, ref) => {
    return (
      <button
        ref={ref}
        disabled={disabled || loading}
        className={cn(
          "relative inline-flex items-center justify-center gap-2 px-6 h-11 rounded-lg font-medium text-white",
          "bg-gradient-to-br from-accent-violet to-accent-blue",
          "hover:shadow-lg hover:shadow-accent-violet/25 transition-all duration-200",
          "disabled:opacity-50 disabled:cursor-not-allowed",
          className
        )}
        {...props}
      >
        {loading ? (
          <Loader2 className="w-4 h-4 animate-spin" />
        ) : icon ? (
          icon
        ) : null}
        {children}
      </button>
    );
  }
);

GradientButton.displayName = 'GradientButton';
