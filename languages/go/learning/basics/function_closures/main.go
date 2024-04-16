package main

import "fmt"

type transformerFunc func(int, int) (int, int)

func main() {
	print("\n")

	myFunc := sumThese()
	for counter := range 6 {
		fmt.Printf("Summing %d gives: %d\n", counter, myFunc(counter))
	}

	print("\n")
	print("\n")

	fmt.Println(swap(1, 2))
	fmt.Println(increment(1, 2))
	fmt.Println(transform(swap, 1, 2))
	fmt.Println(transform(increment, 1, 2))
	fmt.Println(transform_alternative(swap, 1, 2))
	fmt.Println(transform_alternative(increment, 1, 2))

	fmt.Println(multiply(2, 4))
}

func sumThese() func(int) int {
	total := 0
	return func(input int) int {
		total += input
		return total
	}
}

func swap(left, right int) (int, int) {
	return right, left
}

func increment(one, two int) (int, int) {
	return one + 1, two + 1
}

func multiply(one, two int) int {
	return one * two
}

func transform(operation func(int, int) (int, int), one, two int) (int, int) {
	return operation(one, two)
}

func transform_alternative(operation transformerFunc, one, two int) (int, int) {
	return operation(one, two)
}
