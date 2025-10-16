package helper

import "mazes/point"

// ------------------------------------------------------------------------------------------------
func ContainsPoint(needle point.Point, haystack []point.Point) bool {
	for _, h := range haystack {
		if h.Y == needle.Y && h.X == needle.X {
			return true
		}
	}
	return false
}
