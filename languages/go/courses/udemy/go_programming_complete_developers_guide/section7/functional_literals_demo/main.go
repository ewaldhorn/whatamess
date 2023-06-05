package main

import "fmt"

func add(ls, rs int) int {
	return ls + rs
}

func compute(ls, rs int, op func(ls, rs int) int) int {
	fmt.Printf("Running a computation with %d and %d\n", ls, rs)
	return op(ls, rs)
}

func main() {
	fmt.Printf("\nFunction Literals Demo.\n")

	fmt.Printf("Adding %d and %d yields %d.\n", 10, 20, add(10, 20))

	fmt.Println()
	fmt.Printf("Calling compute: %d\n", compute(10, 20, add))

	fmt.Println()
	fmt.Printf("Calling compute: %d\n", compute(10, 20, func(l, r int) int { return r - l }))

	fmt.Println()
}
