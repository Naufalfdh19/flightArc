import { useEffect, useRef } from "react"
import * as THREE from "three"
import { cn } from "../../utils/cn"

interface GlobeProps {
    className?: string
    /** Scatter a field of stars behind the globe */
    stars?: boolean
}

type City = { name: string; lat: number; lon: number }

const CITIES: Record<string, City> = {
    jakarta:   { name: "Jakarta",   lat: -6.2,  lon: 106.8 },
    bali:      { name: "Bali",      lat: -8.4,  lon: 115.2 },
    singapore: { name: "Singapore", lat: 1.35,  lon: 103.8 },
    tokyo:     { name: "Tokyo",     lat: 35.7,  lon: 139.7 },
    sydney:    { name: "Sydney",    lat: -33.9, lon: 151.2 },
    dubai:     { name: "Dubai",     lat: 25.2,  lon: 55.3 },
    paris:     { name: "Paris",     lat: 48.9,  lon: 2.35 },
    london:    { name: "London",    lat: 51.5,  lon: -0.13 },
    newYork:   { name: "New York",  lat: 40.7,  lon: -74.0 },
}

const ROUTES: [keyof typeof CITIES, keyof typeof CITIES][] = [
    ["jakarta", "tokyo"],
    ["jakarta", "dubai"],
    ["bali", "sydney"],
    ["singapore", "london"],
    ["dubai", "paris"],
    ["london", "newYork"],
    ["tokyo", "newYork"],
    ["singapore", "bali"],
]

const GOLD = new THREE.Color("#e0b062")
const GOLD_LIGHT = new THREE.Color("#f5dfb0")
const RADIUS = 1
const ARC_SEGMENTS = 96
const TRAIL = 0.22 // fraction of the arc lit by the moving comet

function latLonToVector(lat: number, lon: number, radius = RADIUS) {
    const phi = THREE.MathUtils.degToRad(90 - lat)
    const theta = THREE.MathUtils.degToRad(lon + 180)
    return new THREE.Vector3(
        -radius * Math.sin(phi) * Math.cos(theta),
        radius * Math.cos(phi),
        radius * Math.sin(phi) * Math.sin(theta),
    )
}

/** Evenly spread points over a sphere (Fibonacci lattice). */
function fibonacciSphere(count: number, radius: number) {
    const positions = new Float32Array(count * 3)
    const golden = Math.PI * (3 - Math.sqrt(5))
    for (let i = 0; i < count; i++) {
        const y = 1 - (i / (count - 1)) * 2
        const r = Math.sqrt(1 - y * y)
        const t = golden * i
        positions.set([Math.cos(t) * r * radius, y * radius, Math.sin(t) * r * radius], i * 3)
    }
    return positions
}

