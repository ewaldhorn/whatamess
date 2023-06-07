package main

import (
	"fmt"
	"syscall/js"
)

// RegisterFunction allows the front-end to call Go functions.
func RegisterFunction(funcName string, myfunc func(i []js.Value)) {
	js.Global().Set(funcName, js.NewCallback(myfunc))
}

func main() {
	fmt.Println("All fired up!")
}
