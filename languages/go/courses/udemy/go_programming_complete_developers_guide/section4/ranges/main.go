package main

import "fmt"

func main() {
	myArray := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for idx, value := range myArray {
		fmt.Printf("We have %2d at index %d\n", value, idx)
	}
	fmt.Println()

	myString := "This is a big fat string!"
	for i, character := range myString {
		fmt.Printf("We have %q at position %2d\n", character, i)
	}
	fmt.Println()
}
