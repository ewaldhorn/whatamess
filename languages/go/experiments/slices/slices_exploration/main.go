package main

import "fmt"

func main() {
	fmt.Println("Playing with slices")
	watSlice()
	makeSlices()
	sliceSize()
}

func watSlice() {
	s1 := []int{1, 2, 3}
	fmt.Println("S1:", s1)
	head := s1[:1] // {1}
	tail := s1[1:] // {2, 3}
	fmt.Println("Head is", head)
	fmt.Println("Tail is", tail)

	head = append(head, 4)
	fmt.Println("Head is now:", head)
	fmt.Println("Tail is now:", tail)
	fmt.Println("S1 is now  :", s1)
}

func makeArrays() {
	var a1 [1]int                    // declared but not initialised
	var a2 = [3]int{1, 2, 3}         // declare and initialise
	var a3 = [...]int{1, 2, 3, 4, 5} // declare but you count it for me thanks
	a4 := [...]int{1, 2}             // declare and initialise with shorthand

	fmt.Println(a1, a2, a3, a4)
}

func makeSlices() {
	// useful to remember slices are views sitting on top of an underlying supporting array of the same type
	var s1 []int             // uninitialised slice, value is nil
	var s2 = []int{1, 2, 3}  // initialised slice
	s3 := []int{1, 2, 3}     // shorthand initialisation
	s4 := make([]int, 0, 10) // pre-allocate capacity of 10, handy for performance tuning

	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Printf("s2 len: %d, cap: %d\n", len(s2), cap(s2))
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))
	fmt.Printf("s4 len: %d, cap: %d\n", len(s4), cap(s4))
}

func sliceAndDice() {
	// s := a[start:end] includes start, excludes end, indexing starts at 0
	a := [...]int{1, 2, 3} // array [3]int {1,2,3}
	s1 := a[:2]            // slice []int{1,2} (from 0..1)
	s2 := a[1:]            // slice []int{2,3} (from 1..end)
	s3 := a[1:2]           // slice []int{2} (from 1..1)

	fmt.Println(s1, s2, s3)
}

func sliceSize() {
	fmt.Println("\nSlice growth")
	s1 := make([]int, 0, 5)
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add five numbers.")
	for i := range 5 {
		s1 = append(s1, i+1)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another five numbers.")
	for i := range 5 {
		s1 = append(s1, i+6)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another three numbers. (Forces underlying array capacity to grow)")
	for i := range 3 {
		s1 = append(s1, i+11)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("\nNow we add another 11 numbers. (Forces underlying array capacity to grow again)")
	for i := range 11 {
		s1 = append(s1, i+14)
	}
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Println("s1:", s1)

	fmt.Println("Seems like array size doubles...")
}
