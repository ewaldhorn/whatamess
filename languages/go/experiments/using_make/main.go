package main

import "fmt"

func main() {
	// Using make() to create a slice with a specified length and capacity
	s1 := make([]int, 10, 15)

	// Initializing the elements
	for i := 0; i < 10; i++ {
		s1[i] = i + 1
	}

	fmt.Println(s1)

	// Using short-hand to do the same
	s2 := []int{}

	for i := 0; i < 10; i++ {
		s2 = append(s2, i+1)
	}

	fmt.Println(s2)
}
