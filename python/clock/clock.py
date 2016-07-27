class Clock:
    def __init__(self, hours, minutes):
        self.minutes = (hours * 60 + minutes) % 1440

    def add(self, minutes):
        self.minutes = (self.minutes + minutes) % 1440
        return self

    def __repr__(self):
        return "{h:0>2d}:{m:0>2d}".format(
                h=int(self.minutes/60), m=int(self.minutes%60))

    def __eq__(self, other):
        return repr(self) == repr(other)
