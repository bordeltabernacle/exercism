package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(n int) (raindrop string) {

	sounds := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for k, v := range sounds {
		if n%k == 0 {
			raindrop += v
		}
	}

	if raindrop == "" {
		raindrop = strconv.Itoa(n)
	}
	return
}
