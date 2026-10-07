import { ArrowUpRight, Star } from "lucide-react";
import { CARD_HEIGHT_LG, CARD_HEIGHT_MD, CARD_HEIGHT_SM, CARD_HEIGHT_XL, CARD_HEIGHT_XS, CARD_WIDTH_FULL, CARD_WIDTH_LG, CARD_WIDTH_MD, CARD_WIDTH_SM, CARD_WIDTH_XL, CARD_WIDTH_XS } from "../../../const/const";
import { cn } from "../../../utils/cn";


interface CardProps {
    /** Any CSS background: image url(...) or gradient */
    background?: string
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    country: string;
    city: string;
    rating: string;
    reviews: string;
    price: string;
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

export default function DestinationCard({
    background,
    height,
    width,
    country,
    city,
    rating,
    reviews,
    price,
    className}:CardProps
) {
    const cardHeight = height ? heightClasses[height] ?? height : "h-80 sm:h-96";
    const cardWidth = width ? widthClasses[width] ?? width : CARD_WIDTH_FULL;

    return (
        <a
            href="#"
            className={cn(
                "group relative isolate flex flex-col justify-between overflow-hidden rounded-3xl border border-white/10 p-5",
                "transition-all duration-500 ease-out hover:-translate-y-1.5 hover:border-gold-400/40 hover:shadow-[0_30px_60px_-20px_rgb(224_176_98/0.35)]",
                cardHeight,
                cardWidth,
                className,
            )}
        >
            {/* Artwork */}
            <div
                className="absolute inset-0 -z-20 bg-cover bg-center transition-transform duration-700 ease-out group-hover:scale-110"
                style={{ background }}
            />
            {/* Dotted texture + bottom fade for legibility */}
            <div className="absolute inset-0 -z-10 bg-[radial-gradient(rgb(255_255_255/0.12)_1px,transparent_1px)] bg-size-[14px_14px] opacity-40 mask-[linear-gradient(to_bottom,black,transparent_70%)]" />
            <div className="absolute inset-0 -z-10 bg-linear-to-t from-ink via-ink/40 to-transparent" />
            {/* Big watermark-style city initial */}
            <span className="pointer-events-none absolute -top-6 -right-2 -z-10 font-display text-[11rem] leading-none font-light text-white/[0.07] transition-transform duration-700 group-hover:-translate-y-2">
                {city.charAt(0)}
            </span>

            <div className="flex items-start justify-between">
                <div className="glass rounded-2xl px-3.5 py-2">
                    <p className="text-[10px] font-semibold tracking-[0.18em] text-mist uppercase">From</p>
                    <p className="font-display text-lg text-cream">{price}</p>
                </div>
                <span className="grid h-10 w-10 place-items-center rounded-full bg-gold-400 text-ink opacity-0 transition-all duration-500 group-hover:rotate-0 group-hover:opacity-100 -rotate-45">
                    <ArrowUpRight size={18} />
                </span>
            </div>

            <div>
                <span className="mb-2 inline-block rounded-full border border-gold-400/40 bg-gold-400/15 px-2.5 py-0.5 text-[10px] font-semibold tracking-[0.18em] text-gold-200 uppercase">
                    {country}
                </span>
                <p className="mb-1 font-display text-3xl leading-none font-light text-cream">{city}</p>
                <p className="flex items-center gap-1 text-xs text-cream/70">
                    <Star size={12} className="fill-gold-400 text-gold-400" />
                    {rating} · {reviews} reviews
                </p>
            </div>
        </a>
    )
}
