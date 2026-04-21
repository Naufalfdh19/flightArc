import type { ReactNode } from "react";
import { BTN_HEIGHT_MD, BTN_HEIGHT_SM, BTN_HEIGHT_XS, BTN_WIDTH_MD, BTN_WIDTH_SM, BTN_WIDTH_XS } from "../../const/const";



interface ButtonProps {
    children: ReactNode
    height?: "xs" | "sm" | "md" | "l" | "xl"
    width?: "xs" | "sm" | "md" | "l" | "xl"
    type?:  "dark-gold" | "black"
    square?: "half" | "full"
    border?: "white-1"
    onClick?: () => void
}


export default function Button({
    children, 
    height,
    width, 
    square,
    border,
    type,
}: ButtonProps) {

    let btnType = "bg-[#ac8743] text-primary-white"
    let btnHeight = BTN_HEIGHT_XS
    let btnWidth = BTN_WIDTH_XS
    let btnBorder = ""
    let btnSquare = ""
    
    if (type === "dark-gold") {
        btnType = "bg-[#523a0d] text-amber-50"
    } else if (type === "black") {
        btnType = "bg-black text-amber-50"
    }

    if (height === "sm") {
        btnHeight = BTN_HEIGHT_SM
    } else if (height === "md") {
        btnHeight = BTN_HEIGHT_MD
    }

    if (width === "sm") {
        btnWidth = BTN_WIDTH_SM
    } else if (width === "md") {
        btnWidth = BTN_WIDTH_MD
    }

    if (square === "full") {
        btnSquare = "rounded-full"
    }

    if (border === "white-1") {
        btnBorder= "border-1 border-amber-50"
    }

    return (
        <>
            <button className={`${btnHeight} ${btnWidth} ${btnType} ${btnSquare} ${btnBorder}`}>
                {children}
            </button>
        </>
    );
};