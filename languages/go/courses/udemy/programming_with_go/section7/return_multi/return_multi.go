package main

import "fmt"

// ----------------------------------------------------------------------------
func main() {
	val := 10
	dbl, trp := return_double_triple(val)
	fmt.Printf("For %d, double is %d and triple is %d\n", val, dbl, trp)
}

// ----------------------------------------------------------------------------
func return_double_triple(input int) (int, int) {
	return input * 2, input * 3
}
