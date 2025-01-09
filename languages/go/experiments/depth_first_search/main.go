package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
func main() {
	fmt.Println("Please use 'go test .' to run the tests.")
	testBigWordList()
	testBigWordListNoAlloc()
}

// ----------------------------------------------------------------------------
func testBigWordList() {
	words := make([]string, 0, 10_000)
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := range 100 {
		for _, ch := range characters {
			words = append(words, fmt.Sprintf("%c%d", ch, i))
		}
	}

	fmt.Printf("\nList: We have %d entries!\n", len(words))
	for i := range 100 {
		fmt.Printf("%s,", words[i])
	}
	fmt.Println()
}

// ----------------------------------------------------------------------------
func testBigWordListNoAlloc() {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	sizeNeeded := len(characters) * 100
	words := make([]string, sizeNeeded)
	pos := 0

	var builder strings.Builder

	for i := range 100 {
		for _, ch := range characters {
			builder.WriteString(string(ch))
			builder.WriteString(strconv.Itoa(i))
			words[pos] = builder.String()

			builder.Reset()
			pos += 1
		}
	}

	fmt.Printf("\nArray: We have %d entries!\n", len(words))
	for i := range 100 {
		fmt.Printf("%s,", words[i])
	}
	fmt.Println()

}
