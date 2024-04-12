package main

import "testing"

func Test_check(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 1, expected: "1"},
		{input: 2, expected: "2"},
		{input: 3, expected: "Fizz"},
		{input: 4, expected: "4"},
		{input: 5, expected: "Buzz"},
		{input: 6, expected: "Fizz"},
		{input: 7, expected: "7"},
		{input: 8, expected: "8"},
		{input: 9, expected: "Fizz"},
		{input: 10, expected: "Buzz"},
		{input: 11, expected: "11"},
		{input: 12, expected: "Fizz"},
		{input: 13, expected: "13"},
		{input: 14, expected: "14"},
		{input: 15, expected: "Fizz Buzz"},
	}

	for _, test := range tests {
		result := check(test.input)

		if result != test.expected {
			t.Errorf("%d resulted in '%s', not '%s'\n", test.input, result, test.expected)
		}
	}
}
