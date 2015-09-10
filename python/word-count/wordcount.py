from collections import defaultdict


def word_count(string):
    word_dict = defaultdict(int)
    for word in string.split():
        word_dict[word] += 1
    return word_dict
