package helper

import (
	"mazes/globals"
	"mazes/point"
	"os"
)

// ------------------------------------------------------------------------------------------------
func ContainsPoint(needle point.Point, haystack []point.Point) bool {
	for _, h := range haystack {
		if h.Y == needle.Y && h.X == needle.X {
			return true
		}
	}
	return false
}

// ------------------------------------------------------------------------------------------------
func ClearTempDirectory() {
	directory := globals.TemporaryDirectory
	dir, _ := os.Open(directory)
	filesToDelete, _ := dir.ReadDir(0)

	for idx := range filesToDelete {
		f := filesToDelete[idx]
		fullPath := directory + f.Name()
		_ = os.Remove(fullPath)
	}
}

// ------------------------------------------------------------------------------------------------
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
