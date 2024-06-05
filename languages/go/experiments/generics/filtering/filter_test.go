package main

import (
	"slices"
	"testing"
)

// ---------------------------------------------------------- Test_filtering_evens_anonymous_filter
func Test_filtering_evens_anonymous_filter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int16
		expected []int16
	}{
		{name: "simple", input: []int16{1, 2, 3, 4, 5, 6}, expected: []int16{2, 4, 6}},
		{name: "easy", input: []int16{2, 2, 3, 4, 5, 6}, expected: []int16{2, 2, 4, 6}},
	}

	for nr, test := range tests {
		result := Filter(test.input, func(value int16) bool { return value%2 == 0 })
		if !slices.Equal(result, test.expected) {
			t.Fatalf("test %d failed with %v instead of %v\n", nr, result, test.expected)
		}
	}
}

// ----------------------------------------------------------- Test_filtering_evens_external_filter
func Test_filtering_evens_external_filter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int16
		expected []int16
	}{
		{name: "simple", input: []int16{1, 2, 3, 4, 5, 6}, expected: []int16{2, 4, 6}},
		{name: "easy", input: []int16{2, 2, 3, 4, 5, 6}, expected: []int16{2, 2, 4, 6}},
	}

	ffunc := func(value int16) bool {
		return value%2 == 0
	}

	for nr, test := range tests {
		result := Filter(test.input, ffunc)
		if !slices.Equal(result, test.expected) {
			t.Fatalf("test %d failed with %v instead of %v\n", nr, result, test.expected)
		}
	}
}

// ------------------------------------------------------------------------- Test_filtering_all_a_s
func Test_filtering_all_a_s(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{name: "abcde", input: []byte{'a', 'b', 'c', 'd'}, expected: []byte{'b', 'c', 'd'}},
	}

	for nr, test := range tests {
		result := Filter(test.input, func(value byte) bool {
			return value != 'a'
		})
		if !slices.Equal(result, test.expected) {
			t.Fatalf("test %d failed with %v instead of %v\n", nr, result, test.expected)
		}
	}
}

// ---------------------------------------------------------------------------- Test_generic_filter
func Test_generic_filter(t *testing.T) {

	tests := []struct {
		name     string
		input    []int16
		expected []int16
	}{
		{name: "simple", input: []int16{1, 2, 3, 4, 5, 6}, expected: []int16{2, 4, 6}},
		{name: "easy", input: []int16{2, 2, 3, 4, 5, 6}, expected: []int16{2, 2, 4, 6}},
	}

	for nr, test := range tests {
		result := Filter(test.input, isEven)
		if !slices.Equal(result, test.expected) {
			t.Fatalf("test %d failed with %v instead of %v\n", nr, result, test.expected)
		}
	}
}
