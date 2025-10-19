package node

import (
	"mazes/helper"
	"mazes/point"
)

// ------------------------------------------------------------------------------------------------
type Node struct {
	index      int
	State      point.Point
	Parent     *Node
	Action     string
	CostToGoal int
}

// ------------------------------------------------------------------------------------------------
func (n *Node) ManhattanDistance(goal point.Point) int {
	return helper.AbsInt(n.State.Y-goal.Y) + helper.AbsInt(n.State.X-goal.X)
}
