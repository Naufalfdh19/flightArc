
import React, { useState } from "react"
import { Link, useNavigate } from "react-router-dom"
import { ArrowRight, Lock, Mail } from "lucide-react"
import Button from "../../components/ui/Button"
import AuthLayout from "../../components/wrapper/AuthLayout"
import { Field, SocialButtons } from "../../components/ui/AuthFields"
import useFetch from "../../hooks/useFetch";
import type { Base } from "../../object-types/types";
import type { LoginRequest } from "../../object-types/request/loginPage";
import { METHOD_POST } from "../../const/const";

export default function Login() {
    const navigate = useNavigate();

    const [email, setEmail ] = useState("")
    const [password, setPassword] = useState("")

    const { data, isLoading, fetchData } = useFetch<Base<LoginRequest>>("http://localhost:9000/api/v1/user/auth/login")

    async function loginSubmit(e: React.FormEvent) {
        e.preventDefault();

        await fetchData(
            {
                method: METHOD_POST,
                headers: {

                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    email: email,
                    password: password
                })
            }
        )

        if (!data?.error) {
            navigate("/")
        }
    }

    return (
        <AuthLayout
            eyebrow="Welcome back"
            title="Sign in to your account"
            subtitle={<>
                New to FlightArc?{" "}
                <Link to="/register" className="font-semibold text-gold-300 transition-colors hover:text-gold-200">
                    Create a free account
                </Link>
            </>}
        >
            <SocialButtons />
            <form onSubmit={loginSubmit} className="flex flex-col gap-5">
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
                    placeholder="••••••••"
                    type="password"
                    autoComplete="current-password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <div className="-mt-1 flex justify-end">
                    <a href="#" className="text-sm text-mist transition-colors hover:text-gold-300">Forgot password?</a>
                </div>
                <Button height="md" square="md" isSubmit width="w-full" disabled={isLoading}>
                    {isLoading ? "Signing in…" : <>Log In <ArrowRight size={18} /></>}
                </Button>
            </form>
        </AuthLayout>
    )
}
