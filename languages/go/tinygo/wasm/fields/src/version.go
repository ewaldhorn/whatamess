package main

import (
	"fmt"
	"syscall/js"
)

// ----------------------------------------------------------------------------
const (
	version = "0.0.1f"
	name    = "TinyCanvas Flow Fields Experiment"
)

// ----------------------------------------------------------------------------
func getVersionString() string {
	return fmt.Sprintf("%s v%s", name, version)
}

// ----------------------------------------------------------------------------
func setVersionCallback() {
	js.Global().Set("getVersion", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return getVersionString()
	}))
}
