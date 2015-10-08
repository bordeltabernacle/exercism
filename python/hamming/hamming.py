def distance(strandA, strandB):
    """
    Calculates the Hamming difference between two DNA strands
    """
    distance = 0
    if len(strandA) != len(strandB):
        return "Strands are of unequal length."
    for i in range(len(strandA)):
        if strandA[i] != strandB[i]:
            distance += 1
    return distance
