package main

import "core:fmt"

// ----------------------------------------------------------------------------
Person :: struct {
	name: string,
	age:  int,
}
// ----------------------------------------------------------------------------
structs :: proc() {
	fmt.println("Looking at how Odin supports structs.")
	bob := Person{"Bob", 31}
	fmt.println(bob)
}
