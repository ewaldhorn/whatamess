package main

import "fmt"

// ----------------------------------------------------------------------------
func modifyValue(ptr *int) {
	*ptr = 100
}

// ----------------------------------------------------------------------------
func main() {
	var x int = 42
	var ptr *int // Declare a pointer to an integer

	ptr = &x // Assign the address of x to the pointer

	fmt.Println()
	fmt.Println("Value of x  :", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value stored at the pointer:", *ptr) // Dereference the pointer to get the value

	fmt.Println("...Changing the value...")
	modifyValue(ptr)
	fmt.Println("Value of x  :", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value stored at the pointer:", *ptr) // Dereference the pointer to get the value

	fmt.Println()
}
