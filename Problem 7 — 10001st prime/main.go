package main

import (
	"fmt"
	"math"
)

func VerifyPrimeNumber(number int) bool {
	if number < 2 {
		return false
	}
	is := true
	max := math.Sqrt(float64(number))
	for i := 2; i <= int(max); i++ {
		if number%i == 0 {
			is = false
			break
		}
	}
	return is
}

func main() {
	var primePosition int = 10_001
	var counter int = 0
	var i int = 2
	for counter < primePosition {
		if VerifyPrimeNumber(i) {
			counter++
		}
		i++
	}
	fmt.Println("the 10001st prime number is", i-1)
}
