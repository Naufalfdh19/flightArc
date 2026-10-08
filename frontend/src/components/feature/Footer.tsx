import { Globe, Headphones, Mail, Send } from "lucide-react";
import Logo from "../ui/Logo";

const FOOTER_LINKS = [
    { title: "Products", links: ["Flights", "Hotels", "Trains", "Car Rentals", "Packages"] },
    { title: "Company", links: ["About Us", "Careers", "Newsroom", "Investors"] },
    { title: "Support", links: ["Help Center", "Contact Us", "Privacy Policy", "Terms"] },
]

const SOCIALS = [
    { label: "Website", icon: <Globe size={16} /> },
    { label: "Email", icon: <Mail size={16} /> },
    { label: "Newsletter", icon: <Send size={16} /> },
    { label: "Support", icon: <Headphones size={16} /> },
]

export default function Footer() {
    return (
        <footer className="border-t border-white/5 bg-night">
            <div className="mx-auto max-w-7xl px-4 py-16 sm:px-6 lg:px-8">
                <div className="grid gap-12 lg:grid-cols-[1.2fr_2fr]">
                    <div className="max-w-sm">
                        <Logo />
                        <p className="mt-5 leading-relaxed text-mist">
                            Your all-in-one travel companion. Book smarter and journey beautifully.
                        </p>
                        <div className="mt-6 flex gap-3">
                            {SOCIALS.map((s) => (
                                <a
                                    key={s.label}
                                    href="#"
                                    aria-label={s.label}
                                    className="grid h-10 w-10 place-items-center rounded-full border border-white/10 text-mist transition-all duration-300 hover:-translate-y-0.5 hover:border-gold-400/50 hover:text-gold-300"
                                >
                                    {s.icon}
                                </a>
                            ))}
                        </div>
                    </div>

                    <div className="grid grid-cols-2 gap-8 sm:grid-cols-3">
                        {FOOTER_LINKS.map((col) => (
                            <div key={col.title}>
                                <h4 className="mb-4 text-xs font-semibold tracking-[0.2em] text-cream uppercase">{col.title}</h4>
                                <ul className="space-y-3">
                                    {col.links.map((link) => (
                                        <li key={link}>
                                            <a href="#" className="text-sm text-mist transition-colors hover:text-gold-300">
                                                {link}
                                            </a>
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        ))}
                    </div>
                </div>

                <div className="mt-14 flex flex-col items-center justify-between gap-4 border-t border-white/5 pt-8 text-xs text-mist/70 sm:flex-row">
                    <p>© {new Date().getFullYear()} FlightArc. All rights reserved.</p>
                    <p>Crafted for travellers, everywhere.</p>
                </div>
            </div>
        </footer>
    )
}
