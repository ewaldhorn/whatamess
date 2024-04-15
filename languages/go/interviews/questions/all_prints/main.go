package main

import (
	"fmt"
	"os"
)

func main() {
	// to the console
	println("Also to the", "screen")
	fmt.Printf("A formatted %s\n", "string")

	// to a string
	myOutput := fmt.Sprintln("Prints to a string")

	// output the string
	println(myOutput)

	// print to stdout
	_, err := fmt.Fprint(os.Stdout, "this", " is ", "text", "\n")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
}
