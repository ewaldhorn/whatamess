package main

import (
	"math/rand"

	"github.com/ewaldhorn/tinycanvas/colour"
)

const colour_count = 10

var colours [colour_count][colour_count]colour.Colour

// ----------------------------------------------------------------------------
func InitColours() {

	for i := range colour_count {
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
}

// ----------------------------------------------------------------------------
func getBaseColour(index int) (uint8, uint8, uint8) {
	switch {
	case index < 2:
		return uint8(128 + rand.Intn(32)), uint8(32 + rand.Intn(32)), 0
	case index < 4:
		return uint8(32 + rand.Intn(32)), uint8(64 + rand.Intn(32)), 0
	case index < 6:
		return 0, uint8(32 + rand.Intn(32)), uint8(32 + rand.Intn(32))
	case index < 8:
		return uint8(150 + rand.Intn(50)), uint8(75 + rand.Intn(25)), 0
	default:
		return uint8(32 + rand.Intn(32)), 0, uint8(32 + rand.Intn(32))
	}
}

// ----------------------------------------------------------------------------
func generateColours(r, g, b uint8) [colour_count]colour.Colour {
	var colours [colour_count]colour.Colour

	for i := 0; i < colour_count; i++ {
		a := uint8(150 + rand.Intn(100))
		switch {
		case r == 0:
			r = uint8(100 + rand.Intn(150))
		case g == 0:
			g = uint8(100 + rand.Intn(150))
		case b == 0:
			b = uint8(100 + rand.Intn(150))
		}
		colours[i] = *colour.NewColour(r, g, b, a)
	}

	return colours
}
