package main

import (
	"fmt"
	"strconv"
	"strings"
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
func buildBigWordListNoPre() []string {
	var words []string
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	for i := range 100 {
		for _, ch := range characters {
			words = append(words, fmt.Sprintf("%c%d", ch, i))
		}
	}

	return words
}

// ----------------------------------------------------------------------------
func buildBigWordListNoAppend() []string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	sizeNeeded := len(characters) * 100
	words := make([]string, sizeNeeded)
	pos := 0
	var builder strings.Builder

	for i := range 100 {
		for _, ch := range characters {
			builder.WriteString(string(ch))
			builder.WriteString(strconv.Itoa(i))
			words[pos] = builder.String()

			builder.Reset()
			pos += 1
		}
	}

	return words
}

// ----------------------------------------------------------------------------
// Benchmark the word list building, just for giggles
func Benchmark_buildBigWordList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildBigWordList()
	}
}

// ----------------------------------------------------------------------------
// Benchmark the word list building, just for giggles
func Benchmark_buildBigWordListNoPre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildBigWordListNoPre()
	}
}

// ----------------------------------------------------------------------------
// Benchmark the word list building, just for giggles
func Benchmark_buildBigWordListNoAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = buildBigWordListNoAppend()
	}
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
// Also benchmark the graph creation function
func Benchmark_makeMassivePopulatedGraph(b *testing.B) {
	for range b.N {
		_ = makeMassivePopulatedGraph()
	}
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
