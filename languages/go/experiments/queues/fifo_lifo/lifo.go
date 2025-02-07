package main

// ----------------------------------------------------------------------------
type Lifo struct {
	items []interface{}
}

// ----------------------------------------------------------------------------
func NewLifo() *Lifo {
	return &Lifo{items: make([]interface{}, 0)}
}

// ----------------------------------------------------------------------------
func (l *Lifo) Push(x interface{}) {
	l.items = append([]interface{}{x}, l.items...)
}

// ----------------------------------------------------------------------------
func (l *Lifo) Pop() interface{} {
	if len(l.items) == 0 {
		return nil
	}
	x := l.items[0]
	l.items = l.items[1:]
	return x
}
