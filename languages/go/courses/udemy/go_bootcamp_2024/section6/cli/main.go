package main

import (
	"fmt"
	"os"
	"strconv"
)

// ------------------------------------------------------------------------------------------------
func main() {
	listAllArgs()
	checkArgCount()
	left, right := getArgNumbers()
	fmt.Printf("The sum of %d and %d is %d.\n", left, right, left+right)
}

// ------------------------------------------------------------------------------------------------
func listAllArgs() {
	fmt.Println("os.Args:", os.Args)
}

// ------------------------------------------------------------------------------------------------
func checkArgCount() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments, expected 2 integers.")
		os.Exit(1)
	}
}

// ------------------------------------------------------------------------------------------------
func getArgNumbers() (int, int) {
	left, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("could not parse '%s' as an integer", os.Args[1])
		os.Exit(2)
	}

	right, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("could not parse '%s' as an integer", os.Args[2])
		os.Exit(2)
	}

	return left, right
}
