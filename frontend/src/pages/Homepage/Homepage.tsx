
import Navbar from "../../components/Navbar/Navbar"
import heroImage from  "../../assets/hero-section.jpeg"

function HomePage() {
    return (
        <>
            <Navbar></Navbar> 
            <section 
                className="h-screen w-screen bg-cover bg-center bg-no-repeat flex flex-col items-center justify-center gap-4"
                style={{backgroundImage:`url(${heroImage})`}}>
                    <div className="p-2 bg-black max-h-xl w-sm border-2 border-amber-300 rounded-full flex justify-center font-bold text-amber-300">
                        2M+ Travellers Trust Us
                    </div>
                    <h1 className="text-7xl max-w-4xl text-center font-bold">
                        <span className="text-amber-300 [text-shadow:1px_1px_0_rgb(0_0_0),-1px_-1px_0_rgb(0_0_0),1px_-1px_0_rgb(0_0_0),-1px_1px_0_rgb(0_0_0)]">Your World</span>, 
                        beautifully within reach
                    </h1>
            </section>
        </>
    )
}

export default HomePage