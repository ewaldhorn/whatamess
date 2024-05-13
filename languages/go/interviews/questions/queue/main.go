package main

import "errors"

type Queue []interface{}

// Length returns length of the given queue.
func (queue Queue) Length() int {
	return len(queue)
}

// Enqueue appends element to the queue.
func (queue *Queue) Enqueue(element interface{}) {
	*queue = append(*queue, element)
}

// Dequeue dequeues element from the queue.
func (queue *Queue) Dequeue() (element interface{}, err error) {
	q := *queue
	l := q.Length()
	if l == 0 {
		err = errors.New("empty queue, can't dequeue")
		return
	}
	element = q[0]
	*queue = q[1:]
	return
}
