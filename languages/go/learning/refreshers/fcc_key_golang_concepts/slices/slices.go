package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	// Using a pointer to a slice
	numbers := []int{1, 2, 3}
	var slicePtr *[]int = &numbers

	// Using a pointer to a map
	person := map[string]string{"name": "Alice", "age": "25"}
	var mapPtr *map[string]string = &person

	fmt.Println("Original slice:", numbers)
	fmt.Println("Modified slice through pointer:", modifySlice(slicePtr))
	fmt.Println("Original map:", person)
	fmt.Println("Modified map through pointer:", modifyMap(mapPtr))
}

// ----------------------------------------------------------------------------
func modifySlice(ptr *[]int) []int {
	// Modify the slice through the pointer
	(*ptr)[0], (*ptr)[1], (*ptr)[2] = 10, 20, 30
	return *ptr
}

// ----------------------------------------------------------------------------
func modifyMap(ptr *map[string]string) map[string]string {
	// Modify the map through the pointer
	(*ptr)["name"] = "Bob"
	(*ptr)["age"] = "30"
	return *ptr
}
