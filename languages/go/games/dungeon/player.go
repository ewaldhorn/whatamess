package main

import "math/rand"

// ----------------------------------------------------------------------------
type Player struct {
	name              string
	health, maxHealth int
	attack, defense   int
	luck              int
}

// ----------------------------------------------------------------------------
func (p *Player) LoseHealth(amount int) bool {
	amountLost := amount - rand.Intn(p.luck/2)

	if amountLost < 0 {
		amountLost = 0
	}

	p.health -= amountLost

	if p.health <= 0 {
		p.health = 0
	}

	return p.health > 0
}

// ----------------------------------------------------------------------------
func (p *Player) RegenerateHealth() {}

// ----------------------------------------------------------------------------
func (p *Player) Attack() int {
	return p.luck + rand.Intn(p.attack)
}

// ----------------------------------------------------------------------------
func (p *Player) Defend(incoming int) int {
	damageTaken := incoming - rand.Intn(p.luck)

	if damageTaken < 0 {
		damageTaken = 0
	}

	return damageTaken
}
