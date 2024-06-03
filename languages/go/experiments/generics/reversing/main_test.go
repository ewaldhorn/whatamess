package main

import "testing"

// ------------------------------------------------------------------------------- Test_reverse_int
func Test_reverse_int(t *testing.T) {
	tests := []struct {
		name     string
		input    []int32
		expected []int32
	}{
		{name: "Simple", input: []int32{4, 3, 2, 1}, expected: []int32{1, 2, 3, 4}},
		{name: "Same", input: []int32{1, 2, 2, 1}, expected: []int32{1, 2, 2, 1}},
	}

	for _, test := range tests {
		result := reverse(test.input)

		for i := range len(result) {
			if result[i] != test.expected[i] {
				t.Fatalf("'%s' failed. expected %v, got %v\n", test.name, test.expected, result)
			}
		}
	}
}

// ------------------------------------------------------------------------- Test_reverse_character
func Test_reverse_character(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{name: "Simple", input: []byte{4, 3, 2, 1}, expected: []byte{1, 2, 3, 4}},
		{name: "Same", input: []byte{1, 2, 2, 1}, expected: []byte{1, 2, 2, 1}},
	}

	for _, test := range tests {
		result := reverse(test.input)

		for i := range len(result) {
			if result[i] != test.expected[i] {
				t.Fatalf("'%s' failed. expected %v, got %v\n", test.name, test.expected, result)
			}
		}
	}
}
