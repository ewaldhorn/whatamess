package main

import "fmt"

type Totals struct {
	runningTotal int
}

func addSomething(total *Totals) {
	total.runningTotal += 5
}

// won't work since it receives a COPY of the struct
func addSomethingBut(total Totals) {
	total.runningTotal += 5
}

func main() {
	value := 10

	var valuePtr *int
	fmt.Println("valuePtr is:", valuePtr)

	valuePtr = &value
	fmt.Println("valuePtr address is:", valuePtr)
	fmt.Println("valuePtr value   is:", *valuePtr)

	secondValue := 20
	secondValuePtr := &secondValue

	fmt.Println("secondValuePtr points to the value:", *secondValuePtr)
	secondValue = 21
	fmt.Println("secondValuePtr points to the value:", *secondValuePtr)

	total := Totals{0}

	fmt.Println(total)
	addSomething(&total)
	fmt.Println(total)
	addSomethingBut(total)
	fmt.Println(total)

}
