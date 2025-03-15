package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	fmt.Printf("The sum of the numbers 1..5 is: %d\n", sum_many(1, 2, 3, 4, 5))
}

// ----------------------------------------------------------------------------
func sum_many(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
