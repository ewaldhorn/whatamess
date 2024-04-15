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

type VendingMachine struct {
	CurrentState state
}

func (v *VendingMachine) setState(s state) {
	v.CurrentState = s
}

func (v *VendingMachine) insertMoney() {
	v.CurrentState.insertMoney()
}

func (v *VendingMachine) dispenseItem() {
	v.CurrentState.dispenseItem()
}

func main() {
	vendingMachine := &VendingMachine{CurrentState: &NoMoneyState{}}

	vendingMachine.insertMoney()
	vendingMachine.dispenseItem()

	vendingMachine.setState(&HasMoneyState{})

	vendingMachine.insertMoney()
	vendingMachine.dispenseItem()
}
