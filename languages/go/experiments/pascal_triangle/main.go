package main

import "fmt"

func main() {
	rows := 20
	coef := 1

	for i := 0; i < rows; i++ {
		for space := 1; space <= rows-i; space++ {
			fmt.Printf("  ")
		}

		for j := 0; j <= i; j++ {
			if j == 0 || i == 0 {
				coef = 1
			} else {
				coef = coef * (i - j + 1) / j
			}

			fmt.Printf("%d   ", coef)
		}

		fmt.Println()
	}
}
