def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    if len(strandA) != len(strandB):
        raise ValueError("Strands are of unequal length.")

    pairs = zip(strandA, strandB)
    return sum([1 for pair in pairs if different(pair)])


def different(pair):
    """
    Asserts whether the pair in a tuple are different
    """
    (x, y) = pair
    return x != y

