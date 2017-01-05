package grains

import (
	"fmt"
	"math"
)

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Invalid square: %d", n)
	}
	return uint64(math.Pow(2, float64(n-1))), nil
}

func Total() (t uint64) {
	for i := 1; i < 65; i++ {
		val, _ := Square(i)
		t += val
	}
	return
}
