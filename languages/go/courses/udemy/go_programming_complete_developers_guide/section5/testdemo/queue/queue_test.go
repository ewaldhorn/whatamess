package queue

import "testing"

func TestCanCreateQueue(t *testing.T) {
	q := New(10)

	if q.capacity != 10 {
		t.Errorf("Queue capacity is not what I expected!")
	}
}

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 1; i < 4; i++ {
		q.Append(i)
	}

	if len(q.items) != 3 {
		t.Errorf("Incorrect number of items in queue!")
	}
}
