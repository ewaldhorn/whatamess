package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ------------------------------------------------------------------------------------------------
//
//go:embed *
var assets embed.FS

// ------------------------------------------------------------------------------------------------
var PlayerSprite = mustLoadImage("images/player.png")
var TitleFont = titleFont("fonts/title.ttf")

// ------------------------------------------------------------------------------------------------
func titleFont(name string) *text.GoTextFace {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	source, err := text.NewGoTextFaceSource(f)
	if err != nil {
		panic(err)
	}

	face := &text.GoTextFace{
		Source: source,
		Size:   48,
	}

	return face
}

// ------------------------------------------------------------------------------------------------
func ReportAssets() {
	err := fs.WalkDir(assets, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			fmt.Println("DIR", d)
		} else {
			fmt.Println("\tFILE", path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error reading embedded files:", err)
	}
}

// ------------------------------------------------------------------------------------------------
func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			log.Printf("failed to close file %s: %v", name, closeErr)
			// panic(closeErr) // if closing failures should be fatal, I don't think they should
		}
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

// ------------------------------------------------------------------------------------------------
// loads all the images in the specified path into an array
func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}
