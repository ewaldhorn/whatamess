package collection

import (
	"someanimation/entity"
	"sync"
)

type Collection struct {
	amoebas []entity.Amoeba
	mutex   sync.RWMutex
}

func (c *Collection) Value(x, y float32) float32 {
	rl := c.mutex.RLocker()
	rl.Lock()
	defer rl.Unlock()

	var res float32
	for _, a := range c.amoebas {
		dx := x - a.XPos
		dy := y - a.XPos
		res += a.Radius * a.Radius / (dx*dx + dy*dy)
	}

	return res
}

func NewRandomCollection(n int) *Collection {
	amoebas := make([]entity.Amoeba, n)

	for i := 0; i < n; i++ {
		amoebas[i] = entity.NewRandomAmoeba()
	}

	return &Collection{amoebas: amoebas}
}

func (c *Collection) Move() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for i := range c.amoebas {
		c.amoebas[i].Move()
	}
}
