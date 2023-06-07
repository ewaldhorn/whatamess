package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	intro()

	// create a channel to allow us to keep going until the user wants to exit
	doneChan := make(chan bool)

	go readUserInput(os.Stdin, doneChan)

	// block until this channel gets a value
	<-doneChan

	// now close the channel, the user wants to exit
	close(doneChan)
	outtro()
}

func intro() {
	fmt.Println("It begins...\n")
}
func outtro() {
	fmt.Println("All done!\n")
}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
	}
}
