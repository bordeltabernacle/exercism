package hamming

import (
	"errors"
	"strings"
)

const testVersion = 4

func Distance(a, b string) (int, error) {
	aSplit := strings.Split(a, "")
	bSplit := strings.Split(b, "")
	distance := 0
	if len(aSplit) != len(bSplit) {
		return 0, errors.New("math: square root of negative number")
	}
	for c := 0; c < len(aSplit); c++ {
		if aSplit[c] != bSplit[c] {
			distance++
		}
	}
	return distance, nil
}
