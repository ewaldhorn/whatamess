package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sayWelcome()

	// create channel to indicate when the user wants to quit
	allDoneChannel := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(allDoneChannel)
	// block until the doneChannel gets a value
	<-allDoneChannel

	// close channel
	close(allDoneChannel)

	sayGoodbye()
}

func sayWelcome() { fmt.Println("\nHello! Pick a number.\nOr (Q)uit : "); prompt() }
func sayGoodbye() { fmt.Println("\nThat was fun, let's do it again soon!") }

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(signal chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			signal <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	number, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole, positive number", false
	}

	_, msg := isPrime(number)

	return msg, false
}

func isPrime(number int) (bool, string) {
	if number < 2 {
		return false, fmt.Sprintf("%d is not a prime number.", number)
	}

	// use mod to figure out if it's a prime
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number, divisible by %d.", number, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", number)
}
