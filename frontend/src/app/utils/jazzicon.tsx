import { useEffect, useRef } from 'react';

// Simple Jazzicon-like avatar generator
export function Jazzicon({ address, size = 40 }: { address: string; size?: number }) {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    if (!canvasRef.current) return;
    
    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    // Generate deterministic colors from address
    const hash = address.slice(2, 10);
    const seed = parseInt(hash, 16);
    
    const colors = [
      `hsl(${(seed % 360)}, 70%, 60%)`,
      `hsl(${((seed * 7) % 360)}, 70%, 50%)`,
      `hsl(${((seed * 13) % 360)}, 70%, 55%)`,
    ];

    // Clear canvas
    ctx.clearRect(0, 0, size, size);

    // Draw background
    ctx.fillStyle = colors[0];
    ctx.fillRect(0, 0, size, size);

    // Draw circles
    for (let i = 0; i < 4; i++) {
      const x = ((seed * (i + 3)) % size);
      const y = ((seed * (i + 7)) % size);
      const radius = ((seed * (i + 11)) % (size / 3)) + 10;
      
      ctx.fillStyle = colors[(i % 2) + 1];
      ctx.beginPath();
      ctx.arc(x, y, radius, 0, Math.PI * 2);
      ctx.fill();
    }
  }, [address, size]);

  return (
    <canvas
      ref={canvasRef}
      width={size}
      height={size}
      className="rounded-full"
      style={{ width: size, height: size }}
    />
  );
}
