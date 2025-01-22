package main

import "math"

func getSecondLargest(numbers ...int) int {
	// empty
	if len(numbers) == 0 {
		return math.MinInt
	}

	// just one
	if len(numbers) == 1 {
		return numbers[0]
	}

	return numbers[0]
}
