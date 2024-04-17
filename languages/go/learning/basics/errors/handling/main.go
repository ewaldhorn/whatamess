package main

import (
	"errors"
	"fmt"
)

// -------------------------------------------------------------------------------------------------
func noProblem() (int, error) {
	return 123, nil
}

// -------------------------------------------------------------------------------------------------
func gotAProblem() (int, error) {
	return -1, errors.New("this is dumb")
}

// -------------------------------------------------------------------------------------------------
func noProblemHere() (int, error) {
	return 321, nil
}

// -------------------------------------------------------------------------------------------------
func handleProblem() {
	println("Sorted!")
}

// -------------------------------------------------------------------------------------------------
func main() {
	result, err := noProblem()

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	println("Read:", result)

	if result, err = gotAProblem(); err != nil {
		fmt.Println("error: ", err)
		return
	}

	println("Read:", result)
}
