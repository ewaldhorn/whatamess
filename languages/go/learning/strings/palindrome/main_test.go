package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "lol", input: "lol", expected: true},
		{name: "topspot", input: "topspot", expected: true},
		{name: "nope", input: "nope", expected: false},
	}

	for _, test := range tests {
		if test.expected != isPalindrome(test.input) {
			t.Errorf("%s failed: expected %t, got %t", test.name, test.expected, isPalindrome(test.input))
		}
	}
}
