import string


def encode(plaintext):
    encoded = _codec(plaintext.lower())
    return " ".join(encoded[i:i+5] for i in range(0, len(encoded), 5))


def decode(cyphertext):
    return _codec(cyphertext)


def _codec(text):
    a = string.ascii_lowercase
    t = str.maketrans(a, a[::-1])
    return "".join(list(c.translate(t) if c.isalnum() else "" for c in text))
