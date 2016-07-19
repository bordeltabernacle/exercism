from datetime import date
from calendar import monthrange


DOTW = {
        "monday": 0,
        "tuesday": 1,
        "wednesday": 2,
        "thursday": 3,
        "friday": 4,
        "saturday": 5,
        "sunday": 6
}

WEEK = {
        "1st": 1,
        "2nd": 8,
        "3rd": 15,
        "4th": 22,
        "teenth": 13
}


def meetup_day(year, month, dotw, when):
    (_, no_of_days) = monthrange(year, month)
    if when == "last":
        while no_of_days > 1:
            if date(year, month, no_of_days).weekday() == DOTW[dotw.lower()]:
                return date(year, month, no_of_days)
            no_of_days -= 1
    else:
        i = WEEK[when]
        while i < no_of_days:
            if date(year, month, i).weekday() == DOTW[dotw.lower()]:
                return date(year, month, i)
            i += 1

