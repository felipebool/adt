package stack

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	Root  *Node
	Count int
}

func (s *Stack) Push(value int) {
	node := Node{
		Value: value,
		Next:  nil,
	}

	if s.Count == 0 {
		s.Root = &node
		s.Count++
		return
	}

	current := s.Root
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &node
	s.Count++
}

func (s *Stack) Pop() *Node {
	if s.Count == 0 {
		return &Node{}
	}

	current := s.Root

	if s.Count == 1 {
		s.Count--
		s.Root = nil

		return current
	}

	previous := current
	for current.Next != nil {
		previous = current
		current = current.Next
	}

	previous.Next = nil

	s.Count--
	return current
}

func (s *Stack) String() string {
	result := ""

	if s.Root == nil {
		return result
	}

	current := s.Root
	for current.Next != nil {
		result += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}

	return result + fmt.Sprintf("%d", current.Value)
}

func NewStack() *Stack {
	return &Stack{}
}
