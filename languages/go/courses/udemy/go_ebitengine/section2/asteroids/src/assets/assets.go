package assets

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")

// ----------------------------------------------------------------------------
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
		fmt.Println("Error reading embbeded files:", err)
	}
}

// ----------------------------------------------------------------------------
func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
