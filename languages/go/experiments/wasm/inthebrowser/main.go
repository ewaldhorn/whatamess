package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func multiply(this js.Value, inputs []js.Value) interface{} {
	fmt.Println("Reading from JSON", inputs[0])
	a, err := strconv.Atoi(inputs[0].String())
	if err != nil {
		fmt.Println("Error", err)
	}
	b, _ := strconv.Atoi(inputs[1].String())
	return a * b
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("multiply", js.FuncOf(multiply))
	<-c
}
