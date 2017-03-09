def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    return sum(x != y for (x, y) in zip(strandA, strandB))
