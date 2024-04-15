package main

import "fmt"

func main() {
	fmt.Println("\nLooking at defer in Go:")
	fmt.Println("")

	defer fmt.Println("")
	defer fmt.Println("Want to go up the hill?")
	defer hello("Jill")
	hello("Jack")
}

func hello(name string) {
	fmt.Println("Hello", name)
}
