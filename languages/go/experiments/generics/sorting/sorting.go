package main

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Skips the boilerplate for ordered entities like numbers
func EasySort[E constraints.Ordered](collection []E) []E {
	// make a copy of the slice, so that we can manipulate positions
	result := make([]E, len(collection))
	copy(result, collection)

	// use the built-in sort, with a comparison function
	sort.Slice(result, func(left, right int) bool {
		return result[left] < result[right]
	})

	// return the result
	return result
}
