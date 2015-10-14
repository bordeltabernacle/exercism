def hey(what, response=None):
    """
    Given statements, returns the answers of Bob, a lackadaisical teenager.
    """
    response_dict = {"stating": "Whatever.",
                     "shouting": "Whoa, chill out!",
                     "question": "Sure.",
                     "silence": "Fine. Be that way!"}
    if what.strip() == "":
        response = response_dict["silence"]
    elif what.isupper():
        response = response_dict["shouting"]
    elif what.strip().endswith("?"):
        response = response_dict["question"]
    else:
        response = response_dict["stating"]
    return response
