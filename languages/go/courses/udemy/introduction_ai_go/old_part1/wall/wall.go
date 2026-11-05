package wall

import "mazes/point"

// ------------------------------------------------------------------------------------------------
type Wall struct {
	State   point.Point
	IsSolid bool
}
