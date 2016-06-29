from itertools import groupby


def encode(data):
    encoded = ''
    for k, v in groupby(data):
        count = len(list(v))
        if count > 1:
            encoded += '{1}{0}'.format(k, count)
        else:
            encoded += k
    return encoded


def decode(data):
    decoded, count = '', 1
    for k, v in groupby(data, str.isdigit):
        elem = "".join(v)
        if k:
            count = int(elem)
        else:
            print(elem)
            decoded += '{0}{1}'.format(elem[0] * count, elem[1:])
    return decoded
