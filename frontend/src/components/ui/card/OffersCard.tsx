import { ArrowRight } from "lucide-react";
import { CARD_HEIGHT_FULL, CARD_HEIGHT_LG, CARD_HEIGHT_MD, CARD_HEIGHT_SM, CARD_HEIGHT_XL, CARD_HEIGHT_XS, CARD_WIDTH_FULL, CARD_WIDTH_LG, CARD_WIDTH_MD, CARD_WIDTH_SM, CARD_WIDTH_XL, CARD_WIDTH_XS } from "../../../const/const";
import { cn } from "../../../utils/cn";
import Button from "../Button";


interface CardProps {
    bgImage?: string
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    eyebrow: string;
    title: string;
    titleSize: "small" | "big" | (string & {});
    description: string;
    cta: string;
    badge: string;
    className?: string;
}

const heightClasses: Record<string, string> = {
    xs: CARD_HEIGHT_XS,
    sm: CARD_HEIGHT_SM,
    md: CARD_HEIGHT_MD,
    l: CARD_HEIGHT_LG,
    xl: CARD_HEIGHT_XL
};

const widthClasses: Record<string, string> = {
    xs: CARD_WIDTH_XS,
    sm: CARD_WIDTH_SM,
    md: CARD_WIDTH_MD,
    l: CARD_WIDTH_LG,
    xl: CARD_WIDTH_XL
};

const titleSizeClasses: Record<string, string> = {
    "big": "text-3xl sm:text-4xl lg:text-5xl",
    "small": "text-2xl sm:text-3xl"
};

export default function OfferCard({
    bgImage,
    height,
    width,
    eyebrow,
    title,
    titleSize,
    description,
    cta,
    badge,
    className}:CardProps
) {
    const cardHeight = height ? heightClasses[height] ?? height : CARD_HEIGHT_FULL;
    const cardWidth = width ? widthClasses[width] ?? width : CARD_WIDTH_FULL;
    const cardTitleSize = titleSizeClasses[titleSize] ?? "text-2xl";
    const isBig = titleSize === "big"

    return (
        <div
            className={cn(
                "group relative isolate flex min-h-64 flex-col overflow-hidden rounded-3xl border border-white/10 p-6 sm:p-8",
                "transition-all duration-500 hover:border-gold-400/30",
                isBig ? "justify-end" : "justify-between",
                cardHeight,
                cardWidth,
                className,
            )}
            style={{ backgroundImage: bgImage }}
        >
            {/* Glowing orb drifts on hover */}
            <div className="absolute -top-24 -right-24 -z-10 h-72 w-72 rounded-full bg-gold-400/20 blur-3xl transition-transform duration-700 group-hover:-translate-x-10 group-hover:translate-y-10" />
            {isBig && (
                <svg className="absolute top-10 right-0 -z-10 h-2/3 w-2/3 text-gold-400/25" viewBox="0 0 200 200" fill="none" aria-hidden="true">
                    <circle cx="140" cy="60" r="110" stroke="currentColor" strokeDasharray="2 6" />
                    <circle cx="140" cy="60" r="70" stroke="currentColor" strokeDasharray="2 6" />
                    <path d="M20 170 Q 100 20 190 70" stroke="currentColor" strokeWidth="1.5" />
                </svg>
            )}

            <span className="glass absolute top-6 right-6 rounded-full px-3 py-1 text-[11px] font-semibold tracking-wider text-gold-200 uppercase sm:top-8 sm:right-8">
                {badge}
            </span>

            <div className={cn(!isBig && "pr-24")}>
                <p className="mb-3 text-xs font-semibold tracking-[0.2em] text-gold-400 uppercase">{eyebrow}</p>
                <h3 className={cn("mb-3 font-display leading-tight font-light whitespace-pre-line text-cream", cardTitleSize)}>{title}</h3>
                <p className="mb-6 text-mist">{description}</p>
            </div>
            <div>
                <Button type={isBig ? "primary" : "outline"} height={isBig ? "md" : "sm"}>
                    {cta}
                    <ArrowRight size={16} className="transition-transform duration-300 group-hover:translate-x-1" />
                </Button>
            </div>
        </div>
    )
}
