package main

import "fmt"

func main() {

	// count in threes
	for i := 3; i < 100; i += 3 {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	i := 0

	for i <= 10 {
		i++
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	i = 20

	for {
		if i < 1 {
			break
		}
		fmt.Printf("%d ", i)
		i--
	}
	fmt.Println()

	for j := 1; j <= 50; j++ {
		if j%2 != 0 {
			continue
		}

		fmt.Printf("%d, ", j)
	}
	fmt.Println()
}
