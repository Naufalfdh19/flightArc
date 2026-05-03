
import React, { useState } from "react"
import Button from "../../components/Button/Button"
import { useNavigate } from "react-router-dom"
import useFetch from "../../hooks/useFetch";
import type { Base } from "../../object-types/types";
import type { LoginRequest } from "../../object-types/request/loginPage";
import { METHOD_POST } from "../../const/const";
import loginImage from  "../../assets/login-register-section.jpeg"

export default function Login() {
    const navigate = useNavigate();

    const [email, setEmail ] = useState("")
    const [password, setPassword] = useState("")

    const { data, error, isLoading, fetchData } = useFetch<Base<LoginRequest>>("http://localhost:9000/api/v1/user/auth/login")

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
        <>
            <div className="h-screen w-screen flex">
                <img className="h-screen hidden md:block" src={loginImage} alt="" />
                <div className="h-full w-full bg-black flex items-center justify-center">
                    <div className="max-h-lg max-w-2xl p-10 m:p-5">
                        <p className="text-amber-500">Welcome Back</p>
                        <p className="mt-2 mb-5 text-amber-50 text-6xl">Sign in to your account</p>
                        <span className="text-amber-50">New to ArcFlight? </span> 
                        <span 
                            onClick={() => navigate("/register")} 
                            className="text-amber-500 hover:text-amber-300 cursor-pointer">Create a free account
                        </span>
                        <div className="my-10 gap-5 flex">
                            <Button border="white-1" width="md" height="md" type="black" square="full">Google</Button>
                            <Button border="white-1" width="md" height="md" type="black" square="full">Apple</Button>
                        </div>
                        <form onSubmit={loginSubmit} className="flex flex-col gap-5">
                            <input 
                                placeholder="Email" 
                                type="text" 
                                className="h-10 p-2 bg-gray-300 rounded-sm" 
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                required
                            />

                            <input 
                                placeholder="Password" 
                                type="password" 
                                className="h-10 p-2 bg-gray-300 rounded-sm" 
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                            />
                            <Button 
                                height="sm" type="primary-700" square="sm"
                                isSubmit={true}
                                width="w-[100%]"
                            > Log In
                            </Button>
                        </form>
                    </div>
                    
                </div>
            </div>
        </>
    )
}