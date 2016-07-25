// Package diffsquares provides ...
package diffsquares

import "math"

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}
