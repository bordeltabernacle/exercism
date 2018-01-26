class Transcriptor {
  translations: { [dna: string]: string } = {
    G: "C",
    C: "G",
    T: "A",
    A: "U"
  }
  toRna(dna: string): string {
    return [...dna]
      .map(x => {
        const nucleotide: string = this.translations[x]
        if (nucleotide === undefined) {
          throw Error("Invalid input DNA.")
        }
        return nucleotide
      })
      .join("")
  }
}

export default Transcriptor
