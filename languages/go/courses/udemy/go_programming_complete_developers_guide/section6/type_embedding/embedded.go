package main

import "fmt"

type Whisperer interface {
	Whisper() string
}

type Yeller interface {
	Yell() string
}

type Talker interface {
	Whisperer
	Yeller
}

func makeNoise(t Talker) {
	t.Whisper()
	t.Yell()
}

type Account struct {
	accountId int
	balance   int
	name      string
}

type ManagerAccount struct {
	Account
	title string
}

func testManagerAccount() {
	managerAccount := ManagerAccount{title: "Sir", Account: Account{accountId: 10, balance: 100, name: "Bob"}}
	fmt.Println(managerAccount)
}

func reportManagerAccount(mgr ManagerAccount) {
	fmt.Println("The manager is called", mgr.name)
}
