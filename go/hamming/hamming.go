package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (dist int, err error) {

	if len(a) != len(b) {
		return dist, errors.New("DNA strands of unequal length")
	}

	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}

	return
}
