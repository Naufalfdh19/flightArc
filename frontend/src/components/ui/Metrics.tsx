import { type ReactNode } from "react"
import { cn } from "../../utils/cn"

interface MetricsProps {
    icon: ReactNode
    title: string
    subtitle: string
    className?: string
}

export default function Metrics({ icon, title, subtitle, className }: MetricsProps) {
    return (
        <div className={cn("group flex items-center gap-4", className)}>
            <div className="grid h-12 w-12 shrink-0 place-items-center rounded-2xl border border-gold-400/20 bg-gold-400/10 text-gold-300 transition-all duration-300 group-hover:scale-110 group-hover:bg-gold-400/20">
                {icon}
            </div>
            <div>
                <p className="font-display text-lg font-medium text-cream sm:text-xl">{title}</p>
                <p className="text-sm text-mist">{subtitle}</p>
            </div>
        </div>
    )
}
