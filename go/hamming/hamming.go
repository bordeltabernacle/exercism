// Calculates the Hamming difference between two DNA strands.
package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (dist int, err error) {
	// check DNA strands are the same length
	if len(a) != len(b) {
		return dist, errors.New("DNA strands of unequal length")
	}
	// compare each nucleotide in strands
	// incrementing dist if not equal
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return
}
