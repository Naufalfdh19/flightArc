import type { ReactNode } from "react";
import { BTN_HEIGHT_MD, BTN_HEIGHT_SM, BTN_HEIGHT_XS, BTN_WIDTH_MD, BTN_WIDTH_SM, BTN_WIDTH_XS } from "../../const/const";



interface ButtonProps {
    children: ReactNode
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    type?:  "dark-gold" | "black";
    square?: (string & {});
    border?: "white-1" | (string & {});
    onClick?: () => void;
    isSubmit?: boolean
}


export default function Button({
    children, 
    height,
    width, 
    square,
    border,
    onClick,
    type,
    isSubmit, 
}: ButtonProps) {
    
    const typeClasses = {
        "dark-gold": "bg-[#523a0d] text-amber-50",
        "black": "bg-black text-amber-50"
    }

    const heightClasses = {
        xs: BTN_HEIGHT_XS,
        sm: BTN_HEIGHT_SM,
        md: BTN_HEIGHT_MD,
        l: "h-12",
        xl: "h-16"
    };

    const widthClasses = {
        xs: BTN_WIDTH_XS,
        sm: BTN_WIDTH_SM,
        md: BTN_WIDTH_MD,
        l: "w-40",
        xl: "w-56"
    };

    const borderClasses = {
        "white-1": "border-1 border-amber-50"
    }

    const squareClasses = {
        full: "rounded-full",
        md: "rounded-md",
        sm: "rounded-sm",
        none: "rounded-none",
    } ;

    const btnType = type ? typeClasses[type] : "bg-[#ac8743] text-primary-white"
    const btnHeight = height ? heightClasses[height as keyof typeof heightClasses] : BTN_HEIGHT_XS;
    const btnWidth = width ? widthClasses[width as keyof typeof widthClasses] : BTN_WIDTH_XS;
    const btnBorder = border ? borderClasses[border as keyof typeof borderClasses] : "";
    const btnSquare = square ? squareClasses[square as keyof typeof squareClasses] : "";

    return (
        <>
            <button 
                className={`${btnHeight} ${btnWidth} ${btnType} ${btnSquare} ${btnBorder}`} 
                onClick={onClick}
                type={isSubmit ? "submit" : "button"}
            >
                {children}
            </button>
        </>
    );
};