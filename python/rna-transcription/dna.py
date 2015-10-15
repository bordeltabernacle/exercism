from string import maketrans


def to_rna(dna):
    '''
    Given a DNA strand, returns its RNA complement (per RNA transcription)
    '''
    return dna.translate(maketrans("GCAT", "CGUA"))

