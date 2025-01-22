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
		{"one", []int{1}, math.MinInt},
		{"easy", []int{10, 20, 30, 40, 50}, 40},
		{"duplicate max", []int{10, 20, 30, 40, 50, 50}, 40},
		{"sample", []int{10, 5, 8, 12, 19, 6}, 12},
		{"longer", []int{1, 2, 3, 19, 4, 5, 6, 7, 18, 14, 18, 19, 2, 3, 4, 5, 6, 9}, 18},
		{"all same", []int{10, 10, 10, 10, 10}, 10},
	}

	for _, test := range tests {
		result := getSecondLargestBasic(test.inputs...)

		if result != test.expected {
			t.Errorf("%s: failed with %d, expected %d", test.name, result, test.expected)
		}
	}
}
