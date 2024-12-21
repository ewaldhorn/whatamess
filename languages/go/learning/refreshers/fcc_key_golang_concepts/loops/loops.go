package main

import "fmt"

func main() {
	// traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// range
	numbers := []int{2, 4, 6, 8, 10}
	for i, num := range numbers {
		fmt.Printf("We have %3d at position %d\n", num, i)
	}
}
