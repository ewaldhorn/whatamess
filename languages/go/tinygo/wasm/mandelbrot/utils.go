package main

import (
	"math"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
// Converts float from range min-maxF to fixed point 12 decimal places
// representation in range 0-max.
func intToFixed(i, max int, min, maxF float64) int64 {
	f := float64(i)/float64(max)*(maxF-min) + min
	return int64(f * 1e12)
}

// ----------------------------------------------------------------------------
// Converts fixed point 12 decimal places representation to int.
func fixedToInt(f int64) int {
	return int(f / 1e12)
}

// ----------------------------------------------------------------------------
// Converts fixed point 12 decimal places representation to float.
func fixedToFloat(f int64) float64 {
	return float64(f) / 1e12
}

// ----------------------------------------------------------------------------
// Performs a single Mandelbrot calculation.
func mandelbrot(cx, cy float64, maxIter int) int {
	x, y := 0.0, 0.0
	iter := 0

	for iter < maxIter && (x*x+y*y) < 4.0 {
		xNew := x*x - y*y + cx
		y = 2*x*y + cy
		x = xNew
		iter++
	}

	return iter
}

// ----------------------------------------------------------------------------
func getColour(iter, maxIter int) *colour.Colour {
	if iter == maxIter {
		return colour.NewColourBlack() // Black for points in the set
	}

	// Color gradient for points outside the set
	r := uint8(float64(iter) / float64(maxIter) * 255)
	g := uint8(math.Sin(float64(iter)/float64(maxIter)*math.Pi*2)*128 + 128)
	b := uint8(255 - r)
	return colour.NewColour(r, g, b, colour.MAX_COLOUR_VALUE)
}
