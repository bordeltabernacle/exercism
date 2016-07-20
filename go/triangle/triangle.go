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
	return isNotInf(a, b, c) && isNotNaN(a, b, c) && greaterThanZero(a, b, c) && meetsTriIneq(a, b, c)
}

// isNotInf checks all sides are not infinity
func isNotInf(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	return true
}

// isNotNaN checks all sides are not NaN
func isNotNaN(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	return true
}

// greaterThanZero checks all sides are greater than zero
func greaterThanZero(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return true
}

// meetsTriIneq checks all sides together meet
// the Triangle Inequality Theorem
func meetsTriIneq(a, b, c float64) bool {
	if a+b < c || b+c < a || c+a < b {
		return false
	}
	return true
}
