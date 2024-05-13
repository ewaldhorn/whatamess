package main

import "testing"

// ----------------------------------------------------------------------------------------- Length
func Test_length(t *testing.T) {
	tests := []struct {
		descr    string
		input    ListNode
		expected int
	}{
		{descr: "One", input: ListNode{}, expected: 1},
		{descr: "Two", input: ListNode{1, &ListNode{}}, expected: 2},
		{descr: "Three", input: ListNode{1, &ListNode{2, &ListNode{}}}, expected: 3},
	}

	for _, test := range tests {
		result := test.input.Length()

		if result != test.expected {
			t.Fatalf("for '%s', we expected %d, got %d\n", test.descr, test.expected, result)
		}
	}
}

// ------------------------------------------------------------------------------------------- Tail
func Test_tail(t *testing.T) {
	tests := []struct {
		descr    string
		input    ListNode
		expected interface{}
	}{
		{descr: "Tail data is nil", input: ListNode{}, expected: nil},
		{descr: "Tail data is 5", input: ListNode{1, &ListNode{5, nil}}, expected: 5},
		{descr: "Tail data is nil", input: ListNode{1, &ListNode{5, &ListNode{}}}, expected: nil},
	}

	for _, test := range tests {
		result := test.input.Tail().Data
		if result != test.expected {
			t.Fatalf("for '%s', expected %v, got %v\n", test.descr, test.expected, result)
		}
	}
}
