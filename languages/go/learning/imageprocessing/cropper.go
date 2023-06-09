package main

import (
	"image"
	"image/jpeg"
	"os"
)

func CropperExample() {
	file, _ := os.Open("test.jpg")
	img, _, _ := image.Decode(file)

	// Define rectangle for cropping.
	rect := image.Rect(200, 200, 400, 400)

	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)

	out, _ := os.Create("test_cropped.jpg")
	defer out.Close()

	jpeg.Encode(out, croppedImg, nil)
}
