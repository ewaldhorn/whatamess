package main

import "fmt"

/*
Experiments with functional programming in Go.
*/
// ----------------------------------------------------------------------------
type mapFunction[E any] func(E) E

// ----------------------------------------------------------------------------
func Map[S ~[]E, E any](s S, apply mapFunction[E]) (result S) {
	result = make(S, len(s))

	for idx := range s {
		result[idx] = apply(s[idx])
	}

	return
}

// ----------------------------------------------------------------------------
func doubleUp(in int) int {
	return in + in
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("We start with:", numbers)
	fmt.Println("Doubled:", Map(numbers, doubleUp))                         // calling a mapping function
	fmt.Println("Tripled:", Map(numbers, func(s int) int { return s * 3 })) // anonymous mapping function
}
