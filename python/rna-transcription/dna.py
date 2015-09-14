def to_rna(dna):
    '''
    Given a DNA strand, returns its RNA complement (per RNA transcription)
    '''
    dna_to_rna = {'G': 'C', 'C': 'G', 'A': 'U', 'T': 'A'}
    return ''.join(dna_to_rna.get(char) for char in dna)
