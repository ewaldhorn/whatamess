package main

import (
	"fmt"
)

type Stuff struct {
	values []int
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) || index < 0 {
		return 0, fmt.Errorf("no value at index %d", index)
	} else {
		return s.values[index], nil
	}
}

func main() {
	fmt.Println()
	fmt.Println("Go Errors Demo")

	stuff := Stuff{[]int{1, 2, 3, 4, 5, 6}}

	val, err := stuff.Get(3)

	if err != nil {
		fmt.Println("The error is:", err)
	} else {
		fmt.Println("The value at index 3 is:", val)
	}

	val, err = stuff.Get(12)

	if err != nil {
		fmt.Println("The error is:", err)
	} else {
		fmt.Println("The value at index 12 is:", val)
	}

	fmt.Println()
}
