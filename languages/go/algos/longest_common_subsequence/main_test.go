package main

import (
	"testing"
)

// ----------------------------------------------------------------------------
func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"", "", 0},
		{"a", "a", 1},
		{"abc", "def", 0},
		{"abcde", "ace", 3},
		{"abcde", "cde", 3},
		{"banana", "anana", 5},
	}

	for _, test := range tests {
		lcs := longestCommonSubsequence(test.s1, test.s2, len(test.s1), len(test.s2))
		if lcs != test.expected {
			t.Errorf("LCS(%q, %q) = %q, want %q", test.s1, test.s2, lcs, test.expected)
		}
	}
}
