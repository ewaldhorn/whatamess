package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("\n\nSome simple IO exercises:")
	fmt.Print("\n")

	simple_io()

	fmt.Print("\n\n")
}

func simple_io() {
	// fmt.Println() becomes the below
	fmt.Fprintln(os.Stdout, []any{"\tThis is one way to use Print!"}...)
}
