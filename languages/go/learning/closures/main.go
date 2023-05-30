package main

import "fmt"

func main() {
	addFunction := func(a, b int) int {
		return a + b
	}

	num1 := 5
	num2 := 2

	fmt.Printf("Adding %d and %d via Add gives %d\n", num1, num2, Add(num1, num2))
	fmt.Printf("Adding %d and %d via addFunction gives %d\n", num1, num2, addFunction(num1, num2))
}

func Add(a, b int) int {
	return a + b
}
