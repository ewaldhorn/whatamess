package algorithm

import (
	"fmt"
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

// ------------------------------------------------------------------------------------------------
func (b *BreadthFirstSearch) Remove() (*node.Node, error) {
	if len(b.Frontier) > 0 {
		if b.Game.Debug {
			fmt.Println("Frontier before remove:")
			for _, f := range b.Frontier {
				fmt.Println("Node:", f.State)
			}
		}

		node := b.Frontier[0]
		b.Frontier = b.Frontier[1:]
		return node, nil
	}

	return nil, fmt.Errorf("frontier is empty")
}

// ------------------------------------------------------------------------------------------------
