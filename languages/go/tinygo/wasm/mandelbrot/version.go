package main

import "syscall/js"

// ----------------------------------------------------------------------------
const VERSION = "0.0.1d"
const NAME = "TinyCanvas Mandelbrotter"

// ----------------------------------------------------------------------------
func getVersionString() string {
	return NAME + " v" + VERSION
}

// ----------------------------------------------------------------------------
func setVersionCallback() {
	js.Global().Set("getVersion", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return getVersionString()
	}))
}
