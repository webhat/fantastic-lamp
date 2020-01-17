package main

import (
	"fmt"
	"math"
)

var fibs = []int{0, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

// IsFibCalc checks if the provided integer is a Fibonacci number
func IsFibCalc(n int) bool {
	v1 := (5*(math.Pow(float64(n), 2)) + 4)
	v2 := math.Sqrt(v1)
	fmt.Println(fmt.Sprintf("1: v1=%f, v2=%f", v1, v2))
	if v2 == float64(int(v2)) {
		return true
	}

	v1 = (5*(math.Pow(float64(n), 2)) - 4)
	v2 = math.Sqrt(v1)
	fmt.Println(fmt.Sprintf("1: v1=%f, v2=%f", v1, v2))
	if v2 == float64(int(v2)) {
		return true
	}
	return false
}
