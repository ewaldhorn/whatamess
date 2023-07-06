package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	greetCommand := flag.NewFlagSet("hello", flag.ExitOnError)
	greetNamePointer := greetCommand.String("name", "Go CLI", "name to greet")

	if len(os.Args) < 2 {
		fmt.Println("Nothing to do.")
		listCommands()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "hello":
		greetCommand.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *greetNamePointer)
	default:
		fmt.Printf("Unknown command: '%s'.\n", cmd)
		os.Exit(2)
	}
}

func listCommands() {
	fmt.Println("The commands are:")

	fmt.Printf("%-8s - Gives a greeting.\n", "hello")

	fmt.Println()
}
