package diffsquares

func SquareOfSum(n int) int {
	s := (n * (n + 1) >> 1)
	return s * s
}

func SumOfSquares(n int) int {
	r := 0
	for i := 1; i <= n; i++ {
		s := i * i
		r += s
	}
	return r
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
