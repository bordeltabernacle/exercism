from datetime import date
import calendar


WEEK = {
        "1st": 1,
        "2nd": 8,
        "3rd": 15,
        "4th": 22,
        "5th": 28,
        "teenth": 13
}


def meetup_day(year, month, dotw, when):
    DOTW = [d for d in calendar.day_name]
    (_, days) = calendar.monthrange(year, month)
    if when == "last":
        return _get_day(year, month, DOTW.index(dotw), reversed(range(days - 6, days + 1)))
    else:
        return _get_day(year, month, DOTW.index(dotw), range(WEEK[when], days + 1))


def _get_day(year, month, dotw, days):
    for day in days:
        if date(year, month, day).weekday() == dotw:
            return date(year, month, day)
    raise Exception
