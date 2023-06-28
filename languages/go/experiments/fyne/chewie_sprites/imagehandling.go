package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func load(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}

	// now that we could read the file, remember to close it!
	defer imgFile.Close()

	imgData, err := png.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}

	return imgData
}
