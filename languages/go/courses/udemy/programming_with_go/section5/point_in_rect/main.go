package main

type Rect struct {
	topLeftX, topLeftY, width, height int
}

func (r Rect) isInRect(x int, y int) bool {
	return (x >= r.topLeftX && x <= r.topLeftX+r.width) && (y >= r.topLeftY && y <= r.topLeftY+r.height)
}
