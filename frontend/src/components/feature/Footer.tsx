import Button from "../ui/Button";





export default function Footer() {
    return (
        <>
            <section className={`px-13 py-20 bg-black`}>
                <div className="flex justify-between">
                    <div className="mr-20">
                        <h3 className="mb-5 text-white text-2xl">WANDERA</h3>
                        <p className="text-white/40">Your all-in-one luxury travel companion. </p>
                        <p className="mb-5 text-white/40">Book smarter and journey beautifully.</p>
                        <div className="flex gap-3">
                            <Button type="black" square="sm" border="white-1" width="max-w-5">f</Button>
                            <Button type="black" square="sm" border="white-1">in</Button>
                            <Button type="black" square="sm" border="white-1">X</Button>
                            <Button type="black" square="sm" border="white-1">s</Button>
                        </div>
                    </div>
                    <div className="flex w-full max-w-130 justify-between">
                        <div>
                            <h4 className="text-white/20">Products</h4>
                            <p className="text-white/40">Flights</p>
                            <p className="text-white/40">hotels</p>
                            <p className="text-white/40">Trains</p>
                            <p className="text-white/40">Car Rentals</p>
                            <p className="text-white/40">Packages</p>
                        </div>
                        <div>
                            <h4 className="text-white/20">Company</h4>
                            <p className="text-white/40">About Us</p>
                            <p className="text-white/40">Careers</p>
                            <p className="text-white/40">Newsroom</p>
                            <p className="text-white/40">Investors</p>
                        </div>
                        <div>
                            <h4 className="text-white/20">Support</h4>
                            <p className="text-white/40">Help Center</p>
                            <p className="text-white/40">Contact Us</p>
                            <p className="text-white/40">Privacy Policy</p>
                            <p className="text-white/40">Terms</p>
                        </div>
                    </div>
                    
                </div>
                <div>

                </div>
            </section>     
        </>
    )
}