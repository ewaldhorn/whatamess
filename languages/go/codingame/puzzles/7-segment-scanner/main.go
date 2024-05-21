package main

import (
	"bufio"
	"fmt"
	"os"
)

type Digit struct {
	line1 string
	line2 string
	line3 string
}

var zero = Digit{line1: " _ ", line2: "| |", line3: "|_|"}
var one = Digit{line1: "   ", line2: "  |", line3: "  |"}
var two = Digit{line1: " _ ", line2: " _|", line3: "|_ "}
var three = Digit{line1: " _ ", line2: " _|", line3: " _|"}
var four = Digit{line1: "   ", line2: "|_|", line3: "  |"}
var five = Digit{line1: " _ ", line2: "|_ ", line3: " _|"}
var six = Digit{line1: " _ ", line2: "|_ ", line3: "|_|"}
var seven = Digit{line1: " _ ", line2: "  |", line3: "  |"}
var eight = Digit{line1: " _ ", line2: "|_|", line3: "|_|"}
var nine = Digit{line1: " _ ", line2: "|_|", line3: " _|"}

func isMatched(left Digit, right Digit) bool {
	return left.line1 == right.line1 && left.line2 == right.line2 && left.line3 == right.line3
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()
	scanner.Scan()
	line3 := scanner.Text()

	for i := 0; i < len(line1); i += 3 {
		digit := Digit{line1: line1[i : i+3], line2: line2[i : i+3], line3: line3[i : i+3]}
		if isMatched(zero, digit) {
			fmt.Print("0")
		} else if isMatched(one, digit) {
			fmt.Print("1")
		} else if isMatched(two, digit) {
			fmt.Print("2")
		} else if isMatched(three, digit) {
			fmt.Print("3")
		} else if isMatched(four, digit) {
			fmt.Print("4")
		} else if isMatched(five, digit) {
			fmt.Print("5")
		} else if isMatched(six, digit) {
			fmt.Print("6")
		} else if isMatched(seven, digit) {
			fmt.Print("7")
		} else if isMatched(eight, digit) {
			fmt.Print("8")
		} else if isMatched(nine, digit) {
			fmt.Print("9")
		} else {
			fmt.Print("x")
		}
	}

	fmt.Println()

	// fmt.Fprintln(os.Stderr, "Debug messages...")
}
