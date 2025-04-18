package main

import "golang.org/x/exp/constraints"

// define a type for the filter function to establish a contract
type filterFunction[E any] func(E) bool

// use generics to apply a given filter function to a slice
func Filter[E any](slice []E, filterFunction filterFunction[E]) []E {
	result := []E{}

	for _, v := range slice {
		if filterFunction(v) {
			result = append(result, v)
		}
	}

	return result
}

// create a comparison function using generics
func isEven[T constraints.Integer](value T) bool {
	return value%2 == 0
}
