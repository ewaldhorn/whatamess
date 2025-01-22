package main

import "math"

// ----------------------------------------------------------------------------
// Basic implementation
func getSecondLargestBasic(numbers ...int) int {
	// empty
	if len(numbers) == 0 {
		return math.MinInt
	}

	// just one
	if len(numbers) == 1 {
		return numbers[0]
	}

	// give first and second starting values
	first, second := numbers[0], numbers[0]

	for _, num := range numbers[1:] {
		if num > first {
			// if number is greater than first, we have a second number
			second = first
			first = num
		} else {
			// it's not greater than first, but it is greater than second
			if num > second && num != first {
				second = num
			}
		}
	}

	return second
}
