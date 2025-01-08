package main

import (
	"fmt"
	"testing"
)

// ----------------------------------------------------------------------------
func makeNewNodeWithDependants(name string, dependants ...string) *Node {
	n := &Node{Name: name}

	for _, s := range dependants {
		n.Dependants = append(n.Dependants, &Node{Name: fmt.Sprintf("%s:%s", name, s)})
	}

	return n
}

// ----------------------------------------------------------------------------
func Test_emptyGraph(t *testing.T) {
	graph := Graph{nodes: []*Node{makeNewNodeWithDependants("Dacia")}}

	result := graph.nodes[0].DepthFirstSearch("Blinky")

	if result != nil {
		t.Errorf("emptyGraph test failed! Expected nil, got %v", result)
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraphSingleNode(t *testing.T) {
	graph := Graph{}

	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("One", "One", "Two", "Three"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Two", "One", "Two", "Three", "Four"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Three", "One", "Two", "Three", "Four", "Five", "Six"))

	result := graph.nodes[0].DepthFirstSearch("One:Two")

	if result == nil {
		t.Errorf("failed with nil, expected %s", "One:Two")
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph(t *testing.T) {
	graph := Graph{}

	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("One", "One", "Two", "Three"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Two", "One", "Two", "Three", "Four"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Three", "One", "Two", "Three", "Four", "Five", "Six"))

	result := graph.DepthFirstSearch("Two:Two")

	if result == nil {
		t.Errorf("failed with nil, expected %s", "One:Two")
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph_NotFound(t *testing.T) {
	graph := Graph{}

	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("One", "One", "Two", "Three"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Two", "One", "Two", "Three", "Four"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Three", "One", "Two", "Three", "Four", "Five", "Six"))

	result := graph.DepthFirstSearch("Six:Two")

	if result != nil {
		t.Errorf("failed with %s, expected nil", result.Name)
	}
}
