def slices(s, n):
    if n > len(s) or n < 1:
        raise ValueError
    res = []
    for i in range(len(s)):
        sl = s[i:i + n]
        if len(sl) == n:
            sl = [int(j) for j in list(sl)]
            res.append(sl)
    return res
