def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    distance = 0

    if len(strandA) != len(strandB):
        raise ValueError("Strands are of unequal length.")

    for x, y in zip(strandA, strandB):
        if x != y:
            distance += 1

    return distance
