package main

import "fmt"

// ----------------------------------------------------------------------------
func createPlayer1() *Player {
	return &Player{
		name:      "Bob",
		health:    30,
		maxHealth: 30,
		attack:    4,
		defense:   3,
		luck:      3,
	}
}

// ----------------------------------------------------------------------------
func createPlayer2() *Player {
	return &Player{
		name:      "Tax",
		health:    40,
		maxHealth: 40,
		attack:    5,
		defense:   2,
		luck:      1,
	}
}

// ----------------------------------------------------------------------------
func playerAttacks(player1 *Player, opponent *Player) {
	attackStrength := player1.Attack()
	damageDone := opponent.Defend(attackStrength)
	opponent.LoseHealth(damageDone)

	fmt.Printf("%s takes %d damage from the attack!\n", opponent.name, damageDone)
}
