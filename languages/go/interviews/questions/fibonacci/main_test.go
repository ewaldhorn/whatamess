package main

import "testing"

// use the same input/expected results for all tests
var tests = []struct {
	input    int
	expected int
}{
	{input: 0, expected: 0},
	{input: 1, expected: 1},
	{input: 2, expected: 1},
	{input: 3, expected: 2},
	{input: 4, expected: 3},
	{input: 5, expected: 5},
	{input: 6, expected: 8},
	{input: 7, expected: 13},
	{input: 8, expected: 21},
	{input: 9, expected: 34},
	{input: 10, expected: 55},
	{input: 11, expected: 89},
	{input: 12, expected: 144},
	{input: 13, expected: 233},
	{input: 14, expected: 377},
	{input: 15, expected: 610},
}

// -------------------------------------------------------------------------------------- Recursion
func Test_recursion(t *testing.T) {
	for _, test := range tests {
		result := fibonacci_recursive(test.input)

		if result != test.expected {
			t.Fatalf("recursion for %d, expected %d, got %d", test.input, test.expected, result)
		}
	}
}

// -------------------------------------------------------------------------------------- Iteration
func Test_iterative(t *testing.T) {
	for _, test := range tests {
		result := fibonacci_iterative(test.input)

		if result != test.expected {
			t.Fatalf("iterative for %d, expected %d, got %d", test.input, test.expected, result)
		}
	}
}

// ----------------------------------------------------------------------------------- Golden Ratio
func Test_golden_ration(t *testing.T) {
	for _, test := range tests {
		result := fibonacci_golden_ration(test.input)

		if result != test.expected {
			t.Fatalf("golden ratio for %d, expected %d, got %d", test.input, test.expected, result)
		}
	}
}
