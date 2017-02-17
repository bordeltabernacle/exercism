const NUCLEOTIDE_COMPLEMENTS = {'G':'C', 'C':'G', 'T':'A', 'A':'U'};

export default class Transcriptor {

  toRna(dna) {
    let rna = '';
    for (var i = 0; i < dna.length; i++) {
      let nucleotide = NUCLEOTIDE_COMPLEMENTS[dna[i]];
      if (nucleotide == undefined) {
        throw new Error('Invalid input DNA.');
      }
      rna += nucleotide;
    }
    return rna
  }

}
