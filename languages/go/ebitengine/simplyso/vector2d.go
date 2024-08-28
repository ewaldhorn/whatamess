package main

import "math"

type Vector2D struct {
	X float64
	Y float64
}

// ----------------------------------------------------------------------------
// Normalise the vector position based on the X to Y ratio
func (src Vector2D) Normalize() Vector2D {
	ratio := math.Sqrt(src.X*src.X + src.Y*src.Y)
	return Vector2D{src.X / ratio, src.Y / ratio}
}
