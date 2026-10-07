import { lazy, Suspense, type ComponentProps } from "react"

// three.js is heavy — load it in its own chunk so text and layout paint first
const Globe = lazy(() => import("./Globe"))

export default function LazyGlobe(props: ComponentProps<typeof Globe>) {
    return (
        <Suspense fallback={null}>
            <Globe {...props} />
        </Suspense>
    )
}
