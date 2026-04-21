
import Button from "../../components/Button/Button"



export default function Login() {
    return (
        <>
            <div className="h-screen w-screen flex">
                <div className="h-full w-1/2 bg-amber-500"></div>
                <div className="h-full w-1/2 bg-black flex items-center justify-center">
                    <div className="h-[50%] w-[75%]">
                        <p className="text-[#ac8743]">Welcome Back</p>
                        <p className="my-2 text-amber-50 text-6xl">Sign in to your account</p>
                        <span className="pt-1 text-amber-50">New to ArcFlight?</span> <span className="text-[#ac8743]">Create a free account</span>
                        <div className="my-10 gap-5 flex">
                            <Button border="white-1" width="md" height="md" type="black" square="full">Google</Button>
                            <Button border="white-1" width="md" height="md" type="black" square="full">Apple</Button>
                        </div>
                    </div>
                    
                </div>
            </div>
        </>
    )
}