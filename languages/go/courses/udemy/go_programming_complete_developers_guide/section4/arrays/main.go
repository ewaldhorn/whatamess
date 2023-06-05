package main

import "fmt"

type DiscountCard struct {
	id        uint16
	storeCode byte
	rate      float32
}

func main() {
	var (
		myArray      = [3]int{7, 8, 9}
		myOtherArray = [...]int{1, 2, 3, 4, 5}
		myThirdArray = [4]int{1, 2, 3}
		cards        = [...]DiscountCard{{}, {1, 2, 3.0}}
	)

	fmt.Println("myArray contains:", myArray)
	fmt.Println("myOtherArray contains:", myOtherArray)
	fmt.Println("myThirdArray contains:", myThirdArray)

	fmt.Println()
	fmt.Println("cards contains:", cards)

	for i, v := range cards {
		fmt.Printf("Entry %d is %v\n", i, v)
	}
	fmt.Println()

}
