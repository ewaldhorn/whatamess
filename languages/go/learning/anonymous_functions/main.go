package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	// simple anonymous function with no parameters
	go func() {
		fmt.Println("Constant message")
	}()

	// simple anonymous function with parameter
	go func(message string) {
		fmt.Println(message)
	}("hello, dynamic message")
}
