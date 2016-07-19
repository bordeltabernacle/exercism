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
    (_, no_of_days) = calendar.monthrange(year, month)
    if when == "last":
        for day in reversed(range(no_of_days + 1)):
            if date(year, month, day).weekday() == DOTW.index(dotw):
                return date(year, month, day)
        raise Exception
    else:
        for day in range(WEEK[when], no_of_days + 1):
            if date(year, month, day).weekday() == DOTW.index(dotw):
                return date(year, month, day)
        raise Exception

#  def scan(year, month, days, dotw):
