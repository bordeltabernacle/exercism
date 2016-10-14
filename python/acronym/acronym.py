import re

def abbreviate(phrase):
    pattern = "[A-Z]+[a-z]*|[a-z]+"
    return "".join(word[0] for word in re.findall(pattern, phrase)).upper()
