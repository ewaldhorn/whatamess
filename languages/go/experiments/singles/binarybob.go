package main

import "fmt"

func main() {
	for i := 0; i <= 99; i++ {
		fmt.Printf("%2d = %8b in binary\n", i, i)
	}
}
