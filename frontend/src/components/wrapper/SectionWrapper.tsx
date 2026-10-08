import { ArrowRight } from "lucide-react"
import Reveal from "../ui/Reveal"
import { cn } from "../../utils/cn"

interface SectionWrapperProps {
  id?: string
  eyebrow: string
  title: React.ReactNode
  seeAllHref?: string
  children: React.ReactNode
  className?: string
}

export function SectionWrapper({
  id,
  eyebrow,
  title,
  seeAllHref,
  children,
  className,
}: SectionWrapperProps) {
  return (
    <section id={id} className={cn("scroll-mt-20 py-16 sm:py-24", className)}>
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <Reveal className="mb-10 flex items-end justify-between gap-6 sm:mb-14">
          <div className="flex flex-col gap-3">
            <span className="inline-flex items-center gap-2 text-xs font-semibold tracking-[0.25em] text-gold-400 uppercase">
              <span className="h-px w-6 bg-gold-400" />
              {eyebrow}
            </span>
            <h2 className="font-display text-3xl leading-[1.1] font-light text-cream sm:text-4xl lg:text-5xl">
              {title}
            </h2>
          </div>
          {seeAllHref && (
            <a
              href={seeAllHref}
              className="group hidden shrink-0 items-center gap-1.5 text-sm font-semibold text-gold-300 transition-colors hover:text-gold-200 sm:inline-flex"
            >
              See all
              <ArrowRight size={16} className="transition-transform duration-300 group-hover:translate-x-1" />
            </a>
          )}
        </Reveal>
        {children}
      </div>
    </section>
  )
}
