from datetime import date
import calendar


def meetup_day(year, month, dotw, when):
    all_days = [d for d in calendar.day_name]
    (_, days) = calendar.monthrange(year, month)
    start_date = {"1st": 1,"2nd": 8,"3rd": 15,"4th": 22,"5th": 28,"teenth": 13}
    if when == "last":
        return _get_day(year, month, all_days.index(dotw),
                reversed(range(days - 6, days + 1)))
    return _get_day(year, month, all_days.index(dotw),
            range(start_date[when], days + 1))


def _get_day(year, month, dotw, days):
    day = [day for day in days if date(year, month, day).weekday() == dotw][0]
    return date(year, month, day)
