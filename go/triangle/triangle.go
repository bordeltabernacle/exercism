package triangle

import "math"

const testVersion = 2

type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
	if triTest(a, b, c) {
		if a == b && b == c {
			return Equ
		}
		if a == b || b == c || a == c {
			return Iso
		}
		return Sca
	}
	return NaT
}

func triTest(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b < c || b+c < a || c+a < b {
		return false
	}
	return true
}
