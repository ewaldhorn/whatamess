package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for n := range nums {
		total += nums[n]
	}

	return total
}

func addem(nums ...int) int {
	total := 0

	for _, n := range nums {
		total += n
	}

	return total
}

func main() {
	fmt.Println("The total of 1,2,3 and 4 is:", sum(1, 2, 3, 4))
	fmt.Println("The total of 1,2,3 and 4 is:", addem(1, 2, 3, 4))
	vals1 := []int{1, 2}
	vals2 := []int{3, 4}
	all := append(vals1, vals2...)
	fmt.Println("The total of 1,2,3 and 4 is:", addem(all...))
}
