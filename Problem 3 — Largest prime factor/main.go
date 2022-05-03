package main

import (
	"fmt"
)

func VerifyPrimeNumber(number int) bool {
	var primeCounter int = 0
	for j := 1; j <= number; j++ {
		if number%j == 0 {
			primeCounter++
		}
	}
	if primeCounter == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var factorizationOf int = 600_851_475_143
	var i int = 2
	for factorizationOf > 1 {
		if VerifyPrimeNumber(i) && factorizationOf%i == 0 {
			factorizationOf /= i
		}
		i++
	}
	largestPrimeFactor := (i - 1)
	fmt.Println(largestPrimeFactor)
}
