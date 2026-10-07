import type { ReactNode } from "react"
import Globe from "../three/LazyGlobe"
import Logo from "../ui/Logo"

interface AuthLayoutProps {
    eyebrow: string
    title: string
    subtitle: ReactNode
    children: ReactNode
}

/** Split-screen shell for login / register: animated globe on the left, form on the right. */
export default function AuthLayout({ eyebrow, title, subtitle, children }: AuthLayoutProps) {
    return (
        <div className="flex min-h-svh bg-ink">
            {/* Visual panel */}
            <aside className="relative isolate hidden w-[46%] flex-col justify-between overflow-hidden border-r border-white/5 bg-night p-10 lg:flex xl:p-14">
                <div className="absolute inset-0 -z-10 bg-[radial-gradient(ellipse_at_50%_55%,rgb(224_176_98/0.16),transparent_60%)]" />
                <div className="absolute inset-0 -z-10">
                    <Globe stars />
                </div>

                <Logo />

                <figure className="glass max-w-md animate-fade-up rounded-3xl p-6">
                    <blockquote className="font-display text-xl leading-snug font-light text-cream">
                        “Booked Jakarta → Tokyo in two minutes and saved enough for an extra night in Kyoto.”
                    </blockquote>
                    <figcaption className="mt-4 flex items-center gap-3 text-sm">
                        <span className="grid h-9 w-9 place-items-center rounded-full bg-gold-400 font-bold text-ink">R</span>
                        <span>
                            <span className="block font-semibold text-cream">Rina A.</span>
                            <span className="text-mist">Frequent flyer</span>
                        </span>
                    </figcaption>
                </figure>
            </aside>

            {/* Form panel */}
            <main className="relative isolate flex flex-1 items-center justify-center px-4 py-12 sm:px-8">
                <div className="absolute top-0 right-0 -z-10 h-96 w-96 rounded-full bg-gold-400/10 blur-3xl" />
                <div className="w-full max-w-md">
                    <Logo className="mb-10 lg:hidden" />
                    <p className="animate-fade-up text-xs font-semibold tracking-[0.25em] text-gold-400 uppercase">{eyebrow}</p>
                    <h1 className="mt-3 animate-fade-up font-display text-4xl leading-tight font-light text-cream [animation-delay:80ms] sm:text-5xl">
                        {title}
                    </h1>
                    <p className="mt-3 animate-fade-up text-mist [animation-delay:160ms]">{subtitle}</p>
                    <div className="mt-8 animate-fade-up [animation-delay:240ms]">{children}</div>
                </div>
            </main>
        </div>
    )
}
