package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("Nothing to do.")
		listCommands()
		os.Exit(1)
	}

	cmd := arguments[0]

	switch cmd {
	case "hello":
		fmt.Println("Hello, Go CLI!")
	default:
		fmt.Printf("Unknown command: '%s'.\n", cmd)
	}
}

func listCommands() {
	fmt.Println("The commands are:")

	fmt.Printf("%-8s - Gives a greeting.\n", "hello")

	fmt.Println()
}
