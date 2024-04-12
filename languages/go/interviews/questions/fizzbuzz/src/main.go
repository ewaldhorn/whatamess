package main

import "fmt"

func check(input int) string {
	isThrees := input%3 == 0
	isFives := input%5 == 0

	if isThrees && isFives {
		return "Fizz Buzz"
	} else if isThrees {
		return "Fizz"
	} else if isFives {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", input)
	}
}

// use a switch instead of if/else
func checkWithSwitch(input int) string {
	isThrees := input%3 == 0
	isFives := input%5 == 0

	switch {
	case isThrees && isFives:
		return "Fizz Buzz"
	case isThrees:
		return "Fizz"
	case isFives:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", input)
	}
}
