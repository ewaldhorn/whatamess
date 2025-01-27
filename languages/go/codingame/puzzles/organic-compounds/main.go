package main

import (
	"bufio"
	"fmt"
	"os"
)

// ----------------------------------------------------------------------------
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		COMPOUND := scanner.Text()
		fmt.Println(COMPOUND)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println("answer") // Write answer to stdout
}

// ----------------------------------------------------------------------------
func validateCompound(entries ...string) string {
	for i := range entries {
		fmt.Println(entries[i])
	}

	return "VALID"
}
