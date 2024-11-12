package main

import "testing"

// ----------------------------------------------------------------------------
func Test_lengthOfLongestNonRepeatingSubString(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{"easy", "abcabcabc", 3},
		{"sneaky", "abcabcdabc", 4},
		{"more", "ababababababcdefgabcabcdabcdeabcdefghababcabcdeabcdefabcdefg", 8},
	}

	for _, test := range tests {
		result := lengthOfLongestNonRepeatingSubString(test.input)

		if result != test.expected {
			t.Fatalf("%s: expected %d, got %d", test.name, test.expected, result)
		}
	}
}

// ----------------------------------------------------------------------------
func Test_lengthOfLongestNonRepeatingSubStringWithMap(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{"easy", "abcabcabc", 3},
		{"sneaky", "abcabcdabc", 4},
		{"more", "ababababababcdefgabcabcdabcdeabcdefghababcabcdeabcdefabcdefg", 8},
	}

	for _, test := range tests {
		result := lengthOfLongestNonRepeatingSubStringWithMap(test.input)

		if result != test.expected {
			t.Fatalf("%s: expected %d, got %d", test.name, test.expected, result)
		}
	}

}
