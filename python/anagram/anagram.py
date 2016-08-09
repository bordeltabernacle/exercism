def detect_anagrams(word, candidates):
    return [c for c in candidates if is_anagram(c, word)]


def is_anagram(a, b):
    a, b = a.lower(), b.lower()
    return sorted(a) == sorted(b) and a != b

