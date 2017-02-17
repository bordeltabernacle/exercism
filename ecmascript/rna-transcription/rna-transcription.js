const NUCLEOTIDE_COMPLEMENTS = {G:'C', C:'G', T:'A', A:'U'};

export default class Transcriptor {

  toRna(dna) {
    return [...dna].map(x => {
      let nucleotide = NUCLEOTIDE_COMPLEMENTS[x];
      if (nucleotide == undefined) {
        throw new Error('Invalid input DNA.');
      }
      return nucleotide
    }).join('');
  }

}
