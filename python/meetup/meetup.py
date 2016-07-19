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
        while no_of_days > 1:
            if date(year, month, no_of_days).weekday() == DOTW.index(dotw.lower()):
                return date(year, month, no_of_days)
            no_of_days -= 1
        raise Exception
    else:
        i = WEEK[when]
        while i < no_of_days:
            if date(year, month, i).weekday() == DOTW.index(dotw.lower()):
                return date(year, month, i)
            i += 1
        raise Exception

