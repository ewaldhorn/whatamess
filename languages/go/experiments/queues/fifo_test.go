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

func TestNewFifo(t *testing.T) {
	f := NewFifo()
	if f == nil {
		t.Errorf("NewFifo returned nil")
	}
	if len(f.items) != 0 {
		t.Errorf("NewFifo returned a non-empty FIFO")
	}
}

func TestPush(t *testing.T) {
	f := NewFifo()
	f.Push(1)
	if len(f.items) != 1 {
		t.Errorf("Push did not add an item to the FIFO")
	}
	if f.items[0] != 1 {
		t.Errorf("Push added the wrong item to the FIFO")
	}
}

func TestPop(t *testing.T) {
	f := NewFifo()
	f.Push(1)
	f.Push(2)
	x := f.Pop()
	if x != 1 {
		t.Errorf("Pop returned the wrong item")
	}
	if len(f.items) != 1 {
		t.Errorf("Pop did not remove the item from the FIFO")
	}
}

func TestPopEmpty(t *testing.T) {
	f := NewFifo()
	x := f.Pop()
	if x != nil {
		t.Errorf("Pop on an empty FIFO did not return nil")
	}
}

func TestMultiplePushAndPop(t *testing.T) {
	f := NewFifo()
	f.Push(1)
	f.Push(2)
	f.Push(3)
	x := f.Pop()
	if x != 1 {
		t.Errorf("Pop returned the wrong item")
	}
	x = f.Pop()
	if x != 2 {
		t.Errorf("Pop returned the wrong item")
	}
	x = f.Pop()
	if x != 3 {
		t.Errorf("Pop returned the wrong item")
	}
	x = f.Pop()
	if x != nil {
		t.Errorf("Pop on an empty FIFO did not return nil")
	}
}
