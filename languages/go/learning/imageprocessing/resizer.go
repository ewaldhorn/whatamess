package main

import (
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ResizerExample() {
	file, _ := os.Open("test.jpg")
	img, _ := jpeg.Decode(file)

	m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, _ := os.Create("test_resized.jpg")
	defer out.Close()

	jpeg.Encode(out, m, nil)
}
