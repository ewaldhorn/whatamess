package main

import "fmt"

type Integers32 interface {
	~int32 | ~uint32
}

func SumNumbers[T Integers32](arr []T) T {
	var sum T

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

type MyInt int32

func testSums() {
	nums1 := []int32{1, 2, 3, 4, 5}
	nums2 := []uint32{1, 2, 3, 4, 5}
	nums3 := []MyInt{1, 2, 3, 4, 5}

	total1 := SumNumbers(nums1)
	total2 := SumNumbers(nums2)
	total3 := SumNumbers(nums3)

	fmt.Printf("\n\nTotals are:\n%d\n%d\n%d\n", total1, total2, total3)
}
