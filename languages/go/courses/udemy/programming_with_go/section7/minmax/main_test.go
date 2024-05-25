package main

import "testing"

func Test_minMaxOf(t *testing.T) {
	tests := []struct {
		input []int
		max   int
		min   int
	}{
		{input: []int{1, 0, 1}, max: 1, min: 0},
		{input: []int{1, 1, 1}, max: 1, min: 1},
		{input: []int{1, 2, 3}, max: 3, min: 1},
		{input: []int{7, 2, 3, 4, 5, 6}, max: 7, min: 2},
	}

	for _, test := range tests {
		min, max := minMaxOf(test.input)

		if min != test.min || max != test.max {
			t.Fatalf("expected min:%d and max:%d for %v, got %d/%d\n", test.min, test.max, test.input, min, max)
		}
	}
}
