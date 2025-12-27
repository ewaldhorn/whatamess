package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// ------------------------------------------------------------------------------------------------
//
//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")

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
			panic(closeErr) // if closing failures should be fatal, I don't think they should
		}
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
