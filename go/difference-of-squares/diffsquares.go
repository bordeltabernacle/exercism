package diffsquares

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func SquareOfSums(n int) int {
	sum := (n*n + n) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}
