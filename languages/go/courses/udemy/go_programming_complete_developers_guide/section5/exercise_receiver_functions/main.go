package main

import (
	"fmt"
)

type Player struct {
	health, maxHealth uint
	energy, maxEnergy uint
	name              string
}

func (p *Player) reportStatistics() {
	fmt.Printf("%v has %d/%d health and %d/%d energy.\n", p.name, p.health, p.maxHealth, p.energy, p.maxEnergy)
}

func (p *Player) addHealth(amount uint) {
	p.health += amount

	if p.health > p.maxHealth {
		p.health = p.maxHealth
	}
}

func (p *Player) applyDamage(amount uint) {
	if p.health >= amount {
		p.health -= amount
	} else {
		p.health = 0
	}
}

func (p *Player) addEnergy(amount uint) {
	p.energy += amount

	if p.energy > p.maxEnergy {
		p.energy = p.maxEnergy
	}
}

func (p *Player) useEnergy(amount uint) bool {
	if p.energy >= amount {
		p.energy -= amount
		return true
	}
	return false
}

func main() {
	fmt.Println()

	player1 := Player{health: 50, maxHealth: 50, energy: 5, maxEnergy: 15, name: "Rosh"}
	player1.reportStatistics()
	fmt.Println()

	fmt.Print("Do 76 damage: ")
	player1.applyDamage(76)
	player1.reportStatistics()

	fmt.Print("Add health and energy: ")
	player1.addHealth(10)
	player1.addEnergy(1)
	player1.reportStatistics()

	fmt.Print("Use some energy: ")
	player1.useEnergy(4)
	player1.reportStatistics()
}
