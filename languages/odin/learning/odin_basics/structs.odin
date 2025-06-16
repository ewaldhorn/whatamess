package main

import "core:fmt"

// ----------------------------------------------------------------------------
Person :: struct {
	name: string,
	age:  int,
}

// ----------------------------------------------------------------------------
change_name :: proc(some_person: ^Person) {
	some_person.name = "John"
}

// ----------------------------------------------------------------------------
structs :: proc() {
	fmt.println("Looking at how Odin supports structs.")
	bob := Person{"Bob", 31}
	fmt.println(bob)
	change_name(&bob)
	fmt.println(bob)
}
