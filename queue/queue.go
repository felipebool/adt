package queue

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Queue struct {
	Root *Node
	Count int
}

func (q *Queue) Enqueue(value int) {
	node := Node{
		Value: value,
		Next:  nil,
	}

	if q.IsEmpty() {
		q.Root = &node
		q.Count++
		return
	}

	current := q.Root
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &node
	q.Count++
}

func (q *Queue) Dequeue() *Node {
	if q.IsEmpty() {
		return &Node{}
	}

	if q.Root.Next != nil {
		node := q.Root
		q.Root = q.Root.Next
		q.Count--

		return node
	}

	node := q.Root
	q.Root = nil
	q.Count--

	return node
}

func (q *Queue) IsEmpty() bool {
	return q.Count == 0
}

func (q *Queue) String() string {
	result := ""

	if q.IsEmpty() {
		return result
	}

	current := q.Root
	for current.Next != nil {
		result += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}

	return result + fmt.Sprintf("%d", current.Value)
}

func NewQueue() *Queue {
	return &Queue{}
}