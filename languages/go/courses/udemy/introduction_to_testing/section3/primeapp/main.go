package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	printIntro()

	// create a channel to indicate user wants to exit
	userMonitorChannel := make(chan bool)

	// start goroutine to get / check numbers
	go readUserInput(userMonitorChannel)

	// wait until user indicates they are done
	<-userMonitorChannel

	// close channel
	close(userMonitorChannel)

	// print a goodbye message
	printOuttro()
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read the user input
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// now try to convert the string to a number
	numToCheck, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return "Please enter a whole number!", false
	}

	_, msg := isPrime(numToCheck)

	return msg, false
}

func prompt() {}

func printIntro() {
	fmt.Println("Hello!")
	fmt.Println("Enter whole numbers to check if they are prime (q to quit).")
	prompt()
}

func printOuttro() {
	fmt.Println("Ok, all done. Goodbye!")
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}

	// negative numbers are not prime
	if n < 0 {
		return false, "Negative numbers are not prime, by definition!"
	}

	// use the modulus operator repeatedly to see if we have a prime number
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			// not a prime number
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", n)
}
