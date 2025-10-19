package node

import "mazes/point"

// ------------------------------------------------------------------------------------------------
type Node struct {
	index      int
	State      point.Point
	Parent     *Node
	Action     string
	CostToGoal int
}
