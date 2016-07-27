class Clock:
    def __init__(self, hours, minutes):
        self.hours = hours
        self.minutes = minutes
        self.time()

    def time(self):
        h = (self.hours + (self.minutes / 60)) % 24
        m = self.minutes % 60
        if m < 0:
            m += 60
            h -= 1
        if h < 0:
            h += 24
        return (int(h), int(m))

    def add(self, minutes):
        self.minutes += minutes
        return self

    def __repr__(self):
        (h, m) = self.time()
        return "{h:0>2d}:{m:0>2d}".format(h=h, m=m)

    def __eq__(self, other):
        return repr(self) == repr(other)
