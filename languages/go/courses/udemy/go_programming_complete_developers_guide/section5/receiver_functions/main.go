package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) clearSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	testCoord()
	testReceiver()
	testShift()

	lot := ParkingLot{[]Space{{false}, {false}, {false}}}

	fmt.Println("Parking lot state:", lot)
	occupySpace(&lot, 1)
	fmt.Println("Parking lot state:", lot)
	lot.clearSpace(1)
	fmt.Println("Parking lot state:", lot)
}
