import { useNavigate } from "react-router-dom"
import Button from "../../ui/Button"

function HomeNavbar() {

    const navigate = useNavigate();

    return (
        <>  
            <div className="h-18 w-full flex bg-white/1">
                <div className="w-100 flex h-full items-center justify-center border-amber-50">
                    <h1 className="text-amber-50 text-2xl">FlightArc</h1>
                </div>
                <div className="h-full w-full px-20">
                    <div className="h-full w-full flex items-center gap-5 border-amber-50">
                        <h1 className="text-amber-50">Flights</h1>
                        <h1 className="text-amber-50">Hotels</h1>
                    </div>
                </div>
                <div className="h-full w-140 flex items-center justify-center gap-5 border-amber-50 ">
                    <Button width="xs" onClick={() => navigate("/login")} square="full">
                        Log In
                    </Button>
                    <Button width="xs" onClick={() => navigate("/register")} type="primary-700" square="full">
                        Register
                    </Button>
                </div>
            </div>
        </>
    )
}


export default HomeNavbar