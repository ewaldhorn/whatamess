package main

import "fmt"

func IsEqual[T comparable](a, b T) bool {
	return a == b
}

func testComparable() {
	fmt.Println(IsEqual(2, 2))
	fmt.Println(IsEqual("foo", "bar"))
	fmt.Println(IsEqual('a', 'b'))
	fmt.Println(IsEqual[uint8](9, 10))
	fmt.Println(IsEqual[uint8](10, 10))
}
