package main

import (
	"fmt"
	"unsafe"
)

type IntBool struct {
	a int64
	b bool
	c int32
}

type BoolInt struct {
	a bool
	b int64
	c int32
}

func main() {
	if unsafe.Sizeof(IntBool{}) == unsafe.Sizeof(BoolInt{}) {
		fmt.Println("Same.")
	} else {
		fmt.Printf("Different (%v vs %v)\n", unsafe.Sizeof(IntBool{}), unsafe.Sizeof(BoolInt{}))
	}
}
