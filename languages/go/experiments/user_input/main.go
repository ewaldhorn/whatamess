package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	withScanLine()
	withBufIOReader()
	withBufIOScanner()
}

func withScanLine() {
	// great for reading a line at a time
	fmt.Print("scanline: input text (3 items): ")
	var w1, w2, w3 string
	n, err := fmt.Scanln(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nnumber of items read: %d\n", n)
	fmt.Printf("read line: [%s] [%s] [%s]\n", w1, w2, w3)
}

func withBufIOReader() {
	fmt.Print("bufio.Reader: input text: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nread line: [%s]\n", line)
}

func withBufIOScanner() {
	fmt.Print("scanner: input text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read line: [%s]\n", scanner.Text())
}
