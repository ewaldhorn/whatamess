package main

import "syscall/js"

func main() {
	document := js.Global().Get("document")
	document.Call("getElementById", "loading").Get("style").Call("setProperty", "display", "none")
	document.Call("getElementById", "calc").Get("style").Call("setProperty", "display", "block")
}
