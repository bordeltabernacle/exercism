class Allergies:
    def __init__(self, score):
        self.score = score
        self._as = self._scores()
        self.lst = [k for k in self._as.keys() if self._as.get(k)]

    def _scores(self):
        binary_score = list(bin(self.score))[2:][::-1]
        reactions = [int(i) for i in binary_score]
        allergies = """
        eggs peanuts shellfish strawberries tomatoes chocolate pollen cats
        """.split()
        return dict(zip(allergies, reactions))

    def is_allergic_to(self, allergy):
        return bool(self._as.get(allergy))
