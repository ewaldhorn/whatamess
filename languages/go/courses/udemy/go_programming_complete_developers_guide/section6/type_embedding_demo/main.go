package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam Mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %s on %s conveyor.\n", item, item.Convey())
	fmt.Printf("Ship %s via %s.\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t ToxicWaste) String() string {
	return "Toxic Waste"
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func (t *ToxicWaste) Convey() BeltSize {
	return Large
}

func main() {
	fmt.Printf("\nType Embedding Demo\n")

	mail := SpamMail{5000}
	automate(&mail)

	waste := ToxicWaste{9000}
	automate(&waste)
}
