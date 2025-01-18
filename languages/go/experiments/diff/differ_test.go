package main

import "testing"

// ----------------------------------------------------------------------------
func Test_findLCS(t *testing.T) {
	tests := []struct {
		name, left, right, expected string
	}{
		{name: "Same", left: "ABCDEF", right: "ABCDEF", expected: "ABCDEF"},
		{name: "Nothing", left: "ABC", right: "DEF", expected: ""},
		{name: "Just last two", left: "ABCDE", right: "XYZDE", expected: "DE"},
		{name: "Blank", left: "", right: "", expected: ""},
		{name: "Some characters", left: "ABCDE", right: "AC", expected: "AC"},
	}

	for _, test := range tests {
		result := findLCS(test.left, test.right)
		if result != test.expected {
			t.Errorf("%s: failed with '%s', expected '%s'.", test.name, result, test.expected)
		}
	}
}
