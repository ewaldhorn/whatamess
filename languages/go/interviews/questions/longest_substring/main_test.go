package main

import "testing"

func Test_findLongestSubstring(t *testing.T) {
	tests := []struct {
		input          string
		expectedLength int32
		expectedString string
	}{
		{input: "abcabcbb", expectedLength: 3, expectedString: "abc"},
		{input: "bbbbb", expectedLength: 1, expectedString: "b"},
		{input: "pwwkew", expectedLength: 3, expectedString: "wke"},
	}

	for _, test := range tests {
		resultLength := findLongestSubString(test.input)

		if resultLength != test.expectedLength {
			t.Errorf("for '%s' we expected %d, got %d\n", test.input, test.expectedLength, resultLength)
		}
	}
}
