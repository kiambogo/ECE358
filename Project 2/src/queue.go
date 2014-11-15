package main

import(
	"container/list"
)

type Queue struct {
	buffer *list.List
}

// "Push" into the queue
func (q *Queue) enqueue(packet *Packet) {
	q.buffer.PushBack(packet)
}

// "Pop" off the queue
func (q *Queue) dequeue() *Packet {
	e := q.buffer.Front().Value.(*Packet)
	q.buffer.Remove(q.buffer.Front())
	return e
}

// "Peek" at the top of the queue
func (q *Queue) peek() *Packet {
	return q.buffer.Front().Value.(*Packet)
}
