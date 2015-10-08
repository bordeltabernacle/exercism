def distance(strandA, strandB):
    distance = 0
    if len(strandA) != len(strandB):
        return "Strands are of unequal length."
    for i in range(len(strandA)):
        if strandA[i] != strandB[i]:
            distance += 1
    return distance
