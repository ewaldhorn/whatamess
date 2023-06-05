package main

import (
	"fmt"
	"readwrite/readers"
	"readwrite/writers"
)

func main() {
	fmt.Println()
	fmt.Println("Go Readers and Writers")

	fmt.Println()
	readers.SimpleReader()

	fmt.Println()
	readers.SimpleReaderUsingBUFIO()

	// fmt.Println()
	// readers.TestScanner()

	fmt.Println()
	writers.SimpleWriter()

	fmt.Println()
}
