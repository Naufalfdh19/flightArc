import { useState } from "react"
import { useNavigate } from "react-router-dom"
import {
    ArrowLeftRight, ArrowRight, BadgeCheck, BedDouble, CalendarDays, Car, Hotel, MapPin,
    Plane, PlaneLanding, PlaneTakeoff, Search, Star, TrainFront, TreePalm, Users, Zap,
} from "lucide-react"
import Navbar from "../../components/feature/navbar/HomeNavbar"
import Button from "../../components/ui/Button"
import Metrics from "../../components/ui/Metrics"
import Reveal from "../../components/ui/Reveal"
import Globe from "../../components/three/LazyGlobe"
import { SectionWrapper } from "../../components/wrapper/SectionWrapper"
import DestinationCard from "../../components/ui/card/DestinationCard"
import OfferCard from "../../components/ui/card/OffersCard"
import Footer from "../../components/feature/Footer"
import { LIMITED_OFFERS } from "../../const/data/offers"
import { TRENDING_DESTINATIONS } from "../../const/data/destinations"
import { cn } from "../../utils/cn"

const SEARCH_TABS = [
    { label: "Flights", icon: Plane },
    { label: "Hotels", icon: BedDouble },
    { label: "Trains", icon: TrainFront },
]

const SERVICES = [
    { label: "Flights", desc: "500+ airlines", icon: Plane },
    { label: "Hotels", desc: "300K+ stays", icon: Hotel },
    { label: "Trains", desc: "Scenic routes", icon: TrainFront },
    { label: "Car Rentals", desc: "Pick up anywhere", icon: Car },
    { label: "Packages", desc: "All-in-one trips", icon: TreePalm },
]

const METRICS = [
    { icon: <Plane size={22} />, title: "500+ Airlines", subtitle: "Worldwide coverage" },
    { icon: <Hotel size={22} />, title: "300K+ Hotels", subtitle: "Every budget & style" },
    { icon: <BadgeCheck size={22} />, title: "Best Price", subtitle: "Guaranteed" },
    { icon: <Zap size={22} />, title: "Instant", subtitle: "E-ticket delivery" },
]

function SearchField({ label, icon, children }: { label: string; icon: React.ReactNode; children: React.ReactNode }) {
    return (
        <label className="group flex flex-col gap-1.5 rounded-2xl border border-white/10 bg-white/[0.03] px-4 py-3 transition-colors focus-within:border-gold-400/60 focus-within:bg-white/[0.06] hover:border-white/20">
            <span className="flex items-center gap-1.5 text-[11px] font-semibold tracking-[0.18em] text-mist uppercase group-focus-within:text-gold-300">
                {icon}
                {label}
            </span>
            {children}
        </label>
    )
}

const scrollToSection = (id: string) => document.getElementById(id)?.scrollIntoView({ behavior: "smooth" })

const inputClass ="w-full bg-transparent text-base font-medium text-cream outline-none placeholder:text-mist/50"

