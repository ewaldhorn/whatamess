package main

import "fmt"

func main() {
	fmt.Println("The sum of the digits for 1    is:", sumOfDigits(1))
	fmt.Println("The sum of the digits for 15   is:", sumOfDigits(15))
	fmt.Println("The sum of the digits for 123  is:", sumOfDigits(123))
	fmt.Println("The sum of the digits for 1234 is:", sumOfDigits(1234))
}

func sumOfDigits(src int) int {
	sum := 0
	for src > 0 {
		sum += src % 10
		src /= 10
	}
	return sum
}
