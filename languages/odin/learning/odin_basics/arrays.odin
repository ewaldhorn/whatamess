package main

import "core:fmt"
import "core:slice"

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

	blanklines(2)
	fmt.println("We also have dynamic arrays in Odin:")
	list := make([dynamic]int)
	defer delete(list)

	append(&list, 3)
	append(&list, 2)
	append(&list, 1)
	enumerateArray(list[:])

	fmt.println("\nBut I can add more:")
	append(&list, 0)
	enumerateArray(list[:])

	fmt.printf("The last value added to the list is %d\n", pop(&list))
	enumerateArray(list[:])

	fmt.println(
		"We can also remove element in the first position.\nNote this will change the order of the array!:",
	)
	unordered_remove(&list, 0)
	enumerateArray(list[:])

	fmt.println("Let's add some values to the list:")
	append(&list, 3)
	append(&list, 5)
	append(&list, 4)
	append(&list, 6)
	enumerateArray(list[:])

	fmt.println("Sort them:")
	slice.sort(list[:])
	enumerateArray(list[:])

}
