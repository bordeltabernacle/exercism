package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) []Triplet {
	var ts []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if (a*a)+(b*b) == (c * c) {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return ts
}

func Sum(p int) []Triplet {
	var ts []Triplet
	triplets := Range(0, p)
	for i := 0; i < len(triplets); i++ {
		a, b, c := triplets[i][0], triplets[i][1], triplets[i][2]
		if a+b+c == p {
			ts = append(ts, Triplet{a, b, c})
		}
	}
	return ts
}
