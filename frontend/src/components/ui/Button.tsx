import type { ReactNode } from "react";
import { BTN_HEIGHT_MD, BTN_HEIGHT_SM, BTN_HEIGHT_XS, BTN_WIDTH_MD, BTN_WIDTH_SM, BTN_WIDTH_XS } from "../../const/const";
import { cn } from "../../utils/cn";

type Variant = "primary" | "secondary" | "ghost" | "outline"
// Older names kept so existing call sites keep working
type LegacyVariant = "primary-700" | "black" | "black-op-10"

interface ButtonProps {
    children: ReactNode
    height?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    width?: "xs" | "sm" | "md" | "l" | "xl" | (string & {});
    type?: Variant | LegacyVariant;
    square?: "full" | "md" | "sm" | "none" | (string & {});
    border?: "white-1" | (string & {});
    onClick?: () => void;
    isSubmit?: boolean
    disabled?: boolean
    className?: string
    ariaLabel?: string
}

const variantClasses: Record<Variant, string> = {
    primary: "bg-gold-400 text-ink shadow-[0_8px_30px_-8px_rgb(224_176_98/0.6)] hover:bg-gold-300 hover:shadow-[0_10px_40px_-6px_rgb(224_176_98/0.75)]",
    secondary: "bg-surface text-cream hover:bg-surface-2",
    ghost: "bg-white/[0.04] text-cream hover:bg-white/10",
    outline: "bg-transparent text-cream border border-white/15 hover:border-gold-400/60 hover:text-gold-200",
}

const legacyVariants: Record<LegacyVariant, Variant> = {
    "primary-700": "primary",
    "black": "secondary",
    "black-op-10": "ghost",
}

const heightClasses: Record<string, string> = {
    xs: BTN_HEIGHT_XS,
    sm: BTN_HEIGHT_SM,
    md: BTN_HEIGHT_MD,
    l: "h-12",
    xl: "h-16",
};

const widthClasses: Record<string, string> = {
    xs: BTN_WIDTH_XS,
    sm: BTN_WIDTH_SM,
    md: BTN_WIDTH_MD,
    l: "w-40",
    xl: "w-56",
};

const squareClasses: Record<string, string> = {
    full: "rounded-full",
    md: "rounded-xl",
    sm: "rounded-lg",
    none: "rounded-none",
};

const borderClasses: Record<string, string> = {
    "white-1": "border border-white/15",
}

export default function Button({
    children,
    height,
    width,
    square,
    border,
    onClick,
    type = "primary",
    isSubmit,
    disabled,
    className,
    ariaLabel,
}: ButtonProps) {
    const variant = type in legacyVariants ? legacyVariants[type as LegacyVariant] : (type as Variant)

    // Unknown keys fall through as raw Tailwind classes (e.g. width="w-full")
    const pick = (map: Record<string, string>, key?: string) => (key ? map[key] ?? key : "")

    return (
        <button
            className={cn(
                "inline-flex cursor-pointer items-center justify-center gap-2 px-4 text-sm font-semibold whitespace-nowrap",
                "transition-all duration-300 ease-out active:scale-[0.97]",
                "focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold-300",
                "disabled:pointer-events-none disabled:opacity-50",
                variantClasses[variant],
                pick(heightClasses, height) || BTN_HEIGHT_SM,
                pick(widthClasses, width),
                pick(squareClasses, square) || "rounded-full",
                pick(borderClasses, border),
                className,
            )}
            onClick={onClick}
            type={isSubmit ? "submit" : "button"}
            disabled={disabled}
            aria-label={ariaLabel}
        >
            {children}
        </button>
    );
};
