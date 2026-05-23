import { CARD_HEIGHT_FULL, CARD_HEIGHT_LG, CARD_HEIGHT_MD, CARD_HEIGHT_SM, CARD_HEIGHT_XL, CARD_HEIGHT_XS, CARD_WIDTH_FULL, CARD_WIDTH_LG, CARD_WIDTH_MD, CARD_WIDTH_SM, CARD_WIDTH_XL, CARD_WIDTH_XS } from "../../../const/const";
import { cn } from "../../../utils/cn";
import Button from "../Button";


interface CardProps {
    bgImage?: string
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    title: string;
    offer: string;
    offerSize: string;
    detail: string;
    link: string;
    complementary: string;
    children?: React.ReactNode;
    className?: string;
}


export default function OfferCard({
    bgImage, 
    height, 
    width,
    title,
    offer,
    offerSize,
    detail,
    link,
    complementary,
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

    const cardHeight = height ? heightClasses[height as keyof typeof heightClasses] : CARD_HEIGHT_FULL;
    const cardWidth = width ? widthClasses[width as keyof typeof widthClasses] : CARD_WIDTH_FULL;


    return (
        <>  
            <div 
                className={cn(`p-5 ${cardHeight} ${cardWidth} ${className} relative rounded-3xl min-h-60`)}
                style={{ backgroundImage: `url(${bgImage})` }}
            >
                <span className="absolute right-5 inline-block bg-[#C8A96E]/25 border border-[#C8A96E]/40 text-[#C8A96E] text-[13px] tracking-[1.5px] uppercase px-2.5 py-0.5 rounded-full mb-1.5">
                    {complementary}
                </span>
                <p className={`mb-3 text-white/40`}>{title}</p>
                <h1 className={`mb-3 text-white ${offerSize}`}>{offer}</h1>
                <p className={`mb-4 text-white/40`}>{detail}</p>
                <Button square="full">{link}</Button>
            </div>
        </>
    )
}