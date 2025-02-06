package main

// ----------------------------------------------------------------------------
type Fifo struct {
	items []interface{}
}

// ----------------------------------------------------------------------------
func NewFifo() *Fifo {
	return &Fifo{items: make([]interface{}, 0)}
}

// ----------------------------------------------------------------------------
func (f *Fifo) Push(x interface{}) {
	f.items = append(f.items, x)
}

// ----------------------------------------------------------------------------
func (f *Fifo) Pop() interface{} {
	if len(f.items) == 0 {
		return nil
	}
	x := f.items[0]
	f.items = f.items[1:]
	return x
}
