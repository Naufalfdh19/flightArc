import React, { useState } from "react"
import { Link } from "react-router-dom"
import { ArrowRight, Lock, Mail, User } from "lucide-react"
import Button from "../../components/ui/Button"
import AuthLayout from "../../components/wrapper/AuthLayout"
import { Field, SocialButtons } from "../../components/ui/AuthFields"

export default function Register() {
    const [name, setName] = useState("")
    const [email, setEmail ] = useState("")
    const [password, setPassword] = useState("")

    async function registerSubmit(e: React.FormEvent) {
        e.preventDefault();
        console.log({name, email, password})
        const res = await fetch("http://localhost:9000/api/v1/user/auth/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                name: name,
                email: email,
                password: password,
            })
        })

        const resData = await res.json()

        console.log(resData)
    }

    return (
        <AuthLayout
            eyebrow="Get started"
            title="Create your account"
            subtitle={<>
                Already have an account?{" "}
                <Link to="/login" className="font-semibold text-gold-300 transition-colors hover:text-gold-200">
                    Log in
                </Link>
            </>}
        >
            <SocialButtons />
            <form onSubmit={registerSubmit} className="flex flex-col gap-5">
                <Field
                    label="Full name"
                    icon={<User size={18} />}
                    placeholder="Jane Traveller"
                    type="text"
                    autoComplete="name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
                <Field
                    label="Email"
                    icon={<Mail size={18} />}
                    placeholder="you@example.com"
                    type="email"
                    autoComplete="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <Field
                    label="Password"
                    icon={<Lock size={18} />}
                    placeholder="At least 8 characters"
                    type="password"
                    autoComplete="new-password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <Button height="md" square="md" isSubmit width="w-full" className="mt-2">
                    Create Account <ArrowRight size={18} />
                </Button>
                <p className="text-center text-xs text-mist/70">
                    By continuing you agree to our <a href="#" className="underline hover:text-gold-300">Terms</a> and{" "}
                    <a href="#" className="underline hover:text-gold-300">Privacy Policy</a>.
                </p>
            </form>
        </AuthLayout>
    )
}
