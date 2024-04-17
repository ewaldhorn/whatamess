package main

import "fmt"

// -------------------------------------------------------------------------------------------------
func sumOfInts(input *[]int) int {
	total := 0

	for _, val := range *input {
		total += val
	}

	return total
}

// -------------------------------------------------------------------------------------------------
func sumOfFloats(input *[]float32) float32 {
	var total float32

	for _, val := range *input {
		total += val
	}

	return total
}

// -------------------------------------------------------------------------------------------------
func sumOf[T int | float32](input *[]T) T {
	var total T

	for _, val := range *input {
		total += val
	}

	return total
}

// -------------------------------------------------------------------------------------------------
func main() {
	myNumbers := []int{1, 2, 3, 4}
	myFloats := []float32{1.01, 2.02, 3.03, 4.04}

	fmt.Printf("Summing with sumOfInts gives: %d\n", sumOfInts(&myNumbers))
	fmt.Printf("Summing with sumOf gives: %d\n", sumOf(&myNumbers))

	fmt.Printf("Summing with sumOfFloats gives: %.2f\n", sumOfFloats(&myFloats))
	fmt.Printf("Summing with sumOf gives: %.2f\n", sumOf(&myFloats))
}
