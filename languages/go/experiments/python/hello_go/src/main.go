package main

import "fmt"
import "C"

// ----------------------------------------------------------------------------
//
//export HelloFromGo
func HelloFromGo() {
	fmt.Println("This is Go, answering the call.")
}

// ----------------------------------------------------------------------------
func main() {}
