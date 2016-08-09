def detect_anagrams(w, cs):
    return [c for c in cs if sorted(c.lower()) == sorted(w.lower()) and c.lower() != w.lower()]
