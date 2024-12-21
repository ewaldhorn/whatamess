package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// struct composition
type Employee struct {
	Person
	Position string
}

// ----------------------------------------------------------------------------
func main() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name) // Output: Alice
	fmt.Println(p.Age)  // Output: 30

	// struct composition usage
	e := Employee{
		Person:   Person{Name: "Bob", Age: 25},
		Position: "Developer",
	}

	fmt.Println(e.Name)     // Output: Bob
	fmt.Println(e.Position) // Output: Developer
}
