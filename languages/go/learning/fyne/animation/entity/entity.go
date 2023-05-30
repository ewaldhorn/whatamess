package entity

import "math/rand"

type Amoeba struct {
	XPos, YPos, XDelta, YDelta, Radius float32
}

func NewRandomAmoeba() Amoeba {
	return Amoeba{
		XPos:   rand.Float32()*0.5 + 0.1,
		YPos:   rand.Float32()*0.5 + 0.1,
		XDelta: (rand.Float32() - 0.5) * 0.008,
		YDelta: (rand.Float32() - 0.5) * 0.008,
		Radius: rand.Float32()*0.1 + 0.015,
	}
}

func (a *Amoeba) Move() {
	a.XPos += a.XDelta
	if a.XPos <= 0 || a.XPos >= 1 {
		a.XDelta = -a.XDelta
	}

	a.YPos += a.YDelta
	if a.YPos <= 0 || a.YPos >= 1 {
		a.YDelta = -a.YDelta
	}

}
