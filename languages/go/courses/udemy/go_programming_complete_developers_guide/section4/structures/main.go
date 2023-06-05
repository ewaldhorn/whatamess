package main

import "fmt"

type Sample struct {
	name string
	age  int
}

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	a := Sample{
		name: "Bob",
		age:  22,
	}

	b := Sample{"Nolan", 19}

	c := Sample{}

	fmt.Println(a, b, c)

	who, old := a.name, a.age

	fmt.Printf("%v is %d years old.\n", who, old)

	casey := Passenger{"Casey Ribeck", 1, false}
	fmt.Println(casey)

	var (
		bill   = Passenger{"Bill", 2, false}
		connie = Passenger{"Connie", 3, true}
		ted    = Passenger{"Ted", 4, true}
	)

	fmt.Println(bill, connie, ted)

	bus := Bus{casey}

	fmt.Println(bus)
	fmt.Printf("%v is in the front passenger seat.", bus.FrontSeat.Name)
}
