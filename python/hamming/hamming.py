def distance(strandA, strandB):
    distance = 0
    if len(strandA) == len(strandB):
        i = 0
        while i < len(strandA):
            if strandA[i] != strandB[i]:
                distance += 1
            i += 1
    return distance

print distance('A', 'B')
