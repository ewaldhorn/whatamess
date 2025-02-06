package main

import (
	"testing"
)

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
