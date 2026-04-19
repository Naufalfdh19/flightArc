

function Navbar() {
    return (
        <>  
            <div className="h-18 w-full flex bg-black">
                {/* Logo & Brand */}
                <div className="w-100 flex h-full items-center justify-center border-amber-50">
                    <h1 className="text-amber-50 text-2xl">FlightArc</h1>
                </div>
                {/* Menu */}
                <div className="h-full w-full px-20">
                    <div className="h-full w-full flex items-center gap-5 border-amber-50">
                        <h1 className="text-amber-50">Flights</h1>
                        <h1 className="text-amber-50">Hotels</h1>
                    </div>
                </div>
                {/* Login & Register */}
                <div className="h-full w-140 flex items-center justify-center gap-5 border-amber-50 ">
                    <button className="h-10 w-23 text-amber-50 bg-[#ac8743]  rounded-full">Log In</button>
                    <button className="h-10 w-23 text-amber-50 bg-[#523a0d]  rounded-full">Register</button>
                </div>
            </div>
        </>
    )
}


export default Navbar