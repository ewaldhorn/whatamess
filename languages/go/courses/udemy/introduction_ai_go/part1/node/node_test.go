package node

import (
	"mazes/point"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestNode(t *testing.T) {
	// Test node creation with all fields
	p := point.Point{X: 1, Y: 2}
	parentNode := &Node{
		State:  point.Point{X: 3, Y: 4},
		Action: "up",
	}
	
	node := Node{
		State:  p,
		Parent: parentNode,
		Action: "down",
	}

	// Test State field
	if node.State.X != 1 {
		t.Errorf("Expected State.X to be 1, but got %d", node.State.X)
	}
	
	if node.State.Y != 2 {
		t.Errorf("Expected State.Y to be 2, but got %d", node.State.Y)
	}

	// Test Parent field
	if node.Parent == nil {
		t.Error("Expected Parent to not be nil")
	}
	
	if node.Parent.State.X != 3 {
		t.Errorf("Expected Parent.State.X to be 3, but got %d", node.Parent.State.X)
	}
	
	if node.Parent.State.Y != 4 {
		t.Errorf("Expected Parent.State.Y to be 4, but got %d", node.Parent.State.Y)
	}
	
	if node.Parent.Action != "up" {
		t.Errorf("Expected Parent.Action to be 'up', but got %s", node.Parent.Action)
	}

	// Test Action field
	if node.Action != "down" {
		t.Errorf("Expected Action to be 'down', but got %s", node.Action)
	}
}