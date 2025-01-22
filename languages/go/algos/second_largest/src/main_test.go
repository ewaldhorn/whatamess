package main

import (
	"math"
	"testing"
)

func Test_getSecondLargest(t *testing.T) {
	tests := []struct {
		name     string
		inputs   []int
		expected int
	}{
		{"empty", []int{}, math.MinInt},
		{"one", []int{1}, 1},
		{"sample", []int{10, 5, 8, 12, 19, 6}, 12},
	}

	for _, test := range tests {
		result := getSecondLargest(test.inputs...)

		if result != test.expected {
			t.Errorf("%s: failed with %d, expected %d", test.name, result, test.expected)
		}
	}
}
