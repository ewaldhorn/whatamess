package main

import "testing"

func Test_sumOfDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 15, expected: 6},
		{input: 123, expected: 6},
		{input: 132, expected: 6},
		{input: 312, expected: 6},
		{input: 321, expected: 6},
		{input: 1234, expected: 10},
		{input: 123456, expected: 21},
	}

	for _, test := range tests {
		result := sumOfDigits(test.input)

		if result != test.expected {
			t.Fatalf("expected %d for %d, got %d\n", test.expected, test.input, result)
		}
	}
}
