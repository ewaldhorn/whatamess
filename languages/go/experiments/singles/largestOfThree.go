package main

import "fmt"

func main() {
	// get input
	fmt.Println("Enter three integers: ")
	var first, second, third int
	fmt.Scanln(&first)
	fmt.Scanln(&second)
	fmt.Scanln(&third)

	fmt.Println()

	if first >= second && first >= third {
		fmt.Print(first)
	} else if second >= first && second >= third {
		fmt.Print(second)
	} else if third >= first && third >= second {
		fmt.Print(third)
	}

	fmt.Println(" is the largest.")
	fmt.Println()
}
