class Transcriptor {
    static translations: { [dna: string]: string } = {
        G: "C",
        C: "G",
        T: "A",
        A: "U"
    }

    transcribeNucleotide(nucleotide: string): string {
        const complement: string = Transcriptor.translations[nucleotide]
        if (!complement) {
            throw Error("Invalid input DNA.")
        }
        return complement
    }

    toRna(dna: string): string {
        return Array.from(dna)
            .map(
                (nucleotide: string): string => {
                    return this.transcribeNucleotide(nucleotide)
                }
            )
            .join("")
    }
}

export default Transcriptor
