package main

import (
	"fmt"
)

func main() {
	var number int = 1
	var i, counter int
	var limit int = 20
	for counter < limit {
		counter = 0
		for i = 1; i <= limit; i++ {
			if number%i == 0 {
				counter += 1
			}
		}
		number++
	}
	smallestPositiveNumber := number - 1
	fmt.Printf("%v is the smallest number that can be divided by each of the numbers from 1 to %v without any remainder", smallestPositiveNumber, limit)
}
