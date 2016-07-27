class Clock:
    def __init__(self, hours, minutes):
        self.hours = hours
        self.minutes = minutes
        self.time()

    def time(self):
        t = (self.hours * 60 + self.minutes) % 1440
        if t < 0:
            t += 1440
        self.hours = int(t / 60)
        self.minutes = int(t % 60)
        return self

    def add(self, minutes):
        self.minutes += minutes
        return self.time()

    def __repr__(self):
        return "{h:0>2d}:{m:0>2d}".format(h=self.hours, m=self.minutes)

    def __eq__(self, other):
        return repr(self) == repr(other)
