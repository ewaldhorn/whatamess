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
func makeNewNodeWithDependantsOptimised(name string, dependants ...string) *Node {
	n := &Node{
		Name:       name,
		Dependants: make([]*Node, len(dependants)),
	}

	builder := strings.Builder{}

	for idx, s := range dependants {
		builder.WriteString(name)
		builder.WriteString(":")
		builder.WriteString(s)
		n.Dependants[idx] = &Node{Name: builder.String()}
		builder.Reset()
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
// Building a wordlist trying to preallocate some memory.
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
// Building the same wordlist, but with no attempt at preallocation
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
// Try to build the wordlist without using append
func buildBigWordListNoAppend() []string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	sizeNeeded := len(characters) * 100
	words := make([]string, sizeNeeded)
	pos := 0
	var builder strings.Builder

	for i := range 100 {
		for j, ch := range characters {
			builder.WriteString(string(ch))
			builder.WriteString(strconv.Itoa(i))
			words[i*len(characters)+j] = builder.String() // calculate correct offset

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
// Benchmark the word list building, with minimal appends
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
func makeMassivePopulatedGraphOptimised() *Graph {

	words := buildBigWordListNoAppend()
	graph := Graph{nodes: make([]*Node, len(words))}

	for idx, word := range words {
		graph.nodes[idx] = makeNewNodeWithDependantsOptimised(word, words...)
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
func Benchmark_makeMassivePopulatedGraphOptimised(b *testing.B) {
	for range b.N {
		_ = makeMassivePopulatedGraphOptimised()
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

// ----------------------------------------------------------------------------
func TestBuildBigWordList(t *testing.T) {
	expectedLength := 100 * 52 // 52 characters in the alphabet (26 uppercase + 26 lowercase)
	expectedFirstElement := "A0"
	expectedLastElement := "z99"

	words := buildBigWordList()

	if len(words) != expectedLength {
		t.Errorf("Expected length %d, got %d", expectedLength, len(words))
	}

	if words[0] != expectedFirstElement {
		t.Errorf("Expected first element %q, got %q", expectedFirstElement, words[0])
	}

	if words[len(words)-1] != expectedLastElement {
		t.Errorf("Expected last element %q, got %q", expectedLastElement, words[len(words)-1])
	}
}

// ----------------------------------------------------------------------------
func Test_populatedGraph_LargeOptimised(t *testing.T) {
	// graph := makeMassivePopulatedGraphOptimised()
	graph := makeMassivePopulatedGraph()

	last := len(graph.nodes) - 1
	lastLast := len(graph.nodes[last].Dependants) - 1

	t.Logf("%s, (%d):  %s, (%d)", graph.nodes[last].Name, len(graph.nodes), graph.nodes[last].Dependants[lastLast], len(graph.nodes[last].Dependants))

	lookFor := "z99:z99"
	result := graph.DepthFirstSearch(lookFor)
	if result != nil {
		t.Errorf("failed with nil, expected %s", lookFor)
	}
}
