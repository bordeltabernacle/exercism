def shouting(phrase, response):
    """ Returns appropriate response to a shouted phrase. """
    response = response
    if phrase.isupper():
        response = "Whoa, chill out!"
    return response


def silence(phrase, response):
    """ Returns appropriate response to silence. """
    response = response
    if phrase.strip() == "":
        response = "Fine. Be that way!"
    return response


def question(phrase, response):
    """ Returns appropriate response to a question. """
    response = response
    if phrase.strip() and phrase.strip()[-1] == "?":
        response = "Sure."
    return response


def hey(phrase, response="Whatever."):
    """
    Given statements, returns the answers of Bob,
    a lackadaisical teenager.
    """
    response = silence(phrase, response)
    response = question(phrase, response)
    response = shouting(phrase, response)
    return response
