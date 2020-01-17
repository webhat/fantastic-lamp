package task1

import (
	//	"github.com/devaof/fizzbuzz"
	"strconv"
)

func fizzorbuzz(n, div int) bool {
	return n%div == 0
}

// FizzBuzzPinkFlamingo returns a string FizzBuzz, Fizz or Buzz
func FizzBuzzPinkFlamingo(n int) string {
	fizz := fizzorbuzz(n, 3)
	buzz := fizzorbuzz(n, 5)
	fib := IsFibCalc(n)

	if fizz && buzz {
		if fib {
			return "Pink Flamingo"
		}
		return "FizzBuzz"
	}
	if fizz {
		return "Fizz"
	}
	if buzz {
		return "Buzz"
	}
	if fib {
		return "Flamingo"
	}
	return strconv.Itoa(n)
}
