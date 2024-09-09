package main

/*
Experiments with functional programming in Go.

Bespoke Map, Filter, Reduce functions.
*/

import "fmt"

// ----------------------------------------------------------------------------
type mapFunction[E any] func(E) E
type keepFunction[E any] func(E) bool

// ----------------------------------------------------------------------------
func Map[S ~[]E, E any](source S, apply mapFunction[E]) (result S) {
	result = make(S, len(source))

	for idx := range source {
		result[idx] = apply(source[idx])
	}

	return
}

// ----------------------------------------------------------------------------
func doubleUp(in int) int {
	return in + in
}

// ----------------------------------------------------------------------------
func Filter[SourceType ~[]E, E any](source SourceType, mustKeep keepFunction[E]) (result SourceType) {
	result = SourceType{}

	for _, element := range source {
		if mustKeep(element) {
			result = append(result, element)
		}
	}

	return
}

// ----------------------------------------------------------------------------
func IsDivisiblebyTwo(num int) bool {
	return num%2 == 0
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("We start with :", numbers)
	fmt.Println("Doubled (Map) :", Map(numbers, doubleUp))                         // calling a mapping function
	fmt.Println("Tripled (Map) :", Map(numbers, func(s int) int { return s * 3 })) // anonymous mapping function
	fmt.Println("We still have :", numbers)
	fmt.Println("Evens (Filter):", Filter(numbers, IsDivisiblebyTwo))
	fmt.Println("Odds (Filter) :", Filter(numbers, func(in int) bool { return in%2 != 0 }))
	fmt.Println("We still have :", numbers)
}
