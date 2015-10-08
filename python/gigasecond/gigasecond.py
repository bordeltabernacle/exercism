from datetime import timedelta


def add_gigasecond(birthdate):
    """
    Calculates the date that someone turned
    or will celebrate their 1 Gs anniversary
    """
    gs_birthday = birthdate + timedelta(seconds=(10**9))
    return gs_birthday
