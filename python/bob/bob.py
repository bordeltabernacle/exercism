def shouting(phrase):
    """ Determines if the given phrase is shouted """
    return phrase.isupper()


def silence(phrase):
    """ Determines if the given phrase is empty """
    return phrase.strip() == ""


def question(phrase):
    """ Determines if the given phrase is a question """
    return phrase.strip()[-1] == "?"


def hey(phrase):
    """ Given statements, returns the answers of Bob, a lackadaisical teenager. """
    response = "Whatever."
    if silence(phrase):
        response = "Fine. Be that way!"
    elif shouting(phrase):
        response = "Whoa, chill out!"
    elif question(phrase):
        response = "Sure."
    return response