function HomePage() {
    const navigate = useNavigate()
    const [activeTab, setActiveTab] = useState("Flights")
    const [from, setFrom] = useState("Jakarta (CGK)")
    const [to, setTo] = useState("")

    return (
        <div className="overflow-x-clip bg-ink">
            <Navbar />

            {/* ── Hero ── */}
            <section className="relative isolate flex min-h-svh items-center pt-24 pb-36 lg:pt-18 lg:pb-32">
                {/* Ambient background */}
                <div className="absolute inset-0 -z-20 bg-[radial-gradient(ellipse_at_75%_40%,rgb(224_176_98/0.14),transparent_55%),radial-gradient(ellipse_at_10%_90%,rgb(47_163_145/0.10),transparent_50%)]" />
                <div className="absolute inset-0 -z-20 bg-[linear-gradient(rgb(255_255_255/0.025)_1px,transparent_1px),linear-gradient(90deg,rgb(255_255_255/0.025)_1px,transparent_1px)] bg-size-[64px_64px] mask-[radial-gradient(ellipse_at_center,black,transparent_75%)]" />

                <div className="mx-auto grid w-full max-w-7xl items-center gap-6 px-4 sm:px-6 lg:grid-cols-[1.05fr_1fr] lg:px-8">
                    <div className="relative z-10 min-w-0 text-center lg:text-left">
                        <div className="glass mb-6 inline-flex animate-fade-up items-center gap-2.5 rounded-full py-1.5 pr-4 pl-2 text-sm text-cream/90">
                            <span className="relative flex h-2.5 w-2.5">
                                <span className="absolute inline-flex h-full w-full animate-ping-slow rounded-full bg-teal-300" />
                                <span className="relative inline-flex h-2.5 w-2.5 rounded-full bg-teal-300" />
                            </span>
                            Trusted by <strong className="font-semibold text-gold-200">2M+</strong> travellers
                        </div>

                        <h1 className="animate-fade-up font-display text-5xl leading-[1.02] font-light tracking-tight text-cream [animation-delay:120ms] sm:text-6xl lg:text-7xl xl:text-[5.25rem]">
                            Your world, <br className="hidden sm:block" />
                            <em className="text-gradient-gold font-normal">beautifully</em> within reach
                        </h1>

                        <p className="mx-auto mt-6 max-w-xl animate-fade-up text-base leading-relaxed text-mist [animation-delay:240ms] sm:text-lg lg:mx-0">
                            Compare fares from 500+ airlines, find stays you'll love, and watch your next journey arc across the globe — all in one place.
                        </p>

                        <div className="mt-8 flex animate-fade-up flex-col items-center justify-center gap-3 [animation-delay:360ms] sm:flex-row lg:justify-start">
                            <Button height="md" className="w-full px-7 sm:w-auto" onClick={() => scrollToSection("search")}>
                                Start exploring <ArrowRight size={18} />
                            </Button>
                            <Button type="outline" height="md" className="w-full px-7 sm:w-auto" onClick={() => scrollToSection("deals")}>
                                View deals
                            </Button>
                        </div>

                        <div className="mt-10 flex animate-fade-up items-center justify-center gap-4 [animation-delay:480ms] lg:justify-start">
                            <div className="flex -space-x-2.5">
                                {["#e0b062", "#6fd3c1", "#c98bb3", "#8fa8e8"].map((c, i) => (
                                    <span
                                        key={c}
                                        className="grid h-9 w-9 place-items-center rounded-full border-2 border-ink text-xs font-bold text-ink"
                                        style={{ background: c }}
                                    >
                                        {"ARSK"[i]}
                                    </span>
                                ))}
                            </div>
                            <div className="text-left text-sm">
                                <div className="flex items-center gap-0.5 text-gold-400">
                                    {Array.from({ length: 5 }).map((_, i) => <Star key={i} size={14} className="fill-current" />)}
                                    <span className="ml-1.5 font-semibold text-cream">4.9</span>
                                </div>
                                <p className="text-mist">from 120K+ reviews</p>
                            </div>
                        </div>
                    </div>

                    {/* Globe */}
                    <div className="relative -mx-4 h-[360px] min-w-0 animate-fade-up [animation-delay:200ms] sm:mx-0 sm:h-[480px] lg:h-[640px]">
                        <Globe stars />

                        {/* Floating flight cards */}
                        <div className="glass absolute top-[12%] left-0 hidden animate-float rounded-2xl p-3.5 shadow-2xl sm:block lg:-left-6">
                            <div className="flex items-center gap-3">
                                <span className="grid h-9 w-9 place-items-center rounded-xl bg-gold-400/15 text-gold-300"><PlaneTakeoff size={18} /></span>
                                <div>
                                    <p className="text-sm font-semibold text-cream">CGK → HND</p>
                                    <p className="text-xs text-mist">7h 10m · Direct</p>
                                </div>
                            </div>
                        </div>
                        <div className="glass absolute right-0 bottom-[14%] hidden animate-float rounded-2xl p-3.5 shadow-2xl [animation-delay:-3s] sm:block">
                            <p className="text-[10px] font-semibold tracking-[0.18em] text-mist uppercase">Bali escape</p>
                            <p className="font-display text-2xl text-cream">IDR 899K</p>
                            <p className="text-xs text-teal-300">↓ 32% this week</p>
                        </div>
                    </div>
                </div>
            </section>

            {/* ── Search panel ── */}
            <div id="search" className="relative z-20 -mt-28 scroll-mt-28 px-4 sm:px-6 lg:px-8">
                <Reveal className="mx-auto max-w-6xl">
                    <form
                        onSubmit={(e) => e.preventDefault()}
                        className="rounded-3xl border border-white/10 bg-surface/80 p-4 shadow-[0_40px_80px_-30px_rgb(0_0_0/0.8)] backdrop-blur-2xl sm:p-6"
                    >
                        <div className="mb-5 flex gap-1 overflow-x-auto rounded-full bg-white/[0.04] p-1 sm:inline-flex">
                            {SEARCH_TABS.map(({ label, icon: Icon }) => (
                                <button
                                    key={label}
                                    type="button"
                                    onClick={() => setActiveTab(label)}
                                    className={cn(
                                        "flex flex-1 cursor-pointer items-center justify-center gap-2 rounded-full px-4 py-2 text-sm font-semibold transition-all duration-300 sm:flex-none",
                                        activeTab === label ? "bg-gold-400 text-ink shadow-lg" : "text-mist hover:text-cream",
                                    )}
                                >
                                    <Icon size={16} />
                                    {label}
                                </button>
                            ))}
                        </div>

                        <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-[1fr_auto_1fr_1fr_0.8fr_auto] lg:items-stretch">
                            <SearchField label="From" icon={<PlaneTakeoff size={13} />}>
                                <input className={inputClass} value={from} onChange={(e) => setFrom(e.target.value)} placeholder="City or airport" />
                            </SearchField>

                            <button
                                type="button"
                                onClick={() => { setFrom(to); setTo(from) }}
                                aria-label="Swap origin and destination"
                                className="mx-auto hidden h-11 w-11 cursor-pointer place-items-center self-center rounded-full border border-white/10 bg-surface-2 text-gold-300 transition-all duration-500 hover:rotate-180 hover:border-gold-400/50 lg:grid"
                            >
                                <ArrowLeftRight size={16} />
                            </button>

                            <SearchField label="To" icon={<PlaneLanding size={13} />}>
                                <input className={inputClass} value={to} onChange={(e) => setTo(e.target.value)} placeholder="Where to?" />
                            </SearchField>
                            <SearchField label="Depart" icon={<CalendarDays size={13} />}>
                                <input className={cn(inputClass, "[&::-webkit-calendar-picker-indicator]:opacity-50")} type="date" />
                            </SearchField>
                            <SearchField label="Travellers" icon={<Users size={13} />}>
                                <select className={cn(inputClass, "cursor-pointer appearance-none")} defaultValue="1">
                                    {[1, 2, 3, 4, 5, 6].map((n) => (
                                        <option key={n} value={n} className="bg-surface">{n} {n === 1 ? "Adult" : "Adults"}</option>
                                    ))}
                                </select>
                            </SearchField>

                            <Button isSubmit height="h-auto min-h-14" square="md" className="px-7 text-base sm:col-span-2 lg:col-span-1">
                                <Search size={18} />
                                Search
                            </Button>
                        </div>
                    </form>
                </Reveal>
            </div>

            {/* ── Metrics ── */}
            <section className="py-16 sm:py-20">
                <div className="mx-auto grid max-w-7xl grid-cols-1 gap-8 px-4 sm:grid-cols-2 sm:px-6 lg:grid-cols-4 lg:px-8">
                    {METRICS.map((m, i) => (
                        <Reveal key={m.title} delay={i * 100}>
                            <Metrics {...m} />
                        </Reveal>
                    ))}
                </div>
            </section>

            {/* ── Services ── */}
            <SectionWrapper
                id="services"
                eyebrow="Browse services"
                title={<>What are you <br />looking for?</>}
                className="bg-night"
            >
                <div className="grid grid-cols-2 gap-3 sm:grid-cols-3 sm:gap-4 lg:grid-cols-5">
                    {SERVICES.map(({ label, desc, icon: Icon }, i) => (
                        <Reveal key={label} delay={i * 80} className={cn(i === SERVICES.length - 1 && "col-span-2 sm:col-span-1")}>
                            <a
                                href="#search"
                                onClick={() => SEARCH_TABS.some((t) => t.label === label) && setActiveTab(label)}
                                className="group flex h-full flex-col gap-6 rounded-3xl border border-white/10 bg-surface p-5 transition-all duration-500 hover:-translate-y-1 hover:border-gold-400/40 hover:bg-surface-2 sm:p-6"
                            >
                                <span className="grid h-12 w-12 place-items-center rounded-2xl bg-gold-400/10 text-gold-300 transition-all duration-500 group-hover:bg-gold-400 group-hover:text-ink">
                                    <Icon size={22} />
                                </span>
                                <div>
                                    <p className="font-display text-xl text-cream">{label}</p>
                                    <p className="mt-1 flex items-center justify-between text-sm text-mist">
                                        {desc}
                                        <ArrowRight size={16} className="-translate-x-2 text-gold-300 opacity-0 transition-all duration-300 group-hover:translate-x-0 group-hover:opacity-100" />
                                    </p>
                                </div>
                            </a>
                        </Reveal>
                    ))}
                </div>
            </SectionWrapper>

            {/* ── Destinations ── */}
            <SectionWrapper
                id="destinations"
                eyebrow="Trending now"
                title={<>Destinations<br />worth the journey</>}
                seeAllHref="/destinations"
            >
                <div className="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
                    {TRENDING_DESTINATIONS.map((d, i) => (
                        <Reveal key={d.id} delay={i * 100}>
                            <DestinationCard {...d} />
                        </Reveal>
                    ))}
                </div>
            </SectionWrapper>

            {/* ── Offers ── */}
            <SectionWrapper
                id="deals"
                eyebrow="Limited offers"
                title={<>Deals curated<br />just for you</>}
                seeAllHref="/destinations"
                className="bg-night"
            >
                <div className="grid gap-4 lg:grid-cols-[3fr_2fr] lg:min-h-150">
                    {LIMITED_OFFERS.length > 0 && (
                        <Reveal className="h-full">
                            <OfferCard
                                cta={LIMITED_OFFERS[0].cta}
                                eyebrow={LIMITED_OFFERS[0].eyebrow}
                                title={LIMITED_OFFERS[0].title}
                                badge={LIMITED_OFFERS[0].badge}
                                description={LIMITED_OFFERS[0].description}
                                bgImage={LIMITED_OFFERS[0].background}
                                titleSize={LIMITED_OFFERS[0].type}
                                className="min-h-96"
                            />
                        </Reveal>
                    )}

                    {LIMITED_OFFERS.length > 1 && (
                        <div className="flex flex-col gap-4">
                            {LIMITED_OFFERS.slice(1).map((offer, i) => (
                                <Reveal key={offer.id} delay={(i + 1) * 120} className="flex-1">
                                    <OfferCard
                                        cta={offer.cta}
                                        eyebrow={offer.eyebrow}
                                        title={offer.title}
                                        badge={offer.badge}
                                        description={offer.description}
                                        bgImage={offer.background}
                                        titleSize={offer.type}
                                    />
                                </Reveal>
                            ))}
                        </div>
                    )}
                </div>
            </SectionWrapper>

            {/* ── CTA ── */}
            <section className="py-16 sm:py-24">
                <Reveal className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
                    <div className="relative isolate overflow-hidden rounded-[2rem] border border-gold-400/20 bg-[linear-gradient(135deg,#1c1508,#111729_60%)] px-6 py-14 text-center sm:px-12 sm:py-20">
                        <div className="absolute -top-32 left-1/2 -z-10 h-80 w-[36rem] -translate-x-1/2 rounded-full bg-gold-400/20 blur-3xl" />
                        <MapPin className="mx-auto mb-5 text-gold-400" size={28} />
                        <h2 className="mx-auto max-w-2xl font-display text-3xl leading-tight font-light text-cream sm:text-5xl">
                            Ready for your next <em className="text-gradient-gold">arc</em>?
                        </h2>
                        <p className="mx-auto mt-4 max-w-lg text-mist">
                            Create a free account to unlock member-only fares and instant price alerts.
                        </p>
                        <div className="mt-8 flex flex-col justify-center gap-3 sm:flex-row">
                            <Button height="md" className="px-8" onClick={() => navigate("/register")}>
                                Create free account <ArrowRight size={18} />
                            </Button>
                            <Button type="ghost" height="md" className="px-8" onClick={() => navigate("/login")}>
                                I already have one
                            </Button>
                        </div>
                    </div>
                </Reveal>
            </section>

            <Footer />
        </div>
    )
}

export default HomePage
