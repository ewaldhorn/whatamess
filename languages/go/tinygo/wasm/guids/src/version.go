package main

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
const (
	version = "0.0.1a"
	name    = "TinyGo GUID Generator"
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
