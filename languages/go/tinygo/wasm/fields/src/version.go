package main

import "syscall/js"

// ----------------------------------------------------------------------------
const VERSION = "0.0.1c"
const NAME = "TinyCanvas Flow Fields Experiment"

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
