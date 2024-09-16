package main

import "github.com/tinne26/kage-desk/display"

func main() {
	display.SetTitle("basic shader example - load from external file")
	display.SetSize(512, 512)
	display.Shader("shader.kage")
}
