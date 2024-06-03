package main

import "testing"

// ----------------------------------------------------------------------- Test_contains_slice_ints
func Test_contains_slice_ints(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		input    []int32
		toFind   int32
	}{

		{name: "find ints", expected: true, input: []int32{1, 2, 3, 4, 5}, toFind: 3},
		{name: "find ints - not found", expected: false, input: []int32{1, 2, 3, 4, 5}, toFind: 0},
	}

	for _, test := range tests {
		result := contains(test.input, test.toFind)
		if result != test.expected {
			t.Fatalf("error in test '%s'.", test.name)
		}
	}
}

// ---------------------------------------------------------------------- Test_contains_slice_bools
func Test_contains_slice_bools(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		input    []bool
		toFind   bool
	}{

		{name: "find ints", expected: true, input: []bool{false, false, true}, toFind: true},
		{name: "find ints - not found", expected: false, input: []bool{false, false}, toFind: true},
	}

	for _, test := range tests {
		result := contains(test.input, test.toFind)
		if result != test.expected {
			t.Fatalf("error in test '%s'.", test.name)
		}
	}
}
