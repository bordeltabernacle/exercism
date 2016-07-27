class Clock:
    def __init__(self, hours, minutes):
        self.minutes = (hours * 60 + minutes) % 1440

    def add(self, minutes):
        self.minutes = (self.minutes + minutes) % 1440
        return self

    def __repr__(self):
        return "{:02}:{:02}".format(self.minutes // 60, self.minutes % 60)

    def __eq__(self, other):
        return repr(self) == repr(other)
