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
var LaserSprite = mustLoadImage("images/laser.png")
var MeteorSprites = mustLoadImages("images/meteors/*.png")
var MeteorSpritesSmall = mustLoadImages("images/meteors-small/*.png")
var PlayerSprite = mustLoadImage("images/player.png")
var TitleFont = titleFont("fonts/title.ttf")
var ExplosionSprite = mustLoadImage("images/explosion.png")
var ExplosionSpriteSmall = mustLoadImage("images/explosion-small.png")
var Explosion = createExplosion()

// ------------------------------------------------------------------------------------------------
func createExplosion() []*ebiten.Image {
	var frames []*ebiten.Image

	for i := 0; i <= 11; i++ {
		frames = append(frames, mustLoadImage(fmt.Sprintf("images/explosion/%d.png", i+1)))
	}

	return frames
}

// ------------------------------------------------------------------------------------------------
func titleFont(name string) *text.GoTextFace {
	return mustLoadAsset(name, func(f fs.File) (*text.GoTextFace, error) {
		source, err := text.NewGoTextFaceSource(f)
		if err != nil {
			return nil, err
		}

		return &text.GoTextFace{
			Source: source,
			Size:   48,
		}, nil
	})
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
	return mustLoadAsset(name, func(f fs.File) (*ebiten.Image, error) {
		img, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}

		return ebiten.NewImageFromImage(img), nil
	})
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

// ------------------------------------------------------------------------------------------------
func mustLoadAsset[T any](name string, action func(fs.File) (T, error)) T {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file %s: %v", name, err)
		}
	}()

	t, err := action(f)
	if err != nil {
		panic(err)
	}

	return t
}
