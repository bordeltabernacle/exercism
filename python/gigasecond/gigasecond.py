from datetime import timedelta


def add_gigasecond(birthdate):
    """
    Calculates the date that someone turned
    or will celebrate their 1 Gs anniversary
    """
    return birthdate + timedelta(seconds=(10**9))
