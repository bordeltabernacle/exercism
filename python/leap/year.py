def is_leap_year(year):
    """
    Returns True if the given year is a leap year, and False if not
    """
    if year % 400 == 0:
        return True
    if year % 100 == 0:
        return False
    if year % 4 != 0:
        return False
    else:
        return True
