package algorithm

import (
	"mazes/maze"
	"mazes/node"
	"mazes/point"
	"testing"
)

// ------------------------------------------------------------------------------------------------
func TestDepthFirstSearch(t *testing.T) {
	dfs := &DepthFirstSearch{Game: &maze.Maze{}}

	t.Run("is empty on initialization", func(t *testing.T) {
		if !dfs.IsEmpty() {
			t.Error("Expected frontier to be empty on initialization")
		}
	})

	t.Run("is not empty after adding a node", func(t *testing.T) {
		dfs := &DepthFirstSearch{Game: &maze.Maze{}}
		node1 := &node.Node{State: point.Point{Y: 1, X: 1}}
		dfs.Add(node1)
		if dfs.IsEmpty() {
			t.Error("Expected frontier to not be empty after adding a node")
		}
	})

	t.Run("add and remove nodes (LIFO)", func(t *testing.T) {
		dfs := &DepthFirstSearch{Game: &maze.Maze{}}
		node1 := &node.Node{State: point.Point{Y: 1, X: 1}}
		node2 := &node.Node{State: point.Point{Y: 2, X: 2}}
		node3 := &node.Node{State: point.Point{Y: 3, X: 3}}

		dfs.Add(node1)
		dfs.Add(node2)
		dfs.Add(node3)

		removedNode, err := dfs.Remove()
		if err != nil {
			t.Fatalf("Remove failed: %v", err)
		}
		if removedNode != node3 {
			t.Errorf("Expected to remove node3, but got %+v", removedNode)
		}

		removedNode, err = dfs.Remove()
		if err != nil {
			t.Fatalf("Remove failed: %v", err)
		}
		if removedNode != node2 {
			t.Errorf("Expected to remove node2, but got %+v", removedNode)
		}

		removedNode, err = dfs.Remove()
		if err != nil {
			t.Fatalf("Remove failed: %v", err)
		}
		if removedNode != node1 {
			t.Errorf("Expected to remove node1, but got %+v", removedNode)
		}

		if !dfs.IsEmpty() {
			t.Error("Expected frontier to be empty after removing all nodes")
		}
	})

	t.Run("remove from empty frontier", func(t *testing.T) {
		dfs := &DepthFirstSearch{Game: &maze.Maze{}}
		_, err := dfs.Remove()
		if err == nil {
			t.Error("Expected an error when removing from an empty frontier, but got nil")
		}
	})

	t.Run("contains state", func(t *testing.T) {
		dfs := &DepthFirstSearch{Game: &maze.Maze{}}
		node1 := &node.Node{State: point.Point{Y: 1, X: 1}}
		node2 := &node.Node{State: point.Point{Y: 2, X: 2}}
		nodeIn := &node.Node{State: point.Point{Y: 1, X: 1}} // Same state as node1
		nodeOut := &node.Node{State: point.Point{Y: 3, X: 3}}

		dfs.Add(node1)
		dfs.Add(node2)

		if !dfs.ContainsState(nodeIn) {
			t.Error("Expected frontier to contain state of nodeIn")
		}

		if dfs.ContainsState(nodeOut) {
			t.Error("Expected frontier to not contain state of nodeOut")
		}
	})

	t.Run("get frontier", func(t *testing.T) {
		dfs := &DepthFirstSearch{Game: &maze.Maze{}}
		node1 := &node.Node{State: point.Point{Y: 1, X: 1}}
		node2 := &node.Node{State: point.Point{Y: 2, X: 2}}

		dfs.Add(node1)
		dfs.Add(node2)

		frontier := dfs.GetFrontier()
		if len(frontier) != 2 {
			t.Fatalf("Expected frontier length of 2, got %d", len(frontier))
		}
		if frontier[0] != node1 {
			t.Error("Frontier content mismatch at index 0")
		}
		if frontier[1] != node2 {
			t.Error("Frontier content mismatch at index 1")
		}
	})
}
