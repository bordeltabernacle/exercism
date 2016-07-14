package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(n int) string {
	var raindrop string
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
		return strconv.Itoa(n)
	}
	return raindrop
}
