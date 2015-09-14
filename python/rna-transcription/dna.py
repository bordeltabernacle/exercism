def to_rna(dna):
    rna = ''
    for char in dna:
        if char == 'G':
            rna += 'C'
        elif char == 'C':
            rna += 'G'
        elif char == 'T':
            rna += 'A'
        elif char == 'A':
            rna += 'U'
    return rna
