class Transcriptor {
  private translations: { [key: string]: string } = {
    G: "C",
    C: "G",
    T: "A",
    A: "U"
  }

  transcribe(nucleotide: string): string {
    const complement: string = this.translations[nucleotide]
    if (!complement) {
      throw Error("Invalid input DNA.")
    }
    return complement
  }

  toRna(dnaStrand: string): string {
    return Array.from(dnaStrand).reduce((rnaStrand: string, dna: string) => {
      return `${rnaStrand}${this.transcribe(dna)}`
    }, "")
  }
}

export default Transcriptor
