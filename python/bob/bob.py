def hey(phrase):
    """
    Given statements, returns the answers of Bob,
    a lackadaisical teenager.
    """
    response = "Whatever."
    response = (lambda x, y: "Fine. Be that way!"
                if x.strip() == ''
                else y)(phrase, response)
    response = (lambda x, y: "Sure."
                if x.strip().endswith("?")
                else y)(phrase, response)
    response = (lambda x, y: "Whoa, chill out!"
                if x.isupper()
                else y)(phrase, response)

    return response
