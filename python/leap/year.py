def is_leap_year(year):
    """
    Return True if the given year is a leap year, and False if not
    """
    return not year % 4 and year % 100 != 0 or not year % 400
