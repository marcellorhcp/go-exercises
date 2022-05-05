package main

import (
	"fmt"
	"math"
)

func verifyPrimeNumber(number int) bool {
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
	var limit int = 2_000_000
	var sumOfThePrimes int
	var i int = 2
	for i <= limit {
		if verifyPrimeNumber(i) {
			sumOfThePrimes += i
		}
		i++
	}
	fmt.Println(sumOfThePrimes)
}
