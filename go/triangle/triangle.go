package triangle

import "math"

const testVersion = 2

// Kind implements the different kinds of triangle
type Kind string

// Identifiers for the different kinds of triangle
const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides determines which kind of triangle
// is represented by the given three sides; a, b & c
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

// triTest tests if the given sides form a valid triangle
func triTest(a, b, c float64) bool {
	return validSide(a) && validSide(b) && validSide(c) && meetsTriIneq(a, b, c)
}

// validSide tests a side is not infinity,
// not NaN & greater than zero
func validSide(s float64) bool {
	return !(math.IsInf(s, 0) || math.IsNaN(s) || s <= 0)
}

// meetsTriIneq checks all sides together meet
// the Triangle Inequality Theorem
func meetsTriIneq(a, b, c float64) bool {
	return !(a+b < c || b+c < a || c+a < b)
}
