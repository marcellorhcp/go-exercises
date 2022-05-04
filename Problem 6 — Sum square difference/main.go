package main

import (
	"fmt"
)

func main() {
	var i int = 1
	var limit int = 100
	var sumOfSquare int
	var sum int
	for i <= limit {
		sumOfSquare += i * i
		sum += i
		i++
	}
	squareOfTheSum := sum * sum
	difference := squareOfTheSum - sumOfSquare
	fmt.Printf("The difference is %v", difference)
}
