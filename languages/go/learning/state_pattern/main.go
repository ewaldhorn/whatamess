package main

import "fmt"

// Based on https://goexperts.tech/article-details?id=golang-state-pattern-18

type state interface {
	insertMoney()
	dispenseItem()
}

type NoMoneyState struct{}

func (n *NoMoneyState) insertMoney() {
	fmt.Println("Money inserted")
}

func (n *NoMoneyState) dispenseItem() {
	fmt.Println("Please insert money first")
}

type HasMoneyState struct{}

func (h *HasMoneyState) insertMoney() {
	fmt.Println("Money already inserted!")
}

func (h *HasMoneyState) dispenseItem() {
	fmt.Println("Item dispensed")
}

func main() {

}
