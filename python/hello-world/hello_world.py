def hello(name=''):
    """
    Returns 'Hello, World!' by default, and
    replaces 'World' with given name if provided
    """
    name = name or 'World'
    return 'Hello, ' + name + '!'
