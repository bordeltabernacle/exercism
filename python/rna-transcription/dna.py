from string import maketrans


def to_rna(dna):
    '''
    Given a DNA strand, returns its RNA complement (per RNA transcription)
    '''
    key = maketrans("GCAT", "CGUA")
    return dna.translate(key)
