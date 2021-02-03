package stack

import (
	"errors"

	"github.com/felipebool/adt/list"
)

var ErrEmptyStack = errors.New("empty stack")

type Stack interface {
	Push(value int)
	Pop() (int, error)
	String() string
}

type stack struct {
	list list.LinkedList
}

func (s *stack) Push(value int) {
	s.list.PushFront(value)
}

func (s *stack) Pop() (int, error) {
	if s.list.Empty() {
		return -1, ErrEmptyStack
	}

	return s.list.PopFront()
}

func (s *stack) String() string {
	return s.list.String()
}

func NewStack() Stack {
	return &stack{
		list: list.NewSinglyLinkedList(),
	}
}
