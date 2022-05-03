package main

import "fmt"

//Fibonacci sequence 0,1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144 ... < 4_000_000

func main() {

	var previousTerm int = 0
	var nextTerm int = 1
	var temp int
	var sum int

	for previousTerm < 4000000 {
		temp = previousTerm + nextTerm
		previousTerm = nextTerm
		nextTerm = temp

		if previousTerm%2 == 0 {
			sum += previousTerm
		}
	}
	fmt.Println(sum)
}
