def slices(s, n):
    if n > len(s) or n < 1:
        raise ValueError
    return [[int(c) for c in list(j)] for j in
            [s[i:i + n] for i in range(len(s))] if len(j) == n]


