package main

import "fmt"

func main() {
	sports := make([]string, 5)

	sports[0] = "ski"
	sports[1] = "surf"
	sports[2] = "swim"
	sports[3] = "sail"
	sports[4] = "sumo wrestling"

	// reslice the slice, remember this affects the original slice of strings as well
	// third parameter controls capacity, in this case 3-1 = 2, otherwise it would default to capacity of 4
	xs := sports[1:3:3]
	xs[0] = "CHANGED"

	inspectSlice(sports)
	inspectSlice(xs)
}

func inspectSlice(xs []string) {
	fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))

	for i := range xs {
		fmt.Printf("%p \t %v \n", &xs[i], xs[i])
	}
}
