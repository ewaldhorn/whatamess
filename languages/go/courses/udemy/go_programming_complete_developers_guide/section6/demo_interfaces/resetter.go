package main

import "fmt"

type Resetter interface {
	Reset()
}

type Coordinate struct {
	x, y int
}

type Player struct {
	health   int
	position Coordinate
}

func (p *Player) Reset() {
	p.health = 100
	p.position = Coordinate{1, 1}
}

func Reset(r Resetter) {
	r.Reset()
}

func ResetWithPenalty(r Resetter) {
	if player, ok := r.(Player); ok {
		player.health = 40
	} else {
		r.Reset()
	}
}

func Test() {
	player := Player{50, Coordinate{2, 2}}
	fmt.Println(player)

	Reset(&player)
	fmt.Println(player)

	player.health -= 33
	fmt.Println(player)

	player.Reset()
	fmt.Println(player)
}
