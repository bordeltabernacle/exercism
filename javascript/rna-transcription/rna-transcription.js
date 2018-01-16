class DnaTranscriber {
  toRna(dna) {
    return [...dna]
      .map(x => {
        var nucleotide = { G: "C", C: "G", T: "A", A: "U" }[x];
        if (nucleotide === undefined) {
          throw Error("Invalid input");
        }
        return nucleotide;
      })
      .join("");
  }
}

module.exports = DnaTranscriber;
