package main

import (
	"testing"
)

func TestFifoPushPop(t *testing.T) {
	fifo := NewFifo()

	// Test Push
	fifo.Push(1)
	fifo.Push(2)
	fifo.Push(3)

	// Test Pop
	if item := fifo.Pop(); item != 1 {
		t.Errorf("Expected 1, got %v", item)
	}
	if item := fifo.Pop(); item != 2 {
		t.Errorf("Expected 2, got %v", item)
	}
	if item := fifo.Pop(); item != 3 {
		t.Errorf("Expected 3, got %v", item)
	}
	if item := fifo.Pop(); item != nil {
		t.Errorf("Expected nil, got %v", item)
	}
}

func TestFifoEmptyPop(t *testing.T) {
	fifo := NewFifo()

	// Test Pop on empty fifo
	if item := fifo.Pop(); item != nil {
		t.Errorf("Expected nil, got %v", item)
	}
}
