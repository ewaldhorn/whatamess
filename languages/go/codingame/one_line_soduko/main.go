package main

import (
	"fmt"
	"strings"
)

func main() {
	all := "0123456789"
	sample := "102357849"
	fmt.Println()
	fmt.Println("Sample string: ", sample, "is missing '6'.")
	fmt.Println("To find difference with", all, ".")
	fmt.Println()

	for i := 0; i < len(all); i++ {
		r := all[i]
		if !strings.Contains(sample, string(r)) {
			fmt.Println(string(r))
		}
	}
}
