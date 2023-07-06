package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	HelloCmd = "hello"
)

func main() {
	greetCmd := flag.NewFlagSet(HelloCmd, flag.ExitOnError)
	greetFormalCmd := flag.NewFlagSet("formal", flag.ExitOnError)
	greetNamePtr := greetFormalCmd.String("name", "Go CLI", "name to greet")

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	switch os.Args[1] {
	case HelloCmd:
		greetCmd.Parse(os.Args[2:])
		fmt.Println("Hi there!")
		if greetCmd.Arg(0) == "formal" {
			greetFormalCmd.Parse(os.Args[3:])
			fmt.Printf("How do you do, %s?\n", *greetNamePtr)
		}
	default:
		fmt.Println("Unknown command")
		os.Exit(2)
	}
}
