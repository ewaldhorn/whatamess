package main

import "fmt"

func main() {
	myArray := [...]string{"one", "two", "three", "four", "five", "six"}

	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println("mySlice contains:", mySlice)

	printSlice(mySlice)
	fmt.Println()
	fmt.Println("Append to the slice:")
	mySlice = append(mySlice, 6)
	printSlice(mySlice)

	fmt.Println()
	printStringSlice(myArray[:])

	fmt.Println()
	fmt.Println("The first three elements are:")
	firstThree := myArray[:3]
	printStringSlice(firstThree)

	fmt.Println()
	fmt.Println("The last two elements are:")
	lastTwo := myArray[len(myArray)-2:]
	printStringSlice(lastTwo)

	fmt.Println()
	fmt.Println("The middle two elements are:")
	middleTwo := myArray[2:4]
	printStringSlice(middleTwo)

	fmt.Println()
	fmt.Println("Now combine firstThree, lastTwo")
	combined := append(firstThree, lastTwo...)
	fmt.Println(combined)

	fmt.Println()
	fmt.Println("And myArray is now:", myArray)
	fmt.Println("Oh no!!! The underlying array is affected by the slices!")

	mySecondSlice := []int{7, 8, 9, 10}
	printSlice(mySecondSlice)

	fmt.Println()
	fmt.Println("Now combine some slices:")
	combinedSlice := append(mySlice, mySecondSlice...)
	printSlice(combinedSlice)

	board := [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	fmt.Println(board)

}

func printStringSlice(slice []string) {
	for i, v := range slice {
		fmt.Printf("At index %d we have value %s\n", i, v)
	}
}

func printSlice(slice []int) {
	for i, v := range slice {
		fmt.Printf("At index %d we have value %d\n", i, v)
	}
}
