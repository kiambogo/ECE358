package main

import(
	"container/list"
)

type Queue struct {
	buffer *list.List
}

func (q Queue) enqueue(packet Packet) {
	q.buffer.PushBack(packet)
}

func (q Queue) dequeue() (packet Packet) {
	e := Packet(q.buffer.Front().Value.(Packet))
	q.buffer.Remove(q.buffer.Front())
	return e
}

func (q Queue) peek() (Packet) {
	return q.buffer.Front().Value.(Packet)
}
