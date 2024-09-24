package main

import (
	"fmt"
	"strings"
)

// ----------------------------------------------------------------------------
func main() {
	justBasicSlice()
	sliceOfSlice()
	removingElements()
}

// ----------------------------------------------------------------------------
func printHeader(title string) {
	fmt.Println("")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
}

// ----------------------------------------------------------------------------
func justBasicSlice() {
	printHeader("Basic Slices:")

	names := []string{"Bob", "Sal", "Rose", "Dennis"}

	fmt.Printf("There are %d names in the list.\n", len(names))

	fmt.Println("Adding \"Taylor\" to the list.")
	names = append(names, "Taylor")

	fmt.Printf("There are now %d names in the list.\n", len(names))

	fmt.Println("Adding \"Ron\", \"Dopey\" and \"Angela\" to the list.")
	names = append(names, "Ron", "Dopey", "Angela")

	fmt.Printf("There are now %d names in the list.\n", len(names))
}

// ----------------------------------------------------------------------------
func sliceOfSlice() {
	printHeader("Slices of Slices:")
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Numbers:", numbers)
	fmt.Printf("There are %d elements in numbers.\n", len(numbers))
	fmt.Println("Slice from position 3:6", numbers[3:6]) // 3,4,5  end is excluded
	fmt.Println("First five elements are (method1)", numbers[0:5])
	fmt.Println("First five elements are (method2)", numbers[:5])
	fmt.Println("Last element is", numbers[len(numbers)-1:])
	fmt.Println("Last 3 elements are", numbers[len(numbers)-3:])

	lastThree := numbers[len(numbers)-3:]

	fmt.Println("LastThree:", lastThree)
	fmt.Printf("LastThree contains %d elements\n", len(lastThree))

	fmt.Println("Adding '10' to the slice")
	lastThree = append(lastThree, 10)

	fmt.Printf("LastThree now contains %d elements (%v)\n", len(lastThree), lastThree)
	fmt.Printf("There are %d elements in numbers (%v).\n", len(numbers), numbers)

	fmt.Println("Changing last item in numbers to 11")
	numbers[len(numbers)-1] = 11
	fmt.Printf("There are %d elements in numbers (%v).\n", len(numbers), numbers)
	fmt.Printf("LastThree now contains %d elements (%v)\n", len(lastThree), lastThree)
}

// ----------------------------------------------------------------------------
func removingElements() {
	printHeader("Removing elements:")
	items := []string{"beans", "pasta", "rice", "chicken", "sweetcorn"}

	fmt.Println("We have    :", items)
	fmt.Println("Let's remove 'chicken'")
	items = append(items[:3], items[4:]...)
	fmt.Println("We now have:", items)
}
