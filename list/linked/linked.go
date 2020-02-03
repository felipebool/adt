package linked

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Linked struct {
	Root *Node
	Count int
}

func (ll *Linked) Insert(value int) {
	node := Node{
		Value: value,
		Next:  nil,
	}

	if ll.Root == nil {
		ll.Root = &node
		return
	}

	current := ll.Root
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &node
	ll.Count++
}

func (ll *Linked) Delete(value int) {
	current := ll.Root

	if ll.IsEmpty() {
		return
	}

	if ll.Root.Value == value {
		ll.Root = current.Next
		return
	}

	previous := current
	for current.Next != nil {
		if current.Value != value {
			previous = current
			current = current.Next
			continue
		}

		previous.Next = current.Next
		ll.Count--
		return
	}

	if current.Value == value {
		previous.Next = nil
	}
}

func (ll *Linked) String() string {
	current := ll.Root
	result := ""

	if ll.IsEmpty() {
		return result
	}

	for current.Next != nil {
		result += fmt.Sprintf("%d -> ", current.Value)
		current = current.Next
	}

	return result + fmt.Sprintf("%d", current.Value)
}

func (ll *Linked) Search(value int) bool {
	current := ll.Root

	for current.Next != nil {
		if current.Value == value {
			return true
		}
	}

	return current.Value == value
}

func (ll *Linked) IsEmpty() bool {
	return ll.Root == nil
}

