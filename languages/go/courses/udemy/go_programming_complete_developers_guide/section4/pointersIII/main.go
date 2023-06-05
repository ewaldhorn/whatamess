package main

import "fmt"

type SecurityTag bool

const (
	Active   = true
	Inactive = false
)

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("Checking out:")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	shirt := Item{"Shirt", Active}
	socks := Item{"Socks", Active}

	basket := []Item{shirt, socks}

	fmt.Println(basket)
	checkout(basket)
	fmt.Println(basket)
}
