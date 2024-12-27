package main

import (
	"smallgo/src/dom"
	"strconv"
)

// ----------------------------------------------------------------------------
func main() {
	// handle loading status update
	dom.Hide("loading")
	dom.Show("calc")

	// now get the calculator ready for action
	dom.SetValue("first-number", "value", "")
	dom.SetValue("second-number", "value", "0")
	dom.SetValue("result", "value", "0")

	// set focus on the first number
	dom.SetFocus("first-number")

	// add listeners for value changes on either input field
	dom.AddEventListener("first-number", "input", performCalculation)
	dom.AddEventListener("second-number", "input", performCalculation)

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch
}

// ----------------------------------------------------------------------------
func performCalculation() {
	firstNumber, firstNumberErr := strconv.Atoi(dom.GetString("first-number", "value"))
	secondNumber, secondNumberErr := strconv.Atoi(dom.GetString("second-number", "value"))

	if firstNumberErr == nil && secondNumberErr == nil {
		dom.SetValue("result", "value", strconv.Itoa(firstNumber+secondNumber))
		dom.RemoveClass("result", "error")
	} else {
		dom.SetValue("result", "value", "ERR")
		dom.AddClass("result", "error")
	}
}
