NORTH, EAST, SOUTH, WEST = range(4)


class Robot:
    def __init__(self, bearing=NORTH, x=0, y=0):
        self.coordinates = (x, y)
        self.bearing = bearing

    def turn_right(self):
        self.bearing = (self.bearing + 1) % 4

    def turn_left(self):
        self.bearing = (self.bearing - 1) % 4

    def advance(self):
        self.coordinates = {
            NORTH: (self.coordinates[0], self.coordinates[1] + 1),
            SOUTH: (self.coordinates[0], self.coordinates[1] - 1),
            EAST: (self.coordinates[0] + 1, self.coordinates[1]),
            WEST: (self.coordinates[0] - 1, self.coordinates[1])
        }[self.bearing]

    def simulate(self, movements):
        [{"R": self.turn_right,
          "L": self.turn_left,
          "A": self.advance}[movement]() for movement in movements]
