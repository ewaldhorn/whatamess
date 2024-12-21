package main

import "fmt"

// ----------------------------------------------------------------------------
type Person struct {
	Name string
	Age  int
}

// ----------------------------------------------------------------------------
func main() {
	p := Person{Name: "John", Age: 30}
	fmt.Println("Before modification:", p)

	modifyPerson(&p)

	fmt.Println("After modification :", p)
}

// ----------------------------------------------------------------------------
func modifyPerson(ptr *Person) {
	ptr.Age = 31
	ptr.Name = "John Doe"
}
