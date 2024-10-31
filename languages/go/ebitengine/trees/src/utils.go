package main

import (
	"math"
	"math/rand"
)

// ----------------------------------------------------------------------------
// Clamps the value to the min/max
func clampFloat32(value, min, max float32) float32 {
	if min > max {
		panic("min must less than max")
	}

	if min == max {
		return min
	}

	if value <= min {
		return min
	}

	if max <= value {
		return max
	}

	return value
}

// ----------------------------------------------------------------------------
func randomBetween(start float32, end float32) float32 {
	return clampFloat32(start+((start+end)*rand.Float32()), start, end)
}

// ----------------------------------------------------------------------------
func rotateBy(x float32, y float32, len float32, degrees float32) (float32, float32) {
	newX := x + (len * float32(math.Cos(float64(degrees*math.Pi/180))))
	newY := y - (len * float32(math.Sin(float64(degrees*math.Pi/180))))

	return newX, newY
}
