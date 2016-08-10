NORTH = "north"
SOUTH = "south"
EAST = "east"
WEST = "west"


class Robot:
    def __init__(self, bearing=NORTH, x=0, y=0):
        self._directions = [NORTH, EAST, SOUTH, WEST]
        self.coordinates = (x, y)
        self.bearing = bearing

    def turn_right(self):
        i = self._directions.index(self.bearing) + 1
        if i == 4:
            i = 0
        self.bearing = self._directions[i]

    def turn_left(self):
        i = self._directions.index(self.bearing) - 1
        if i == -1:
            i = 3
        self.bearing = self._directions[i]

    def advance(self):
        actions = {NORTH: (self.coordinates[0], self.coordinates[1] + 1),
                   SOUTH: (self.coordinates[0], self.coordinates[1] - 1),
                   EAST: (self.coordinates[0] + 1, self.coordinates[1]),
                   WEST: (self.coordinates[0] - 1, self.coordinates[1])}
        self.coordinates = actions[self.bearing]

    def simulate(self, movements):
        actions = {"R": self.turn_right,
                   "L": self.turn_left,
                   "A": self.advance}
        [actions[movement]() for movement in movements]
