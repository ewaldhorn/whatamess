package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width   = 800
	height  = 800
	maxIter = 255
)

func main() {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			cx := float64(x-width/2) * 4.0 / float64(width)
			cy := float64(y-height/2) * 4.0 / float64(height)
			iter := mandelbrot(cx, cy, maxIter)
			col := color.NRGBA{uint8(iter), uint8(iter), uint8(iter), 255}
			img.SetNRGBA(x, y, col)
		}
	}

	f, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}

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
