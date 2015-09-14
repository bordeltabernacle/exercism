def to_rna(dna):
    '''
    Given a DNA strand, returns its RNA complement (per RNA transcription)
    '''
    rna = ''
    dna_to_rna = {'G': 'C', 'C': 'G', 'A': 'U', 'T': 'A'}
    for char in dna:
        rna += dna_to_rna.get(char)
    return rna
