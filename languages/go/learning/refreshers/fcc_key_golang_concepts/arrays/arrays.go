package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	numbers := [3]int{1, 2, 3}
	var ptr *[3]int // Pointer to an array of integers

	ptr = &numbers // Assign the address of the array to the pointer

	fmt.Println("Original array (1):", numbers)
	fmt.Println("Modified array through pointer:", modifyArray(ptr))
	fmt.Println("Original array (2):", numbers)
}

// ----------------------------------------------------------------------------
func modifyArray(ptr *[3]int) [3]int {
	// Modify the array through the pointer
	ptr[0], ptr[1], ptr[2] = 10, 20, 30
	return *ptr
}
