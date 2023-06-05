package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string
type Noodles string

func (c Chicken) PrepareDish() {
	fmt.Println("Making a chicken dish!")
}

func (s Salad) PrepareDish() {
	fmt.Println("Making a salad!")
}

func (n Noodles) PrepareDish() {
	fmt.Println("Making noodles!")
}

func PrepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")

	for _, dish := range dishes {
		fmt.Println()
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println()
}

func main() {
	PrepareDishes([]Preparer{Chicken("Chicken Wings"), Noodles("Fried Noodles"), Salad("Ceasar Salad")})
}
