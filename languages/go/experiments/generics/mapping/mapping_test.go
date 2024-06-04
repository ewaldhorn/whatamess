package main

import (
	"slices"
	"testing"
)

// --------------------------------------------------------------------------- Test_doubling_values
func Test_doubling_values(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected []int32
	}{
		{name: "doubled up", input: []int32{1, 2, 3, 4}, expected: []int32{2, 4, 6, 8}},
	}

	for nr, test := range tests {
		result := Map(test.input, func(i int32) int32 { return i + i })

		if !slices.Equal(test.expected, result) {
			t.Fatalf("test %d failed with result %v instead of %v\n", nr, result, test.expected)
		}
	}
}

// ---------------------------------------------------------------------- Test_adding_one_to_values
func Test_adding_one_to_values(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected []int32
	}{
		{name: "plus one", input: []int32{1, 2, 3, 4}, expected: []int32{2, 3, 4, 5}},
	}

	for nr, test := range tests {
		result := Map(test.input, func(i int32) int32 { return i + 1 })

		if !slices.Equal(test.expected, result) {
			t.Fatalf("test %d failed with result %v instead of %v\n", nr, result, test.expected)
		}
	}
}
