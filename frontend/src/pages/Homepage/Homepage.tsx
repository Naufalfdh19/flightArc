
import Navbar from "../../components/feature/Navbar"
import heroImage from  "../../assets/hero-section.jpeg"
import Button from "../../components/ui/Button"
import Metrics from "../../components/ui/Metrics"
import { SectionWrapper } from "../../components/wrapper/SectionWrapper"
import DestinationCard from "../../components/ui/card/DestinationCard"
import OfferCard from "../../components/ui/card/OffersCard"


function HomePage() {
    return (
        <>
            <div className="bg-ink">
                <Navbar></Navbar> 
                <section 
                    className="h-screen w-full bg-cover bg-center bg-no-repeat flex flex-col items-center justify-center gap-4"
                    style={{backgroundImage:`url(${heroImage})`}}>
                        <div className="p-2 bg-black max-h-xl w-sm border-2 border-amber-300 rounded-full flex justify-center font-bold text-amber-300">
                            2M+ Travellers Trust Us
                        </div>
                        <h1 className="text-7xl max-w-4xl text-center font-bold">
                            <span className="text-amber-300 [text-shadow:1px_1px_0_rgb(0_0_0),-1px_-1px_0_rgb(0_0_0),1px_-1px_0_rgb(0_0_0),-1px_1px_0_rgb(0_0_0)]">Your World</span>, 
                            beautifully within reach
                        </h1>
                </section>
                <div className="bg-white/1 h-100 flex justify-center">
                    <div className="p-3 relative  bg-white/10 my-20 mx-10 h-45  w-full max-w-3xl border-white/20 border rounded-2xl">
                        <div className="w-full">
                            <Button type="black-op-10" border="white-1" height="sm" square="md">Flight</Button>
                        </div>
                        <div className="my-4 w-full flex gap-3 items-end">
                            <div className="w-full max-w-30 flex flex-col items-center">
                                <p className="text-white/30">FROM</p>
                                <input className="w-full border border-gray-400 rounded-md" type="text" />
                            </div>
                            <div className="w-full max-w-30 flex flex-col items-center">
                                <p className="text-white/30">TO</p>
                                <input className="w-full border border-gray-400 rounded-md" type="text" />
                            </div>
                            <div className="w-full max-w-50 flex flex-col items-center">
                                <p className="text-white/30">DEPART</p>
                                <input className="w-full border border-gray-400 rounded-md" type="text" />
                            </div>
                            <div className="absolute bottom-0 right-0 p-3">
                                <Button type="black-op-10" border="white-1" height="sm" width="sm" square="md">Search &rarr;</Button>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="w-full p-20 h-50 bg-black flex justify-between">
                    <Metrics logo="✈️">500+ Airlines <span className="text-gray-400">Worldwide coverage</span></Metrics>
                        <div className="h-10 w-px bg-white/20"></div>
                    <Metrics logo="✈️">300K+ Hotels <span className="text-gray-400">Every budget & style</span></Metrics>
                        <div className="h-10 w-px bg-white/20"></div>
                    <Metrics logo="✈️">Best Price <span className="text-gray-400">Guaranteed</span></Metrics>
                        <div className="h-10 w-px bg-white/20"></div>
                    <Metrics logo="✈️">Instant <span className="text-gray-400">E-ticket delivery</span></Metrics>
                </div>
                <SectionWrapper
                    eyebrow="BROWSE SERVICES"
                    title={<>What are you <br></br>looking for?</>}
                    seeAllHref="/destinations"
                    className="bg-ink"
                >
                    <div className="flex gap-5">
                        <Button height="sm" width="lg" square="full">✈️ Flights</Button>
                        <Button height="sm" width="lg" square="full">✈️ Flights</Button>
                        <Button height="sm" width="lg" square="full">✈️ Flights</Button>
                        <Button height="sm" width="lg" square="full">✈️ Flights</Button>
                        <Button height="sm" width="lg" square="full">✈️ Flights</Button>
                    </div>
                </SectionWrapper>
                <SectionWrapper
                    eyebrow="Trending Now"
                    title={<>Destinations<br />worth the journey</>}
                    seeAllHref="/destinations"
                    className="bg-black"
                >
                    <div className="grid grid-cols-1 gap-5 md:grid-cols-2 md:grid-rows-2 lg:grid-cols-4 lg:grid-rows-1">
                        <DestinationCard className="bg-green-900" country={"INDONESIA"} city={"Bali"} rating={"4.9"} reviews={"12K"}></DestinationCard>
                        <DestinationCard className="bg-green-900" country={"JAPAN"} city={"Tokyo"} rating={"4.8"} reviews={"8.5K"}></DestinationCard>
                        <DestinationCard className="bg-green-900" country={"FRANCE"} city={"Paris"} rating={"4.7"} reviews={"6.2K"}></DestinationCard>
                        <DestinationCard className="bg-green-900" country={"UAE"} city={"Dubai"} rating={"4.8"} reviews={"9K"}></DestinationCard>
                    </div>
                </SectionWrapper>
                <SectionWrapper
                    eyebrow="Limited Offers"
                    title={<>Deals curated<br />just for you</>}
                    seeAllHref="/destinations"
                    className="bg-ink"
                >
                    <div className="grid md:grid-cols-[3fr_2fr] gap-4 min-h-150">
                        <OfferCard link="Claim deal &rarr;" 
                            title="Flash Sale" 
                            offer="Up to 50% off international flights" 
                            offerSize="text-4xl" 
                            complementary ="Ends in 2d 14h" 
                            detail="Hundreds of routes, incredible prices. Book before they're gone."
                            className="bg-yellow-900"></OfferCard>
                        <div className="grid grid-cols-1 gap-4">
                            <OfferCard 
                                link="Explore &rarr;"
                                title="Staycation Special"
                                offer="Free breakfast included"
                                offerSize="text-2xl" 
                                complementary="Hotels" 
                                detail="500+ partner hotels, late checkout"
                                className="bg-green-900"></OfferCard>
                            <OfferCard
                                link="Browse &rarr;"
                                title="Weekend Escapes" 
                                offer="Experiences from IDR 99K"
                                offerSize="text-2xl"
                                complementary="Activities" 
                                detail="Tours, spas, local adventures"
                                className="bg-red-900"></OfferCard>
                        </div>
                    </div>
                </SectionWrapper>
            </div>
            
        </>
    )
}

export default HomePage