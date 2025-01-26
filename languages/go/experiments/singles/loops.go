package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	standardLoop()
	infiniteLoop()
	whileLoop()

	fmt.Println()
}

func standardLoop() {
	fmt.Print("Standard loop: ")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%3d ", i)
	}

	fmt.Println()
}

func infiniteLoop() {
	fmt.Print("Infinite loop: ")

	i := 1

	for {
		fmt.Printf("%3d ", i)
		i++

		if i > 10 {
			break
		}
	}

	fmt.Println()
}

func whileLoop() {
	fmt.Print("While loop   : ")

	i := 1

	for i <= 10 {
		fmt.Printf("%3d ", i)
		i++
	}

	fmt.Println()
}
