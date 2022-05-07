package main

import (
	"fmt"
)

func CheckLetters(letters string) []string {
	var missingLetters = []string{}
	var limit int
	index := int(letters[0])

	for n := 1; n < len(letters); n++ {
		if int(letters[n-1]) > int(letters[n]) {
			return missingLetters
		}
	}

	if int(letters[0]) > 64 && int(letters[0]) < 91 {
		limit = 91
	}
	if int(letters[0]) > 96 && int(letters[0]) < 123 {
		limit = 123
	}

	for i := 0; i < len(letters); i++ {
		for j := index; j < limit; j++ {
			if int(letters[i]) == j {
				index += 1
				break
			} else {
				missingLetters = append(missingLetters, string(j))
				index += 1
			}
		}
	}
	return missingLetters
}

func main() {
	fmt.Println(CheckLetters("abcdf"))
	fmt.Println(CheckLetters("acdf"))
	fmt.Println(CheckLetters("OQRS"))
	fmt.Println(CheckLetters("acdb"))
	fmt.Println(CheckLetters("abcz"))
}
