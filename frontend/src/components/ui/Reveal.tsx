import { useEffect, useRef, useState, type ReactNode } from "react"
import { cn } from "../../utils/cn"

interface RevealProps {
    children: ReactNode
    /** Delay in ms, handy for staggering siblings */
    delay?: number
    className?: string
}

/** Fades and lifts its children in the first time they scroll into view. */
export default function Reveal({ children, delay = 0, className }: RevealProps) {
    const ref = useRef<HTMLDivElement>(null)
    const [visible, setVisible] = useState(false)

    useEffect(() => {
        const el = ref.current
        if (!el) return
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setVisible(true)
                    observer.disconnect()
                }
            },
            { threshold: 0.15, rootMargin: "0px 0px -40px 0px" },
        )
        observer.observe(el)
        return () => observer.disconnect()
    }, [])

    return (
        <div
            ref={ref}
            className={cn("reveal", visible && "is-visible", className)}
            style={{ transitionDelay: `${delay}ms` }}
        >
            {children}
        </div>
    )
}
