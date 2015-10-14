def hey(what, response=None):
    """
    Given statements, returns the answers of Bob, a lackadaisical teenager.
    """
    if what.strip() == "":
        return "Fine. Be that way!"
    elif what.isupper():
        return "Whoa, chill out!"
    elif what.strip().endswith("?"):
        return "Sure."
    else:
        return "Whatever."
