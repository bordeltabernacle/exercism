package raindrops

import (
	"strconv"
)

const testVersion = 2

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
