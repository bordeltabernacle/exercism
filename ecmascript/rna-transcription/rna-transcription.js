class Transcriptor {
  constructor() {
    this.translations = {
      G: 'C',
      C: 'G',
      T: 'A',
      A: 'U',
    };
  }

  transcribe(nucleotide) {
    const complement = this.translations[nucleotide];
    if (complement === undefined) {
      throw Error('Invalid input DNA.');
    }
    return complement;
  }

  toRna(dna) {
    return Array.from(dna)
      .map(x => this.transcribe(x))
      .join('');
  }
}

export default Transcriptor;
