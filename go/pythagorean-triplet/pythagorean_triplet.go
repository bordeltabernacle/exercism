package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (ts []Triplet) {
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

func Sum(p int) (ts []Triplet) {
	triplets := Range(0, p)
	for _, t := range triplets {
		if t[0]+t[1]+t[2] == p {
			ts = append(ts, t)
		}
	}
	return ts
}
