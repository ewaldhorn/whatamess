package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	// N: Number of elements which make up the association table.
	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)

	// Q: Number Q of file names to be analyzed.
	var Q int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &Q)

	extMap := map[string]string{}

	for i := 0; i < N; i++ {
		// EXT: file extension
		// MT: MIME type.
		var EXT, MT string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &EXT, &MT)
		extMap[strings.ToLower(EXT)] = MT
	}

	for i := 0; i < Q; i++ {
		scanner.Scan()
		FNAME := scanner.Text()
		extensionPos := strings.LastIndex(FNAME, ".")
		extension := ""

		if extensionPos == -1 {
			fmt.Println("UNKNOWN")
		} else {
			extension = strings.ToLower(FNAME[extensionPos+1:])
			name, ok := extMap[extension]

			if ok {
				fmt.Println(name)
			} else {
				fmt.Println("UNKNOWN")
			}
		}
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
}
