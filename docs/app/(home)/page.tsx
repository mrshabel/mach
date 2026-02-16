import { CTASection } from "@/components/cta";
import { FeaturesSection } from "@/components/features";
import { HeroSection } from "@/components/hero";

export default function HomePage() {
    return (
        <main className="relative overflow-hidden">
            <HeroSection />
            <FeaturesSection />
            <CTASection />
        </main>
    );
}
