package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	sayWelcome()

	// create channel to indicate when the user wants to quit
	allDoneChannel := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, allDoneChannel)
	// block until the doneChannel gets a value
	<-allDoneChannel

	// close channel
	close(allDoneChannel)

	sayGoodbye()
}

// ----------------------------------------------------------------- sayWelcome
func sayWelcome() {
	fmt.Println("\nHello! Pick a number.\nOr (Q)uit : ")
	prompt()
}

// ----------------------------------------------------------------- sayGoodbye
func sayGoodbye() {
	fmt.Println("\nThat was fun, let's do it again soon!")
}

// --------------------------------------------------------------------- prompt
func prompt() {
	fmt.Print("-> ")
}

// -------------------------------------------------------------- readUserInput
func readUserInput(in io.Reader, signal chan bool) {
	scanner := bufio.NewScanner(in)

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

// --------------------------------------------------------------- checkNumbers
func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	number, err := strconv.Atoi(scanner.Text())

	if err != nil || number < 0 {
		return "Please enter a whole, positive number", false
	}

	_, msg := isPrime(number)

	return msg, false
}

// -------------------------------------------------------------------- isPrime
func isPrime(number int) (bool, string) {
	if number < 2 {
		return false, fmt.Sprintf("%d is not a prime number.", number)
	}

	// use mod to figure out if it's a prime
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false,
				fmt.Sprintf("%d is not a prime number, divisible by %d.", number, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", number)
}
