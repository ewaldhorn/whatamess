package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) String() string {
	return fmt.Sprintf("Car: %v", string(c))
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) String() string {
	return fmt.Sprintf("Truck: %v", string(t))
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Println("Use a small lift for", p)
	case StandardLift:
		fmt.Println("Use a standard lift for", p)
	default:
		fmt.Println("Use a large lift for", p)
	}
}

func main() {
	c := Car("Toyota")
	sendToLift(c)

	sendToLift(Truck("Mitsubishi"))
	sendToLift(Motorcycle("Moto Mia"))
}
