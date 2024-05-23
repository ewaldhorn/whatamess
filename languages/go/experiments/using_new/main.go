package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// Using new() to allocate memory for a Person struct
	p1 := new(Person)

	// Using short-hand to do the same
	p2 := &Person{Name: "Bob Doe", Age: 30, Gender: "Male"}

	// Initializing the fields
	p1.Name = "John Doe"
	p1.Age = 30
	p1.Gender = "Male"

	fmt.Println("p1:", p1)
	fmt.Println("Person 1 name:", p1.Name)

	fmt.Println("p2:", p2)
	fmt.Println("Person 2 name:", p2.Name)

}
