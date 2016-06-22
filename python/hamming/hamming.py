def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    if len(strandA) != len(strandB):
        raise ValueError("Strands are of unequal length.")

    pairs = zip(strandA, strandB)
    return sum(1 for (x, y) in pairs if x != y)

