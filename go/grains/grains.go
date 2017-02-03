package grains

import "fmt"

const testVersion = 1

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Invalid square: %d", n)
	}
	return 1 << uint(n-1), nil
}

func Total() (t uint64) {
	for i := 0; i < 64; i++ {
		t += 1 << uint(i)
	}
	return
}
