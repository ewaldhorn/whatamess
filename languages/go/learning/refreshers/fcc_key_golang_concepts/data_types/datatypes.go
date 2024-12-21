package main

import "fmt"

func main() {

	// Data Types
	var isEmployed bool = true         // boolean
	var count int = 42                 // integer
	var greeting string = "Hello, Go!" // string

	// Struct
	type Rectangle struct {
		width  float64
		height float64
	}
	rect := Rectangle{width: 10.5, height: 5.2}

	fmt.Println("Is Employed:", isEmployed)
	fmt.Println("Count:", count)
	fmt.Println("Greeting:", greeting)
	fmt.Printf("Rectangle: width = %.2f, height = %.2f\n", rect.width, rect.height)
}
