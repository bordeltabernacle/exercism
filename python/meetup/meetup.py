from datetime import date
from calendar import monthrange


DOTW = ["monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday",
        "saturday",
        "sunday"]
WEEK = {
        "1st": 1,
        "2nd": 8,
        "3rd": 15,
        "4th": 22,
        "5th": 28,
        "teenth": 13
}


def meetup_day(year, month, dotw, when):
    (_, no_of_days) = monthrange(year, month)
    if when == "last":
        while no_of_days > 1:
            if date(year, month, no_of_days).weekday() == DOTW.index(dotw.lower()):
                return date(year, month, no_of_days)
            no_of_days -= 1
    else:
        i = WEEK[when]
        while i < no_of_days:
            if date(year, month, i).weekday() == DOTW.index(dotw.lower()):
                return date(year, month, i)
            i += 1

