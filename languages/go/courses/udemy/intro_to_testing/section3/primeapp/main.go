package main

import "fmt"

func main() {
}

func isPrime(number int) (bool, string) {
	if number < 2 {
		return false, fmt.Sprintf("%d is not a prime number.", number)
	}

	// use mod to figure out if it's a prime
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number, divisible by %d.", number, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number!", number)
}
