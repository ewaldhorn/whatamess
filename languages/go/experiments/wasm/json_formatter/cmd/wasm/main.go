package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")

	done := make(chan struct{}, 0)

	// link the functions to JavaScript space
	js.Global().Set("formatJSON", jsonWrapper())

	// keep Go waiting around on a channel
	<-done
}

const PREFIX = ""

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(raw, PREFIX, "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}

// Wraps prettyJson so that it can be called from JavaScript
func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)

		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}

		return pretty
	})

	return jsonFunc
}
