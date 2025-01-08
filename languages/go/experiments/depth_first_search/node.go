package main

// ----------------------------------------------------------------------------
type Node struct {
	Name       string
	Dependants []*Node
}

// ----------------------------------------------------------------------------
type Graph struct {
	nodes []*Node
}
