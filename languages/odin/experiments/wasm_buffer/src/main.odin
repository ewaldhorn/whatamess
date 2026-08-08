package main

import "core:fmt"

main :: proc() {
	fmt.printfln("Testing")
}

// ------------------------------------------------------------------------------------------------
@(export)
bootup :: proc "c" () {}
