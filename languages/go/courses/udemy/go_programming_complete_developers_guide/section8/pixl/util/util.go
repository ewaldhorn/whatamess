package util

import (
	"image"
	"image/color"
)

func GetImageColours(img image.Image) map[color.Color]struct{} {
	colours := make(map[color.Color]struct{})
	var empty struct{}

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			colours[img.At(x, y)] = empty
		}
	}

	return colours
}
