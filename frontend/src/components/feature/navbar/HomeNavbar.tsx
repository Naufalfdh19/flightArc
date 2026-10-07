import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { Menu, X } from "lucide-react"
import Button from "../../ui/Button"
import Logo from "../../ui/Logo"
import { cn } from "../../../utils/cn"

const NAV_LINKS = [
    { label: "Flights", href: "#search" },
    { label: "Hotels", href: "#services" },
    { label: "Destinations", href: "#destinations" },
    { label: "Deals", href: "#deals" },
]

function HomeNavbar() {
    const navigate = useNavigate();
    const [scrolled, setScrolled] = useState(false)
    const [open, setOpen] = useState(false)

    useEffect(() => {
        const onScroll = () => setScrolled(window.scrollY > 24)
        onScroll()
        window.addEventListener("scroll", onScroll, { passive: true })
        return () => window.removeEventListener("scroll", onScroll)
    }, [])

    return (
        <header
            className={cn(
                "fixed inset-x-0 top-0 z-50 transition-all duration-500",
                scrolled || open ? "border-b border-white/5 bg-ink/75 backdrop-blur-xl" : "bg-transparent",
            )}
        >
            <nav className="mx-auto flex h-18 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
                <Logo />

                <ul className="hidden items-center gap-1 md:flex">
                    {NAV_LINKS.map((link) => (
                        <li key={link.label}>
                            <a
                                href={link.href}
                                className="relative rounded-full px-4 py-2 text-sm font-medium text-mist transition-colors hover:text-cream after:absolute after:inset-x-4 after:bottom-1 after:h-px after:origin-left after:scale-x-0 after:bg-gold-400 after:transition-transform after:duration-300 hover:after:scale-x-100"
                            >
                                {link.label}
                            </a>
                        </li>
                    ))}
                </ul>

                <div className="hidden items-center gap-2 md:flex">
                    <Button type="ghost" onClick={() => navigate("/login")}>Log In</Button>
                    <Button onClick={() => navigate("/register")}>Register</Button>
                </div>

                <button
                    className="grid h-10 w-10 cursor-pointer place-items-center rounded-full text-cream transition-colors hover:bg-white/10 md:hidden"
                    onClick={() => setOpen((o) => !o)}
                    aria-label={open ? "Close menu" : "Open menu"}
                    aria-expanded={open}
                >
                    {open ? <X size={22} /> : <Menu size={22} />}
                </button>
            </nav>

            {/* Mobile menu */}
            <div
                className={cn(
                    "grid transition-[grid-template-rows,opacity] duration-400 ease-out md:hidden",
                    open ? "grid-rows-[1fr] opacity-100" : "grid-rows-[0fr] opacity-0",
                )}
            >
                <div className="overflow-hidden">
                    <div className="flex flex-col gap-1 px-4 pt-2 pb-6">
                        {NAV_LINKS.map((link) => (
                            <a
                                key={link.label}
                                href={link.href}
                                onClick={() => setOpen(false)}
                                className="rounded-xl px-4 py-3 text-base font-medium text-cream/90 hover:bg-white/5"
                            >
                                {link.label}
                            </a>
                        ))}
                        <div className="mt-3 grid grid-cols-2 gap-3">
                            <Button type="outline" height="md" onClick={() => navigate("/login")}>Log In</Button>
                            <Button height="md" onClick={() => navigate("/register")}>Register</Button>
                        </div>
                    </div>
                </div>
            </div>
        </header>
    )
}


export default HomeNavbar
