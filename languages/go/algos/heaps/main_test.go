package main

import "testing"

// ----------------------------------------------------------------------------
func Test_IntegerHeap(t *testing.T) {
	var heap IntegerHeap

	if heap.Len() != 0 {
		t.Errorf("Heap should be empty, with a length of 0, at this time, it is %d instead.", heap.Len())
	}

	heap.Push(50)
	heap.Push(80)

	if heap.Len() != 2 {
		t.Errorf("Heap should have two items at this time, it has %d instead.", heap.Len())
	}

	result := heap.Pop()

	if result != 80 {
		t.Errorf("Expected the popped value to be 80, got %d instead", result)
	}

	if heap.Len() != 1 {
		t.Errorf("Heap should have one item at this time, it has %d instead.", heap.Len())
	}

}
