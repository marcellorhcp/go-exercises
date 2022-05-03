package main

import (
	"fmt"
)

//https://www.tutorialspoint.com/write-a-golang-program-to-reverse-a-number
func ReverseNumber(num int) int {
	res := 0
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	return res
}

func main() {
	var result int
	var LargestPalindrome int = 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			result = i * j
			if result == ReverseNumber(result) && result > LargestPalindrome {
				LargestPalindrome = result
			}
		}
	}

	fmt.Println(LargestPalindrome)

}
