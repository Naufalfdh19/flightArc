import { CARD_HEIGHT_LG, CARD_HEIGHT_MD, CARD_HEIGHT_SM, CARD_HEIGHT_XL, CARD_HEIGHT_XS, CARD_WIDTH_LG, CARD_WIDTH_MD, CARD_WIDTH_SM, CARD_WIDTH_XL, CARD_WIDTH_XS } from "../../../const/const";
import { cn } from "../../../utils/cn";


interface CardProps {
    bgImage?: string
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    country: (string & {});
    city: (string & {});
    rating: (string & {});
    reviews: (string & {});
    children?: React.ReactNode;
    className?: string;
}


export default function DestinationCard({
    bgImage, 
    height, 
    width,
    country,
    city,
    rating,
    reviews,
    className}:CardProps
) {

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

    const cardHeight = height ? heightClasses[height as keyof typeof heightClasses] : CARD_HEIGHT_MD;
    const cardWidth = width ? widthClasses[width as keyof typeof widthClasses] : CARD_WIDTH_XS;


    return (
        <>  
            <div 
                className={cn(`p-5 ${cardHeight} ${cardWidth} ${className} flex flex-col items-end justify-between`)}
                style={{ backgroundImage: `url(${bgImage})` }}
            >
                <div className="bg-black w-12.5 h-5"></div>
                <div className="bg-black w-full h-40">
                    <span className="inline-block bg-[#C8A96E]/25 border border-[#C8A96E]/40 text-[#C8A96E] text-[9px] tracking-[1.5px] uppercase px-2.5 py-0.5 rounded-full mb-1.5">
                        {country}
                    </span>
                    {/* City name */}
                    <p className="font-serif text-[22px] font-light text-white leading-none mb-0.5">
                        {city}
                    </p>
                    {/* Rating */}
                    <p className="text-[11px] text-white/55">
                        ★ {rating} · {reviews} reviews
                    </p>
                </div>
            </div>
        </>
    )
}