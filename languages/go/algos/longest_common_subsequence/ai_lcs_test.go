package main

import (
	"testing"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected string
	}{
		{"a", "a", "a"},
		{"abc", "def", ""},
		{"abcde", "ace", "ace"},
		{"abcde", "cde", "cde"},
		{"banana", "anana", "anana"},
	}

	for _, test := range tests {
		lcs, err := LCS(test.s1, test.s2)
		if err != nil {
			t.Errorf("LCS(%q, %q) returned error: %v", test.s1, test.s2, err)
			continue
		}
		if lcs != test.expected {
			t.Errorf("LCS(%q, %q) = %q, want %q", test.s1, test.s2, lcs, test.expected)
		}
	}
}
