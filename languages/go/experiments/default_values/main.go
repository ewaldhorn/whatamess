package main

import (
	"cmp"
	"fmt"
)

var (
	defaultName = "Stranger"
)

// ----------------------------------------------------------------------------
func getName(really bool) string {
	if really {
		return "Josephine"
	} else {
		return ""
	}
}

// ----------------------------------------------------------------------------
func sayHello(really bool) {
	fmt.Printf("Hello there %s!\n", cmp.Or(getName(really), defaultName))
}

// ----------------------------------------------------------------------------
func main() {
	fmt.Println("Go Default Values:")
	fmt.Println()

	sayHello(true)
	sayHello(false)

	fmt.Println()
}
