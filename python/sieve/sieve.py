def sieve(limit):
    return _shake_sieve(range(2, limit + 1), [])


def _shake_sieve(r, bowl):
    head = r[0]
    bowl.append(head)
    r = [x for x in r if x % head != 0]
    if r:
        _shake_sieve(r, bowl)
    return bowl
