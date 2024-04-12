package main

import "fmt"

func check(input int) string {
	isThrees := input%3 == 0
	isFives := input%5 == 0

	if isThrees && isFives {
		return "Fizz Buzz"
	}

	if isThrees {
		return "Fizz"
	}

	if isFives {
		return "Buzz"
	}

	return fmt.Sprintf("%d", input)
}
