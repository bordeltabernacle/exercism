def detect_anagrams(word, candidates):
    return [c for c in candidates if is_anagram(c, word)]


def is_anagram(a, b):
    return sorted(a.lower()) == sorted(b.lower()) and a.lower() != b.lower()

