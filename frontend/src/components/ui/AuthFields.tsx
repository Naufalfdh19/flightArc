import type { InputHTMLAttributes, ReactNode } from "react"
import Button from "./Button"

interface FieldProps extends InputHTMLAttributes<HTMLInputElement> {
    label: string
    icon: ReactNode
}

export function Field({ label, icon, id, ...props }: FieldProps) {
    const inputId = id ?? label.toLowerCase().replace(/\s+/g, "-")
    return (
        <div className="flex flex-col gap-2">
            <label htmlFor={inputId} className="text-sm font-medium text-cream/90">{label}</label>
            <div className="group relative">
                <span className="pointer-events-none absolute top-1/2 left-4 -translate-y-1/2 text-mist transition-colors group-focus-within:text-gold-300">
                    {icon}
                </span>
                <input id={inputId} className="field pl-11" {...props} />
            </div>
        </div>
    )
}

export function SocialButtons() {
    return (
        <>
            <div className="grid grid-cols-2 gap-3">
                <Button type="outline" height="md" square="md">
                    <svg viewBox="0 0 24 24" className="h-4 w-4" aria-hidden="true">
                        <path fill="#EA4335" d="M12 10.2v3.9h5.5c-.2 1.3-1.6 3.8-5.5 3.8-3.3 0-6-2.7-6-6.1s2.7-6.1 6-6.1c1.9 0 3.1.8 3.8 1.5l2.6-2.5C16.8 3.2 14.6 2.2 12 2.2 6.6 2.2 2.2 6.6 2.2 12s4.4 9.8 9.8 9.8c5.7 0 9.4-4 9.4-9.6 0-.6-.1-1.1-.2-1.6H12z" />
                    </svg>
                    Google
                </Button>
                <Button type="outline" height="md" square="md">
                    <svg viewBox="0 0 24 24" className="h-4 w-4 fill-current" aria-hidden="true">
                        <path d="M16.4 12.6c0-2.4 2-3.6 2.1-3.7-1.1-1.7-2.9-1.9-3.5-1.9-1.5-.2-2.9.9-3.7.9-.8 0-1.9-.9-3.2-.8-1.6 0-3.1 1-4 2.4-1.7 2.9-.4 7.3 1.2 9.7.8 1.2 1.8 2.5 3 2.4 1.2 0 1.7-.8 3.1-.8 1.5 0 1.9.8 3.2.8 1.3 0 2.2-1.2 3-2.4.9-1.4 1.3-2.7 1.3-2.8 0 0-2.5-1-2.5-3.8zM14 5.4c.7-.8 1.1-1.9 1-3-1 0-2.1.7-2.8 1.5-.6.7-1.2 1.8-1 2.9 1 .1 2.1-.6 2.8-1.4z" />
                    </svg>
                    Apple
                </Button>
            </div>
            <div className="my-7 flex items-center gap-4 text-xs tracking-[0.2em] text-mist/70 uppercase">
                <span className="h-px flex-1 bg-white/10" />
                or with email
                <span className="h-px flex-1 bg-white/10" />
            </div>
        </>
    )
}
