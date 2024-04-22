package main

import "fmt"

func main() {
	var x uint8 = 255
	x++
	fmt.Println("X contains", x)

	// causes overflow because it is evaluated at runtime
	a := uint8(255) + 1
	fmt.Println("A contains", a)
}
