import string


def encode(data):
    prev, encoded, count = '', '', 0
    for el in data:
        if el != prev:
            if count > 1:
                encoded += str(count)
            encoded += prev
            prev = el
            count = 0
        if el == prev:
            count += 1
    if count > 1:
        encoded += str(count)
    encoded += prev
    return encoded


def decode(data):
    decoded, count = '', ''
    for el in data:
        if el in string.digits:
            count += el
        else:
            if count == '':
                decoded += el
            else:
                decoded += (el * int(count))
                count = ''
    return decoded
