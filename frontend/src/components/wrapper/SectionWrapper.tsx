interface SectionWrapperProps {
  eyebrow: string
  title: React.ReactNode
  seeAllHref?: string
  children: React.ReactNode
  className?: string
}

export function SectionWrapper({
  eyebrow,
  title,
  seeAllHref,
  children,
  className,
}: SectionWrapperProps) {
  return (
    <section className={`px-13 py-20 ${className}`}>
      <div className="flex justify-between items-end mb-11">
        <div className="flex flex-col gap-1">
          <span className="text-[10px] tracking-[2.5px] uppercase text-[#C8A96E]">
            {eyebrow}
          </span>
          <h2 className="font-serif text-[38px] font-light leading-tight text-white">
            {title}
          </h2>
        </div>
        {seeAllHref && (
          <a href={seeAllHref} className="text-[13px] text-[#C8A96E] border-b border-transparent hover:border-[#C8A96E] transition-all mt-1">
            See all →
          </a>
        )}
      </div>
      {children}
    </section>
  )
}