package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	fmt.Println("Please use 'go test .' to run the tests.")
	testBigWordList()
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
	fmt.Printf("We have %d entries!\n", len(words))
}
