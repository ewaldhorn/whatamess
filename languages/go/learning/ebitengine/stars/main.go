// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	scale        = 16
	starsCount   = 2048
)

type Star struct {
	fromX, fromY, toX, toY, brightness float32
}

func (s *Star) Init() {
	s.toX = rand.Float32() * screenWidth * scale
	s.fromX = s.toX
	s.toY = rand.Float32() * screenHeight * scale
	s.fromY = s.toY
	s.brightness = rand.Float32() * 0xff
}

func (s *Star) Update(x, y float32) {
	s.fromX = s.toX
	s.fromY = s.toY
	s.toX += (s.toX - x) / 32
	s.toY += (s.toY - y) / 32
	s.brightness += 1
	if 0xff < s.brightness {
		s.brightness = 0xff
	}
	if s.fromX < 0 || screenWidth*scale < s.fromX || s.fromY < 0 || screenHeight*scale < s.fromY {
		s.Init()
	}
}

func (s *Star) Draw(screen *ebiten.Image) {
	c := color.RGBA{
		R: uint8(0xaa * s.brightness / 0xff),
		G: uint8(0xcc * s.brightness / 0xff),
		B: uint8(0xee * s.brightness / 0xff),
		A: 0xff}
	vector.StrokeLine(screen, s.fromX/scale, s.fromY/scale, s.toX/scale, s.toY/scale, 1, c, true)
}

type Game struct {
	stars [starsCount]Star
}

func NewGame() *Game {
	g := &Game{}
	for i := 0; i < starsCount; i++ {
		g.stars[i].Init()
	}
	return g
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	for i := 0; i < starsCount; i++ {
		g.stars[i].Update(float32(x*scale), float32(y*scale))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < starsCount; i++ {
		g.stars[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Stars (Ebitengine Demo)")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}