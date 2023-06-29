package main

import (
	"strconv"
	"syscall/js"
)

func multiply(this js.Value, inputs []js.Value) interface{} {
	a, _ := strconv.Atoi(inputs[0].String())
	b, _ := strconv.Atoi(inputs[1].String())
	return a * b
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("multiply", js.FuncOf(multiply))
	<-c
}
