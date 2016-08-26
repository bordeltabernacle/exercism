def sum_of_multiples(n, muls):
    muls = muls[1:] if muls[0] == 0 else muls
    return sum(i for i in range(n) if any(i % m == 0 for m in muls))
