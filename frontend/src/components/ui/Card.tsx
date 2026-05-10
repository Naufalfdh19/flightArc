import { CARD_HEIGHT_LG, CARD_HEIGHT_MD, CARD_HEIGHT_SM, CARD_HEIGHT_XL, CARD_HEIGHT_XS, CARD_WIDTH_LG, CARD_WIDTH_MD, CARD_WIDTH_SM, CARD_WIDTH_XL, CARD_WIDTH_XS } from "../../const/const";
import { cn } from "../../utils/cn";


interface CardProps {
    bgImage?: string
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    children?: React.ReactNode;
    className?: string;
}


export default function Card({
    bgImage, 
    height, 
    width,
    children,
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

    const cardHeight = height ? heightClasses[height as keyof typeof heightClasses] : CARD_HEIGHT_XS;
    const cardWidth = width ? widthClasses[width as keyof typeof widthClasses] : CARD_WIDTH_XS;


    return (
        <>  
            <div 
                className={cn(`p-5 ${cardHeight} ${cardWidth} ${className}`)}
                style={{ backgroundImage: `url(${bgImage})` }}
            >{children}</div>
        </>
    )
}