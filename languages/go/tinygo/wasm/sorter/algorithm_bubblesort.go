package main

type BubbleSorter struct {
	haveSwapped     bool
	position        int
	comparePosition int
	numbers         []int
}

// ----------------------------------------------------------------------------
func NewBubbleSorter() *BubbleSorter {
	sorter := &BubbleSorter{position: 0, comparePosition: 0}

	return sorter
}

// ----------------------------------------------------------------------------
func (b *BubbleSorter) Step() bool {

}
