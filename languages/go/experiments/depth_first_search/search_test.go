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
func makePopulatedGraph() *Graph {
	graph := Graph{}

	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("One", "One", "Two", "Three"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Two", "One", "Two", "Three", "Four"))
	graph.nodes = append(graph.nodes, makeNewNodeWithDependants("Three", "One", "Two", "Three", "Four", "Five", "Six"))

	return &graph
}

// ----------------------------------------------------------------------------
func buildBigWordList() []string {
	words := make([]string, 0, 15_000)
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := range 100 {
		for _, ch := range characters {
			words = append(words, fmt.Sprintf("%c%d", ch, i))
		}
	}

	return words
}

// ----------------------------------------------------------------------------
func makeMassivePopulatedGraph() *Graph {
	graph := Graph{}

	words := buildBigWordList()

	for _, word := range words {
		graph.nodes = append(graph.nodes, makeNewNodeWithDependants(word, words...))
	}

	return &graph
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
	graph := makePopulatedGraph()

	result := graph.nodes[0].DepthFirstSearch("One:Two")

	if result == nil {
		t.Errorf("failed with nil, expected %s", "One:Two")
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph(t *testing.T) {
	graph := makePopulatedGraph()

	result := graph.DepthFirstSearch("Two:Two")

	if result == nil {
		t.Errorf("failed with nil, expected %s", "One:Two")
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph_NotFound(t *testing.T) {
	graph := makePopulatedGraph()

	result := graph.DepthFirstSearch("Six:Two")

	if result != nil {
		t.Errorf("failed with %s, expected nil", result.Name)
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph_Large(t *testing.T) {
	graph := makeMassivePopulatedGraph()

	lookFor := "z99:z99"
	result := graph.DepthFirstSearch(lookFor)

	if result == nil {
		t.Errorf("failed with nil, expected %s", lookFor)
	}
}
