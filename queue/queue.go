package queue

import (
	"errors"

	"github.com/felipebool/adt/list"
)

var ErrEmptyQueue = errors.New("empty queue")

type Queue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	String() string
}

type queue struct {
	list list.LinkedList
}

func (q *queue) Enqueue(value int) {
	q.list.PushBack(value)
}

func (q *queue) Dequeue() (int, error) {
	if q.list.Empty() {
		return -1, ErrEmptyQueue
	}

	return q.list.PopFront()
}

func (q *queue) String() string {
	return q.list.String()
}

func NewQueue() Queue {
	return &queue{
		list: list.NewSinglyLinkedList(),
	}
}
