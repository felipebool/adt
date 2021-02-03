package binary

import (
	"errors"
	"strconv"
)

var (
	ErrValueNotFound = errors.New("element not found")
	ErrEmptyTree     = errors.New("empty tree")
)

/**
value >= left.value
value < right.value
*/
type node struct {
	value int
	left  *node
	right *node
}

type Tree struct {
	root *node
}

func (b *Tree) Add(value int) {
	newNode := node{value: value, left: nil, right: nil}

	if b.Empty() {
		b.root = &newNode
		return
	}

	current, previous := b.root, b.root

	for current != nil {
		previous = current

		if current.value == value {
			return
		}

		if current.value > value {
			current = current.left
			continue
		}

		current = current.right
	}

	if previous.value >= value {
		previous.left = &newNode
		return
	}

	previous.right = &newNode
}

func (b *Tree) Remove(value int) error {
	// @TODO implement me
	return nil
}

func (b *Tree) Search(value int) bool {
	if b.Empty() {
		return false
	}

	current := b.root

	for current != nil {
		if current.value == value {
			return true
		}

		if current.value > value {
			current = current.left
			continue
		}

		current = current.right
	}

	return false
}

func (b *Tree) Empty() bool {
	return b.root == nil
}

func (b *Tree) InOrder() string {
	if b.root == nil {
		return "()"
	}

	return inOrder(b.root)
}

func inOrder(t *node) string {
	str := ""

	if t == nil {
		return str
	}

	str += inOrder(t.left)
	str += "(" + strconv.Itoa(t.value) + ")"
	str += inOrder(t.right)

	return str
}

func (b *Tree) PreOrder() string {
	if b.root == nil {
		return "()"
	}

	return preOrder(b.root)
}

func preOrder(t *node) string {
	str := ""

	if t == nil {
		return str
	}

	str += "(" + strconv.Itoa(t.value) + ")"
	str += preOrder(t.left)
	str += preOrder(t.right)

	return str
}

func (b *Tree) PostOrder() string {
	if b.root == nil {
		return "()"
	}

	return postOrder(b.root)
}

func postOrder(t *node) string {
	str := ""

	if t == nil {
		return str
	}

	str += postOrder(t.left)
	str += postOrder(t.right)
	str += "(" + strconv.Itoa(t.value) + ")"

	return str
}
