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

// ------------------------------------------------------------------------------------------------
func TestManhattanDistance(t *testing.T) {
	// Test case 1: Same point (distance should be 0)
	node := &Node{
		State: point.Point{X: 5, Y: 5},
	}
	goal := point.Point{X: 5, Y: 5}
	distance := node.ManhattanDistance(goal)
	expected := 0
	if distance != expected {
		t.Errorf("ManhattanDistance with same point: expected %d, got %d", expected, distance)
	}

	// Test case 2: Horizontal distance only
	node = &Node{
		State: point.Point{X: 2, Y: 5},
	}
	goal = point.Point{X: 8, Y: 5}
	distance = node.ManhattanDistance(goal)
	expected = 6
	if distance != expected {
		t.Errorf("ManhattanDistance with horizontal distance: expected %d, got %d", expected, distance)
	}

	// Test case 3: Vertical distance only
	node = &Node{
		State: point.Point{X: 5, Y: 2},
	}
	goal = point.Point{X: 5, Y: 8}
	distance = node.ManhattanDistance(goal)
	expected = 6
	if distance != expected {
		t.Errorf("ManhattanDistance with vertical distance: expected %d, got %d", expected, distance)
	}

	// Test case 4: Both horizontal and vertical distance
	node = &Node{
		State: point.Point{X: 2, Y: 3},
	}
	goal = point.Point{X: 8, Y: 7}
	distance = node.ManhattanDistance(goal)
	expected = 10 // |2-8| + |3-7| = 6 + 4 = 10
	if distance != expected {
		t.Errorf("ManhattanDistance with both distances: expected %d, got %d", expected, distance)
	}

	// Test case 5: Negative coordinates
	node = &Node{
		State: point.Point{X: -2, Y: -3},
	}
	goal = point.Point{X: 5, Y: 4}
	distance = node.ManhattanDistance(goal)
	expected = 14 // |-2-5| + |-3-4| = 7 + 7 = 14
	if distance != expected {
		t.Errorf("ManhattanDistance with negative coordinates: expected %d, got %d", expected, distance)
	}
}
