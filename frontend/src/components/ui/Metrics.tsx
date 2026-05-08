import {type ReactNode } from "react"

interface MetricsProps {
    logo: string
    children: ReactNode
    classname?: string | undefined
}

export default function Metrics(props: MetricsProps) {
    return (
        <>
            <div className={`h-full flex items-center justify-center ${props.classname}`}>
                <p className="p-1 px-2 text-white">{props.logo}</p>
                <p className="p-1 px-2 text-white w-25">{props.children}</p>
            </div>
            
        </>
    )
}