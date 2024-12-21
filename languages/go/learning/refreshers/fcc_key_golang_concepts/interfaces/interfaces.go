package main

import "fmt"

// ----------------------------------------------------------------------------
type Speaker interface {
	Speak() string
}

// ----------------------------------------------------------------------------
type Person struct {
	Name string
}

// ----------------------------------------------------------------------------
func (p Person) Speak() string {
	return "Hi, my name is " + p.Name
}

// ----------------------------------------------------------------------------
func saySomething(s Speaker) {
	fmt.Println(s.Speak())
}

// ----------------------------------------------------------------------------
func main() {
	p := Person{Name: "Alice"}
	saySomething(p) // Output: Hi, my name is Alice
}
