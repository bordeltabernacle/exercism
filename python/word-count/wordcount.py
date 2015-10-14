from collections import Counter


def word_count(string):
    """
    Returns a count of each space separated word/symbol in a given string
    """
    return Counter(string.split())
