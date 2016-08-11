import string


ALPHABET = list(string.ascii_lowercase)


def encode(plaintext):
    encoded = ""
    for c in plaintext.lower():
        if c == " " or c in string.punctuation:
            pass
        elif c in string.digits:
            encoded += c
        else:
            i = ALPHABET.index(c)
            encoded += ALPHABET[-(i + 1)]
    encoded = " ".join(encoded[i:i+5] for i in range(0, len(encoded), 5))
    return encoded


def decode(cyphertext):
    decoded = ""
    for c in cyphertext.lower():
        if c == " ":
            pass
        else:
            i = ALPHABET.index(c)
            decoded += ALPHABET[-(i + 1)]
    return decoded
