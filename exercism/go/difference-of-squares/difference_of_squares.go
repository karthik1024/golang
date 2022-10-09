package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	return int(math.Pow(float64(n), 2)*math.Pow(float64(n+1), 2)) >> 2
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
