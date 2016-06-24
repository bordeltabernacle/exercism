import string


def is_pangram(sentence):
    alphabet = list(string.ascii_lowercase)
    test = sum(1 for letter in alphabet if letter in sentence.lower())
    result = True if test == 26 else False
    return result
