package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello   = "hello"
	CmdGoodBye = "bye"
)

func main() {
	fmt.Println("Go Readers Exercise")
	fmt.Println()
	startApplication()
}

func startApplication() {
	scanner := bufio.NewScanner(os.Stdin)
	numLines := 0
	numCommands := 0

	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "Q" {
			break
		} else {
			text := strings.TrimSpace(scanner.Text())

			switch text {
			case CmdHello:
				numCommands += 1
				fmt.Println("command response: Hi")
			case CmdGoodBye:
				numCommands += 1
				fmt.Println("command response: Bye")
			}

			if text != "" {
				numLines += 1
			}
		}
	}

	fmt.Printf("\nYou entered %d lines of text.\n", numLines)
	fmt.Printf("Of those, %d were valid commands.\n\n", numCommands)
}
