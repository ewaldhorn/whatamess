package main

import (
	"fmt"
	"syscall/js"
)

// ----------------------------------------------------------------------------
func main() {
	js.Global().Set("formatJSON", jsonWrapper())

	<-make(chan struct{}) // prevent early exit
}

// ----------------------------------------------------------------------------
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettifyJSON(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}
