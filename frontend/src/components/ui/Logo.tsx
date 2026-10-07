import { Link } from "react-router-dom"
import { cn } from "../../utils/cn"

export default function Logo({ className }: { className?: string }) {
    return (
        <Link to="/" className={cn("group inline-flex items-center gap-2.5", className)} aria-label="FlightArc home">
            <svg viewBox="0 0 32 32" className="h-8 w-8 transition-transform duration-500 group-hover:rotate-12" aria-hidden="true">
                <defs>
                    <linearGradient id="logo-arc" x1="0" y1="1" x2="1" y2="0">
                        <stop offset="0" stopColor="var(--color-gold-500)" />
                        <stop offset="1" stopColor="var(--color-gold-200)" />
                    </linearGradient>
                </defs>
                <circle cx="16" cy="16" r="15" fill="none" stroke="currentColor" strokeOpacity=".15" />
                <path d="M5 23 C 10 8, 22 6, 27 11" fill="none" stroke="url(#logo-arc)" strokeWidth="2.5" strokeLinecap="round" />
                <circle cx="5" cy="23" r="2.2" fill="var(--color-gold-400)" />
                <circle cx="27" cy="11" r="2.2" fill="var(--color-gold-200)" />
            </svg>
            <span className="font-display text-xl font-medium tracking-tight text-cream">
                Flight<span className="text-gold-400">Arc</span>
            </span>
        </Link>
    )
}
