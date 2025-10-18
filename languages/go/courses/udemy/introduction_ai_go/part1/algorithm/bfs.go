package algorithm

import (
	"mazes/maze"
	"mazes/node"
)

// ------------------------------------------------------------------------------------------------
type BreadthFirstSearch struct {
	Frontier []*node.Node
	Game     *maze.Maze
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) GetFrontier() []*node.Node {
	return b.Frontier
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) AddNode(i *node.Node) {
	b.Frontier = append(b.Frontier, i)
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) ContainsState(i *node.Node) bool {
	for _, x := range b.Frontier {
		if x.State == i.State {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) IsEmpty() bool {
	return len(b.Frontier) == 0
}
