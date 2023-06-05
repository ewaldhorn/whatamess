package main

import "fmt"

func main() {
	fmt.Println()
	fmt.Println("All about the maps")

	myFirstMap := make(map[string]int)
	myFirstMap["blue"] = 1
	myFirstMap["green"] = 2
	myFirstMap["red"] = 3
	fmt.Println(myFirstMap)

	mySecondMap := map[string]int{"blue": 1, "green": 2, "red": 3}
	fmt.Println(mySecondMap)

	myFirstMap["orange"] = 4
	fmt.Println(myFirstMap)
	delete(myFirstMap, "orange")
	fmt.Println(myFirstMap)

	val, has := myFirstMap["gray"]
	if has {
		fmt.Println(val)
	} else {
		fmt.Println("Alas, no.")
	}

	fmt.Println("In the map:", myFirstMap)
	for key, val := range myFirstMap {
		fmt.Printf("We have the value of %d for the key %q\n", val, key)
	}
}
