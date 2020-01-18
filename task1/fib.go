package task1

import (
	"math"
)

// IsFibCalc checks if the provided integer is a Fibonacci number
func IsFibCalc(n int) bool {
	v1 := (5*(math.Pow(float64(n), 2)) + 4)
	v2 := math.Sqrt(v1)
	if v2 == float64(int(v2)) {
		return true
	}

	v1 = (5*(math.Pow(float64(n), 2)) - 4)
	v2 = math.Sqrt(v1)
	if v2 == float64(int(v2)) {
		return true
	}
	return false
}
