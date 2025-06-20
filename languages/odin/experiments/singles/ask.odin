package main

import "core:fmt"
import "core:os"

main :: proc() {
	buf: []byte

	fmt.print("Enter something: ")
	os.read(os.stdin, buf)
	fmt.println("You typed:", buf)
}
