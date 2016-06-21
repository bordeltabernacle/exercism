def hey(phrase):
    """
    Given statements, returns the answers of Bob,
    a lackadaisical teenager.
    """
    if phrase.strip() == '':
        return "Fine. Be that way!"
    if phrase.isupper():
        return "Whoa, chill out!"
    if phrase.strip().endswith("?"):
        return "Sure."
    return "Whatever."
