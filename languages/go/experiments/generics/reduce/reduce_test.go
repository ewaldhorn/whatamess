package main

import "testing"

// ------------------------------------------------------------------------------ Test_summing_ints
func Test_summing_ints(t *testing.T) {
	tests := []struct {
		input    []int32
		expected int32
	}{
		{input: []int32{5, 4, 8, 6, 2}, expected: 25},
		{input: []int32{1, 2, 3, 4, 5}, expected: 15},
	}

	for nr, test := range tests {
		result := Fold(test.input, 0, func(current, next int32) int32 { return current + next })

		if result != test.expected {
			t.Fatalf("test %d failed with %d instead of %d\n", nr, result, test.expected)
		}
	}
}

// ------------------------------------------------------------------ Test_calculating_product_ints
func Test_calculating_product_ints(t *testing.T) {
	tests := []struct {
		input    []int32
		expected int32
	}{
		{input: []int32{1, 2, 3}, expected: 6},
		{input: []int32{1, 2, 2}, expected: 4},
		{input: []int32{3, 2, 1}, expected: 6},
	}

	multiplier := func(current, next int32) int32 {
		return current * next
	}

	for nr, test := range tests {
		result := Fold(test.input, 1, multiplier) // always start at 1 for this...

		if result != test.expected {
			t.Fatalf("test %d failed with %d instead of %d\n", nr, result, test.expected)
		}
	}
}

// ------------------------------------------------------------------------------ Test_with_strings
func Test_with_strings(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{input: []string{"o", "n", "e"}, expected: "one"},
		{input: []string{"o", "o", "o"}, expected: "ooo"},
	}

	for nr, test := range tests {
		result := Fold(test.input, "", func(cur, nxt string) string { return cur + nxt })

		if result != test.expected {
			t.Fatalf("test %d failed with %v instead of %v\n", nr, result, test.expected)
		}
	}
}
