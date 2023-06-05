package main

import "fmt"

func helloLiterals() {
	world := func() {
		fmt.Printf("Hello, world!\n")
	}

	world()
}

func moreLiterals() {
	summer := func(a int, b int) int {
		return a + b
	}

	answer := summer(4, 5)

	fmt.Printf("The sum of 4 and 5 is %d\n", answer)
}

func andEvenMoreLiterals(fn func(msg string), msg string) {
	fn(msg)
}

func surround() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "---------------------------------")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "---------------------------------")
	}
}

func main() {
	fmt.Printf("\nFunction Literals:\n\n")

	helloLiterals()
	moreLiterals()
	andEvenMoreLiterals(surround(), "This is it! Literally, it is!")

	fmt.Println()
}
