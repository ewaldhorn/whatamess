package main

import "core:fmt"
import "core:slice"

// ----------------------------------------------------------------------------
tmpMap :: proc() {
	tmp := make(map[string]string, context.temp_allocator)
	defer free_all(context.temp_allocator)
}

// ----------------------------------------------------------------------------
maps :: proc() {
	fmt.println("Odin supports maps too:")
	myMap := make(map[string]int)
	defer delete(myMap)

	myMap["Nic"] = 10
	myMap["Sylvi"] = 9
	myMap["Dorris"] = 11
	myMap["Ron"] = 9

	fmt.printf("The score associated with 'Nic' is %d\n", myMap["Nic"])

	if "Casey" in myMap {
		fmt.printf("Casey's score is %d\n", myMap["Casey"])
	} else {
		fmt.println("There's no 'Casey' in the list!")
	}

	fmt.println("Check if a value exists:")
	score, existed := myMap["Astrid"]
	fmt.println(score, existed)

	fmt.println("\nWe also have Go-style 'if ok' style map handling:")
	if val, ok := myMap["Ron"]; ok {
		fmt.println("Ron's score is", val)
	} else {
		fmt.println("Ron wasn't found in the list")
	}
}
