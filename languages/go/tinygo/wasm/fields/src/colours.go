package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

const colour_count = 10

var colours [colour_count][colour_count]colour.Colour

// ----------------------------------------------------------------------------
func InitColours() {

	for i := 1; i <= 8; i++ {
		r, g, b := getBaseColour(i)
		colours[i] = generateColours(r, g, b)
	}

	// old school colour palette
	colours[9][0] = *colour.NewColour(32, 32, 128, 255)
	colours[9][1] = *colour.NewColour(32, 32, 150, 255)
	colours[9][2] = *colour.NewColour(32, 32, 200, 255)
	colours[9][3] = *colour.NewColour(80, 110, 240, 255)
	colours[9][4] = *colour.NewColour(128, 128, 255, 128)
	colours[9][5] = *colour.NewColour(90, 90, 215, 255)
	colours[9][6] = *colour.NewColour(90, 90, 215, 128)
	colours[9][7] = *colour.NewColour(155, 155, 245, 255)
	colours[9][8] = *colour.NewColour(64, 64, 128, 128)
	colours[9][9] = *colour.NewColour(64, 64, 175, 255)

	// madness colour palette
	for idx := range 10 {
		colours[0][idx] = *colour.NewColour(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
	}
}

// ----------------------------------------------------------------------------
// Sneaky method to randomise the random colours
func RandomiseRandomColour() {
	for idx := range 10 {
		colours[0][idx] = *colour.NewColour(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)))
	}
}

// ----------------------------------------------------------------------------
func getBaseColour(index int) (uint8, uint8, uint8) {
	switch {
	case index < 2:
		return uint8(64 + rand.Intn(200)), uint8(32 + rand.Intn(150)), 0
	case index < 4:
		return uint8(32 + rand.Intn(64)), uint8(64 + rand.Intn(150)), 0
	case index < 6:
		return 0, uint8(64 + rand.Intn(128)), uint8(120 + rand.Intn(128))
	case index < 8:
		return uint8(50 + rand.Intn(200)), uint8(100 + rand.Intn(145)), 0
	default:
		return uint8(32 + rand.Intn(200)), 0, uint8(32 + rand.Intn(200))
	}
}

// ----------------------------------------------------------------------------
func generateColours(r, g, b uint8) [colour_count]colour.Colour {
	var colours [colour_count]colour.Colour

	for i := 0; i < colour_count; i++ {
		a := uint8(150 + rand.Intn(100))
		switch {
		case r == 0:
			r = uint8(50 + rand.Intn(200))
		case g == 0:
			g = uint8(50 + rand.Intn(200))
		case b == 0:
			b = uint8(50 + rand.Intn(200))
		}
		colours[i] = *colour.NewColour(r, g, b, a)
	}

	return colours
}
