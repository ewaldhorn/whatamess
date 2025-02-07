package main

import (
	"testing"
)

// ====================================================================== TESTS
func TestLifoPushPop(t *testing.T) {
	lifo := NewLifo()

	lifo.Push(1)
	lifo.Push(2)
	lifo.Push(3)

	if val := lifo.Pop(); val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}

	if val := lifo.Pop(); val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	if val := lifo.Pop(); val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	if val := lifo.Pop(); val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}

// ----------------------------------------------------------------------------
func TestNewLifo(t *testing.T) {
	l := NewLifo()
	if l == nil {
		t.Errorf("NewLifo returned nil")
	}
	if len(l.items) != 0 {
		t.Errorf("NewLifo returned a non-empty LIFO")
	}
}

// ----------------------------------------------------------------------------
func Test_LifoPush(t *testing.T) {
	l := NewLifo()
	l.Push(1)
	if len(l.items) != 1 {
		t.Errorf("Push did not add an item to the LIFO")
	}
	if l.items[0] != 1 {
		t.Errorf("Push added the wrong item to the LIFO")
	}
}

// ----------------------------------------------------------------------------
func Test_LifoPop(t *testing.T) {
	l := NewLifo()
	l.Push(1)
	l.Push(2)
	x := l.Pop()
	if x != 2 {
		t.Errorf("Pop returned the wrong item")
	}
	if len(l.items) != 1 {
		t.Errorf("Pop did not remove the item from the LIFO")
	}
}

// ----------------------------------------------------------------------------
func Test_LifoPopEmpty(t *testing.T) {
	l := NewLifo()
	x := l.Pop()
	if x != nil {
		t.Errorf("Pop on an empty LIFO did not return nil")
	}
}

// ----------------------------------------------------------------------------
func Test_LifoMultiplePushAndPop(t *testing.T) {
	l := NewLifo()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	x := l.Pop()
	if x != 3 {
		t.Errorf("Pop returned the wrong item")
	}
	x = l.Pop()
	if x != 2 {
		t.Errorf("Pop returned the wrong item")
	}
	x = l.Pop()
	if x != 1 {
		t.Errorf("Pop returned the wrong item")
	}
	x = l.Pop()
	if x != nil {
		t.Errorf("Pop on an empty LIFO did not return nil")
	}
}
