import { cn } from "../components/ui/utils";

type StatusType = 
  | 'completed' 
  | 'verifying' 
  | 'funded' 
  | 'posted' 
  | 'proposed' 
  | 'disputed'
  | 'accepted'
  | 'active'
  | 'draft'
  | 'inactive';

const statusStyles: Record<StatusType, string> = {
  completed: 'bg-success/10 text-success border-success/20',
  verifying: 'bg-accent-violet/10 text-accent-violet border-accent-violet/20',
  funded: 'bg-accent-blue/10 text-accent-blue border-accent-blue/20',
  posted: 'bg-warning/10 text-warning border-warning/20',
  proposed: 'bg-text-muted/10 text-text-muted border-text-muted/20',
  disputed: 'bg-error/10 text-error border-error/20',
  accepted: 'bg-accent-blue/10 text-accent-blue border-accent-blue/20',
  active: 'bg-success/10 text-success border-success/20',
  draft: 'bg-text-muted/10 text-text-muted border-text-muted/20',
  inactive: 'bg-text-muted/10 text-text-muted border-text-muted/20',
};

export function StatusBadge({ status, className }: { status: StatusType; className?: string }) {
  return (
    <span
      className={cn(
        "inline-flex items-center px-3 py-1 rounded-full text-xs font-medium border capitalize",
        statusStyles[status],
        className
      )}
    >
      {status}
    </span>
  );
}
