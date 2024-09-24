package main

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------------
func main() {
	justBasicSlice()
	sliceOfSlice()
}

// ----------------------------------------------------------------------------
func printHeader(title string) {
	fmt.Println("")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

// ----------------------------------------------------------------------------
func justBasicSlice() {
	printHeader("Basic Slices:")

	names := []string{"Bob", "Sal", "Rose", "Dennis"}

	fmt.Printf("There are %d names in the list.\n", len(names))

	fmt.Println("Adding \"Taylor\" to the list.")
	names = append(names, "Taylor")

	fmt.Printf("There are now %d names in the list.\n", len(names))

	fmt.Println("Adding \"Ron\", \"Dopey\" and \"Angela\" to the list.")
	names = append(names, "Ron", "Dopey", "Angela")

	fmt.Printf("There are now %d names in the list.\n", len(names))
}

// ----------------------------------------------------------------------------
func sliceOfSlice() {
	printHeader("Slices of Slices:")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("There are %d elements in numbers.\n", len(numbers))
}