export default function Globe({ className, stars = false }: GlobeProps) {
    const mountRef = useRef<HTMLDivElement>(null)

    useEffect(() => {
        const mount = mountRef.current
        if (!mount) return

        let renderer: THREE.WebGLRenderer
        try {
            renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true, powerPreference: "low-power" })
        } catch {
            return // No WebGL — the surrounding gradient still looks fine.
        }

        const reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches
        const isSmall = window.innerWidth < 768

        renderer.setPixelRatio(Math.min(window.devicePixelRatio, isSmall ? 1.5 : 2))
        renderer.domElement.style.display = "block"
        mount.appendChild(renderer.domElement)

        const scene = new THREE.Scene()
        const camera = new THREE.PerspectiveCamera(42, 1, 0.1, 100)
        camera.position.set(0, 0, 3.4)

        const disposables: { dispose: () => void }[] = []
        const track = <T extends { dispose: () => void }>(item: T) => {
            disposables.push(item)
            return item
        }

        // Tilted group holds everything that spins with the earth
        const tilt = new THREE.Group()
        tilt.rotation.z = THREE.MathUtils.degToRad(-12)
        scene.add(tilt)
        const globe = new THREE.Group()
        // Face South-East Asia toward the camera first
        globe.rotation.y = THREE.MathUtils.degToRad(-90 - 110)
        globe.rotation.x = THREE.MathUtils.degToRad(12)
        tilt.add(globe)

        // Solid core so the back-side dots are hidden
        globe.add(new THREE.Mesh(
            track(new THREE.SphereGeometry(RADIUS * 0.985, 64, 64)),
            track(new THREE.MeshBasicMaterial({ color: "#0b1020" })),
        ))

        // Dotted surface
        const dotGeometry = track(new THREE.BufferGeometry())
        dotGeometry.setAttribute("position", new THREE.BufferAttribute(fibonacciSphere(isSmall ? 2200 : 4200, RADIUS), 3))
        globe.add(new THREE.Points(dotGeometry, track(new THREE.PointsMaterial({
            color: "#8d97b5",
            size: 0.011,
            transparent: true,
            opacity: 0.55,
            sizeAttenuation: true,
        }))))

        // Latitude rings
        const ringMaterial = track(new THREE.LineBasicMaterial({ color: "#2a3558", transparent: true, opacity: 0.6 }))
        for (const lat of [-60, -30, 0, 30, 60]) {
            const pts: THREE.Vector3[] = []
            for (let lon = -180; lon <= 180; lon += 4) pts.push(latLonToVector(lat, lon, RADIUS * 1.001))
            globe.add(new THREE.Line(track(new THREE.BufferGeometry().setFromPoints(pts)), ringMaterial))
        }

        // Atmosphere glow (fresnel on a back-facing shell)
        const atmosphere = new THREE.Mesh(
            track(new THREE.SphereGeometry(RADIUS * 1.18, 64, 64)),
            track(new THREE.ShaderMaterial({
                uniforms: { glowColor: { value: GOLD } },
                vertexShader: /* glsl */ `
                    varying vec3 vNormal;
                    void main() {
                        vNormal = normalize(normalMatrix * normal);
                        gl_Position = projectionMatrix * modelViewMatrix * vec4(position, 1.0);
                    }`,
                fragmentShader: /* glsl */ `
                    uniform vec3 glowColor;
                    varying vec3 vNormal;
                    void main() {
                        float intensity = pow(0.62 - dot(vNormal, vec3(0.0, 0.0, 1.0)), 3.0);
                        gl_FragColor = vec4(glowColor, 1.0) * intensity * 0.6;
                    }`,
                side: THREE.BackSide,
                blending: THREE.AdditiveBlending,
                transparent: true,
                depthWrite: false,
            })),
        )
        scene.add(atmosphere)

        // City markers with pulsing rings
        const markerGeometry = track(new THREE.SphereGeometry(0.014, 12, 12))
        const markerMaterial = track(new THREE.MeshBasicMaterial({ color: GOLD_LIGHT }))
        const pulseGeometry = track(new THREE.RingGeometry(0.018, 0.026, 32))
        const pulses: { mesh: THREE.Mesh; material: THREE.MeshBasicMaterial; offset: number }[] = []

        Object.values(CITIES).forEach((city, i) => {
            const position = latLonToVector(city.lat, city.lon, RADIUS * 1.003)
            const marker = new THREE.Mesh(markerGeometry, markerMaterial)
            marker.position.copy(position)
            globe.add(marker)

            const material = track(new THREE.MeshBasicMaterial({
                color: GOLD,
                transparent: true,
                side: THREE.DoubleSide,
                depthWrite: false,
            }))
            const pulse = new THREE.Mesh(pulseGeometry, material)
            pulse.position.copy(position)
            pulse.lookAt(position.clone().multiplyScalar(2))
            globe.add(pulse)
            pulses.push({ mesh: pulse, material, offset: i * 0.37 })
        })

        // Flight arcs: a faint full path plus a bright moving comet and plane dot
        const planeGeometry = track(new THREE.SphereGeometry(0.012, 10, 10))
        const planeMaterial = track(new THREE.MeshBasicMaterial({ color: "#ffffff" }))
        const arcs = ROUTES.map(([from, to], i) => {
            const start = latLonToVector(CITIES[from].lat, CITIES[from].lon)
            const end = latLonToVector(CITIES[to].lat, CITIES[to].lon)
            const lift = RADIUS + 0.12 + start.distanceTo(end) * 0.32
            const mid = start.clone().add(end).normalize().multiplyScalar(lift)
            const curve = new THREE.QuadraticBezierCurve3(start, mid, end)
            const points = curve.getPoints(ARC_SEGMENTS)

            const path = new THREE.Line(
                track(new THREE.BufferGeometry().setFromPoints(points)),
                track(new THREE.LineBasicMaterial({ color: GOLD, transparent: true, opacity: 0.18 })),
            )
            globe.add(path)

            const cometGeometry = track(new THREE.BufferGeometry().setFromPoints(points))
            const comet = new THREE.Line(
                cometGeometry,
                track(new THREE.LineBasicMaterial({ color: GOLD_LIGHT, transparent: true, opacity: 0.95 })),
            )
            globe.add(comet)

            const plane = new THREE.Mesh(planeGeometry, planeMaterial)
            globe.add(plane)

            return { curve, cometGeometry, plane, speed: 0.11 + (i % 3) * 0.025, offset: i / ROUTES.length }
        })

        // Optional star field
        if (stars) {
            const count = isSmall ? 300 : 700
            const positions = new Float32Array(count * 3)
            for (let i = 0; i < count; i++) {
                const dir = new THREE.Vector3().randomDirection().multiplyScalar(6 + Math.random() * 6)
                positions.set([dir.x, dir.y, dir.z], i * 3)
            }
            const starGeometry = track(new THREE.BufferGeometry())
            starGeometry.setAttribute("position", new THREE.BufferAttribute(positions, 3))
            scene.add(new THREE.Points(starGeometry, track(new THREE.PointsMaterial({
                color: "#cfd6ea",
                size: 0.025,
                transparent: true,
                opacity: 0.6,
            }))))
        }

        // Sizing
        const resize = () => {
            const { clientWidth: w, clientHeight: h } = mount
            if (!w || !h) return
            renderer.setSize(w, h)
            camera.aspect = w / h
            // Keep the whole globe in frame on tall/narrow containers
            camera.position.z = w / h < 1 ? 3.4 / (w / h) ** 0.75 : 3.4
            camera.updateProjectionMatrix()
        }
        const resizeObserver = new ResizeObserver(resize)
        resizeObserver.observe(mount)
        resize()

        // Gentle parallax that follows the pointer
        const pointer = { x: 0, y: 0 }
        const onPointerMove = (e: PointerEvent) => {
            pointer.x = (e.clientX / window.innerWidth) * 2 - 1
            pointer.y = (e.clientY / window.innerHeight) * 2 - 1
        }
        if (!reduceMotion) window.addEventListener("pointermove", onPointerMove)

        const renderFrame = (elapsed: number, delta: number) => {
            if (!reduceMotion) {
                globe.rotation.y += delta * 0.06
                tilt.rotation.x += (pointer.y * 0.18 - tilt.rotation.x) * 0.04
                tilt.rotation.y += (pointer.x * 0.25 - tilt.rotation.y) * 0.04
            }

            for (const { mesh, material, offset } of pulses) {
                const p = ((reduceMotion ? 0.3 : elapsed * 0.5) + offset) % 1
                mesh.scale.setScalar(1 + p * 2.4)
                material.opacity = (1 - p) * 0.8
            }

            for (const arc of arcs) {
                const t = reduceMotion ? 1 : ((elapsed * arc.speed + arc.offset) % 1) * (1 + TRAIL)
                const head = Math.min(t, 1)
                const tail = Math.max(t - TRAIL, 0)
                const startIndex = Math.floor(tail * ARC_SEGMENTS)
                arc.cometGeometry.setDrawRange(startIndex, Math.max(Math.ceil(head * ARC_SEGMENTS) - startIndex + 1, 0))
                arc.plane.visible = t <= 1 && !reduceMotion
                arc.plane.position.copy(arc.curve.getPoint(head))
            }

            renderer.render(scene, camera)
        }

        // Only animate while on screen
        let frame = 0
        let visible = true
        const timer = new THREE.Timer()
        const loop = (time: number) => {
            frame = requestAnimationFrame(loop)
            if (!visible) return
            timer.update(time)
            renderFrame(timer.getElapsed(), Math.min(timer.getDelta(), 0.05))
        }
        const visibilityObserver = new IntersectionObserver(([entry]) => {
            visible = entry.isIntersecting
        })
        visibilityObserver.observe(mount)

        // With reduced motion renderFrame draws a still scene, but keeps redrawing after resizes
        frame = requestAnimationFrame(loop)

        return () => {
            cancelAnimationFrame(frame)
            resizeObserver.disconnect()
            visibilityObserver.disconnect()
            window.removeEventListener("pointermove", onPointerMove)
            disposables.forEach((d) => d.dispose())
            renderer.dispose()
            renderer.domElement.remove()
        }
    }, [stars])

    return <div ref={mountRef} aria-hidden="true" className={cn("h-full w-full", className)} />
}
