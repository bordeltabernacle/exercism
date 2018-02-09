class SpaceAge {
    readonly earthYearInSeconds: number = 365.25 * 24 * 60 * 60
    constructor(readonly seconds: number) {
        this.seconds = seconds
    }
    private years(orbitalPeriod: number): number {
        return (
            Math.round(
                this.seconds / (this.earthYearInSeconds * orbitalPeriod) * 100
            ) / 100
        )
    }
    onEarth(): number {
        return this.years(1)
    }
    onMercury(): number {
        return this.years(0.2408467)
    }
    onVenus(): number {
        return this.years(0.61519726)
    }
    onMars(): number {
        return this.years(1.8808158)
    }
    onJupiter(): number {
        return this.years(11.862615)
    }
    onSaturn(): number {
        return this.years(29.447498)
    }
    onUranus(): number {
        return this.years(84.016846)
    }
    onNeptune(): number {
        return this.years(164.79132)
    }
}

export default SpaceAge
