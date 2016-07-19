package raindrops

import (
	"strconv"
)

const testVersion = 2

// Convert converts a number to a string,
// the contents of which depends on the number's prime factors
func Convert(n int) (raindrop string) {
	if n%3 == 0 {
		raindrop += "Pling"
	}
	if n%5 == 0 {
		raindrop += "Plang"
	}
	if n%7 == 0 {
		raindrop += "Plong"
	}
	if raindrop == "" {
		raindrop = strconv.Itoa(n)
	}
	return
}
