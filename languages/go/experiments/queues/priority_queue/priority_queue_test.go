package main

import (
	"container/heap"
	"testing"
)

// -------------------------------------------------------------------------
func Test_priorityQueue(t *testing.T) {
	// Create a priority queue and add some items
	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Item{value: "task1", priority: 3})
	heap.Push(pq, &Item{value: "task2", priority: 1})
	heap.Push(pq, &Item{value: "task3", priority: 2})

	// Test the length of the priority queue
	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}

	// Test the order of items based on priority
	expectedOrder := []string{"task1", "task3", "task2"}
	for i, expected := range expectedOrder {
		item := heap.Pop(pq).(*Item)
		if item.value != expected {
			t.Errorf("Expected %s, got %s at index %d", expected, item.value, i)
		}
	}

	// Test popping all elements to ensure correct order and length
	if pq.Len() != 0 {
		t.Errorf("Expected length 0, got %d", pq.Len())
	}
}

// -------------------------------------------------------------------------
func Test_PriorityQueue_Len(t *testing.T) {
	pq := &PriorityQueue{}
	if pq.Len() != 0 {
		t.Errorf("Expected length 0, got %d", pq.Len())
	}

	pq = &PriorityQueue{
		&Item{value: "task1", priority: 3},
		&Item{value: "task2", priority: 1},
		&Item{value: "task3", priority: 2},
	}

	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}
}

// -------------------------------------------------------------------------
func Test_PriorityQueue_Less(t *testing.T) {
	pq := &PriorityQueue{
		&Item{value: "task1", priority: 3},
		&Item{value: "task2", priority: 1},
		&Item{value: "task3", priority: 2},
	}

	if !pq.Less(0, 1) {
		t.Errorf("Expected true, got false")
	}
	if pq.Less(1, 0) {
		t.Errorf("Expected false, got true")
	}
	if !pq.Less(2, 1) {
		t.Errorf("Expected true, got false")
	}
	if pq.Less(1, 2) {
		t.Errorf("Expected false, got true")
	}
}
