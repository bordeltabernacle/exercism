class Allergies:
    def __init__(self, score):
        self.score = score
        self.allergens = [
            "eggs",
            "peanuts",
            "shellfish",
            "strawberries",
            "tomatoes",
            "chocolate",
            "pollen",
            "cats",
        ]
        _binary_score = list(reversed(bin(score)))[:-2]
        _reactions = [int(i) for i in _binary_score]
        _chart = dict(zip(self.allergens, _reactions))
        self.lst = [k for k in _chart.keys() if _chart.get(k)]

    def is_allergic_to(self, allergy):
        return allergy in self.lst
