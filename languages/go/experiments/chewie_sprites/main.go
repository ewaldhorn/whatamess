package main

import (
	"fmt"
	"image"
	"image/draw"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	fmt.Printf("\nChewie Chew Chew Chew Begins!\n")

	myApp := app.New()
	w := myApp.NewWindow("Chewie Chew Chew Chew!")

	background := load("./assets/background.png")
	playerSprites := load("./assets/chewie.png")

	now := time.Now().UnixMilli()
	game := &Game{
		800,
		500,
		10,
		now,
		4,
	}

	fpsInterval := int64(1000 / game.fps)

	player := &Player{100, 200, 40, 72, 0, 0, 4, 3, 0, 1, 2, 9, 0, 0}

	img := canvas.NewImageFromImage(background)
	img.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	playerImg := canvas.NewRasterFromImage(sprite)
	spriteSize := image.Pt(player.width, player.height)

	c := container.New(layout.NewMaxLayout(), img, playerImg)
	w.SetContent(c)

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyDown:
			if player.y < int(game.canvasHeight)-player.height-game.margin {
				player.yMov = player.speed
			}
			player.frameY = player.downY
		case fyne.KeyUp:
			if player.y > 100 {
				player.yMov = -player.speed
			}
			player.frameY = player.upY
		case fyne.KeyLeft:
			if player.x > game.margin {
				player.xMov = -player.speed
			}
			player.frameY = player.leftY
		case fyne.KeyRight:
			if player.x < int(game.canvasWidth)-player.width-game.margin {
				player.xMov = player.speed
			}
			player.frameY = player.rightY
		}
	})
	go func() {
		for {
			time.Sleep(time.Millisecond)

			now := time.Now().UnixMilli()
			elapsed := now - game.then

			if elapsed > fpsInterval {
				game.then = now

				spriteDP := image.Pt(player.width*player.frameX, player.height*player.frameY)
				sr := image.Rectangle{spriteDP, spriteDP.Add(spriteSize)}

				dp := image.Pt(player.x, player.y)
				r := image.Rectangle{dp, dp.Add(spriteSize)}

				draw.Draw(sprite, sprite.Bounds(), image.Transparent, image.Point{}, draw.Src)
				draw.Draw(sprite, r, playerSprites, sr.Min, draw.Src)
				playerImg = canvas.NewRasterFromImage(sprite)

				if player.xMov != 0 || player.yMov != 0 {
					player.x += player.xMov
					player.y += player.yMov
					player.frameX = (player.frameX + 1) % player.cyclesX
					player.xMov = 0
					player.yMov = 0
				} else {
					player.frameX = 0
				}

				c.Refresh()
			}
		}
	}()

	w.CenterOnScreen()
	w.ShowAndRun()

	fmt.Println("The End.")
}
