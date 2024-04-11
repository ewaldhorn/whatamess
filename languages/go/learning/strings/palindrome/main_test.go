package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "true: lol", input: "lol", expected: true},
		{name: "true: topspot", input: "topspot", expected: true},
		{name: "false: nope", input: "nope", expected: false},
	}

	for _, test := range tests {
		if test.expected != isPalindrome(test.input) {
			t.Errorf("%s failed: expected %t, got %t", test.name, test.expected, isPalindrome(test.input))
		}
	}
}
