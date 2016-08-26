def sum_of_multiples(n, muls):
    res = []
    for i in range(n):
        for m in muls:
            if m != 0 and i % m == 0:
                res.append(i)
    return sum(set(res))
