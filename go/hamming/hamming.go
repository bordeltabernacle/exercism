package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	var dist int
	if len(a) != len(b) {
		return 0, errors.New("DNA strands of unequal length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
