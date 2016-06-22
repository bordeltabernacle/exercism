def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    if len(strandA) != len(strandB):
        raise ValueError("Strands are of unequal length.")

    combined = zip(strandA, strandB)
    differ = filter(different, combined)
    return len(list(differ))


def different(pair):
    """
    Asserts whether the pair in a tuple are different
    """
    (x, y) = pair
    return x != y
