package list

import (
	"github.com/felipebool/adt/list/singly"
)

type LinkedList interface {
	Front() (int, error)
	PushFront(value int)
	PopFront() (int, error)

	Back() (int, error)
	PushBack(value int)
	PopBack() (int, error)

	PushBefore(value, before int) error
	PushAfter(value, after int) error

	Find(value int) bool
	Erase(value int) error

	Empty() bool
	Size() int

	Reverse() error

	String() string
}

func NewSinglyLinkedList() LinkedList {
	return &singly.SinglyLinkedList{}
}
