package main

import "fmt"

func main() {
	b := &Builder{}
	employee := b.
		Name("Michael Scott").
		Role("manager").
		Build()
	fmt.Println(employee)
	// {Michael Scott manager 20000 60000}
}
