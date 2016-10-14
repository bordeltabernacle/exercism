import re

def abbreviate(phrase):
    findall_words = re.findall(re.compile(r"\w+\b"), phrase)
    findall_camel_case = [re.findall(re.compile(r"[A-Z][a-z]+|\w+"), r) for r in findall_words]
    acronym = "".join((word[0] for words in findall_camel_case for word in words)).upper()
    return acronym
