package main

import (
	"github.com/gowebapi/webapi/core/js"
)

func main() {
	waitChannel := make(chan struct{})

	createCanvas()
	js.Global().Call("requestAnimationFrame", js.FuncOf(drawFrame))

	<-waitChannel
}
