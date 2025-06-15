package main

import "core:fmt"

// ----------------------------------------------------------------------------
modifyNumber :: proc(number: ^int) {
	number^ += 2
}

// ----------------------------------------------------------------------------
pointers :: proc() {
	starter := 0
	fmt.printf("Starter is %d right now, before calling modifyNumber.\n", starter)
	modifyNumber(&starter)
	fmt.printf("Starter is %d now after calling modifyNumber. It was changed in-place.\n", starter)
}

// ----------------------------------------------------------------------------
enumerateArray :: proc(arr: []int) {
	for v, i in arr {
		fmt.printf("Value at index %d is %d\n", i, v)
	}
}

// ----------------------------------------------------------------------------
modifyArray :: proc(arr: []int) {
	arr[2] = 8
}

// ----------------------------------------------------------------------------
arrays :: proc() {
	simpleArray: [5]int

	fmt.println("\nArray with default values:")
	enumerateArray(simpleArray[:])

	fmt.println("\nChanging some values:")
	simpleArray[1] = 10
	simpleArray[3] = 12
	enumerateArray(simpleArray[:])

	fmt.println("\nArray with assigned values:")
	assignedValues := [5]int{2, 4, 6, 8, 10}
	enumerateArray(assignedValues[:])

	fmt.println("\nWe don't need to specify the item count:")
	anyItems := [?]int{1, 2, 3, 3, 2, 1}
	enumerateArray(anyItems[:])

	fmt.println("\nWe can also modify the array:")
	modifyArray(anyItems[:])
	enumerateArray(anyItems[:])
}

// ----------------------------------------------------------------------------
main :: proc() {
	blanklines(2)
	defer blanklines(2)

	defer {
		blanklines(1)
		fmt.println("Tada!")
		fmt.println("And that's all folks!")
	}

	part_one()

	blanklines(2)
	pointers()

	blanklines(2)
	arrays()

}
