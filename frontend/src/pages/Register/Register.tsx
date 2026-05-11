
import React, { useState } from "react"
import Button from "../../components/ui/Button"
import { useNavigate } from "react-router-dom"

export default function Register() {
    const navigate = useNavigate();

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
        <>
            <div className="h-screen w-screen flex">
                <div className="h-full w-1/2 bg-amber-500"></div>
                <div className="h-full w-1/2 bg-black flex items-center justify-center">
                    <div className="h-[75%] w-[75%]">
                        <p className="text-amber-500">Welcome Back</p>
                        <p className="mt-2 mb-5 text-amber-50 text-6xl">Sign up your account</p>
                        <span className="text-amber-50">Already have an account? </span> 
                        <span 
                            onClick={() => navigate("/login")} 
                            className="text-amber-500 cursor-pointer hover:text-amber-300 transition-all">Log In
                        </span>
                        <div className="my-10 gap-5 flex">
                            <Button border="white-1" width="md" height="md" type="black" square="full">Google</Button>
                            <Button border="white-1" width="md" height="md" type="black" square="full">Apple</Button>
                        </div>
                        <form onSubmit={registerSubmit} className="flex flex-col gap-5">
                            <input 
                                placeholder="name" 
                                type="text" 
                                className="h-10 p-2 bg-gray-300 rounded-sm" 
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                required
                            />
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
                            > Create Account
                            </Button>
                        </form>
                    </div>
                    
                </div>
            </div>
        </>
    )
}