package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyArray[T constraints.Ordered] struct {
	inner []T
}

func (m *MyArray[T]) Max() T {
	max := m.inner[0]

	for _, v := range m.inner {
		if v > max {
			max = v
		}
	}

	return max
}

func testMore() {
	fmt.Printf("\nTesting even more constraints\n")

	arr := MyArray[int]{inner: []int{1, 3, 5, 6, 2, 8, 9, 2, 4, 1, 1, 7, 5, 4}}
	fmt.Println("The max is", arr.Max())
}
