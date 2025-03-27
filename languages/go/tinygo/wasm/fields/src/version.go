package main

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
const (
	version = "0.0.1g"
	name    = "TinyCanvas Flow Fields Experiment"
)

// ----------------------------------------------------------------------------
func getVersionString() string {
	return name + " v" + version
}

// ----------------------------------------------------------------------------
func setVersionCallback() {
	js.Global().Set("getVersion", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return getVersionString()
	}))
}
