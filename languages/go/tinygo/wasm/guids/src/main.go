package main

import (
	"fmt"
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

// ------------------------------------------ Externally provided by JavaScript

//export bootstrap
func bootstrap()

//export animateCanvasOne
func animateCanvasOne()

// -------------------------------------------------------------------- GLOBALS
var outputContainer dom.HTMLElement

// ----------------------------------------------------------------------------
func main() {
	setCallbacks()
	setupKeyListeners()
	initUX()
	bootstrap()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func initUX() {
	dom.Hide("loading")
	dom.Show("controls")
	dom.Show("information")
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setRefreshCallback()
}

// ----------------------------------------------------------------------------
func setupKeyListeners() {
	js.Global().Set("onkeydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		keyPressed := args[0].Get("code").String()

		isCtrlPressed := args[0].Get("ctrlKey").Bool()
		isMetaPressed := args[0].Get("metaKey").Bool()

		switch keyPressed {
		case "KeyR":
			if !isCtrlPressed && !isMetaPressed {
				fmt.Println("Random")
			}
		case "KeyG":
			fmt.Println("Uh huh")
		}

		return nil
	}))
}

// ----------------------------------------------------------------------------
func setRefreshCallback() {
	dom.AddEventListener("refreshButton", "click", func() {
		if outputDiv != nil {
			effect.randomise(true)
		}
	})
}
