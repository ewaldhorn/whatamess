package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	add := func(a, b int) int {
		return a + b
	}

	subtract := func(a, b int) int {
		return a - b
	}

	var operation func(int, int) int

	operation = add
	fmt.Println("Result of addition:", performOperation(operation, 5, 3))

	operation = subtract
	fmt.Println("Result of subtraction:", performOperation(operation, 5, 3))
}

// ----------------------------------------------------------------------------
func performOperation(op func(int, int) int, a, b int) int {
	return op(a, b)
}
