package singly

import (
	"errors"
	"strconv"
)

var (
	ErrEmptyList     = errors.New("empty list")
	ErrValueNotFound = errors.New("value not found")
)

type node struct {
	value int
	next  *node
}

type SinglyLinkedList struct {
	head *node
	tail *node
	size int
}

// returns the first node value, without removing it from the list
func (sll SinglyLinkedList) Front() (int, error) {
	if sll.Empty() {
		return -1, ErrEmptyList
	}

	return sll.head.value, nil
}

// inserts value in the first position of the list
func (sll *SinglyLinkedList) PushFront(value int) {
	newNode := node{value: value, next: nil}

	if sll.Empty() {
		sll.head = &newNode
		sll.tail = &newNode
		sll.size++

		return
	}

	newNode.next = sll.head
	sll.head = &newNode
	sll.size++
}

// returns the value of the first position of the list, removing the node
func (sll *SinglyLinkedList) PopFront() (int, error) {
	if sll.Empty() {
		return -1, ErrEmptyList
	}

	firstNode := sll.head
	firstNode.next = nil

	sll.head = sll.head.next

	// if head is nil, the list is now empty
	if sll.head == nil {
		sll.tail = nil
	}

	sll.size--
	return firstNode.value, nil
}

// returns the last node value, without removing it from the list
func (sll SinglyLinkedList) Back() (int, error) {
	if sll.Empty() {
		return -1, ErrEmptyList
	}

	return sll.tail.value, nil
}

func (sll *SinglyLinkedList) PushBack(value int) {
	newNode := node{value: value, next: nil}

	if sll.Empty() {
		sll.head = &newNode
		sll.tail = &newNode
		sll.size++

		return
	}

	lastNode := sll.tail
	lastNode.next = &newNode
	sll.tail = &newNode
	sll.size++
}

func (sll *SinglyLinkedList) PopBack() (int, error) {
	if sll.Empty() {
		return -1, ErrEmptyList
	}

	previous, current := sll.head, sll.head

	for current != nil {
		previous = current
		current = current.next
	}

	sll.tail = previous
	current = nil
	sll.size--

	return previous.value, nil
}

func (sll *SinglyLinkedList) PushBefore(value int, before int) error {
	newNode := node{value: value, next: nil}

	if sll.Empty() {
		return ErrEmptyList
	}

	current, previous := sll.head, sll.head
	for current != nil {
		if current.value == before {
			break
		}

		previous = current
		current = current.next
	}

	// it reached the end of the list without finding the element
	if current == nil {
		return ErrValueNotFound
	}

	// value found in the first position
	if previous == current {
		newNode.next = current
		sll.head = &newNode
		sll.size++

		return nil
	}

	previous.next = &newNode
	newNode.next = current
	sll.size++

	return nil
}

func (sll *SinglyLinkedList) PushAfter(value int, after int) error {
	newNode := node{value: value, next: nil}

	if sll.Empty() {
		return ErrEmptyList
	}

	current := sll.head

	for current != nil {
		if current.value == after {
			break
		}

		current = current.next
	}

	// it reached the end of the list without finding the element
	if current == nil {
		return ErrValueNotFound
	}

	newNode.next = current.next
	current.next = &newNode
	sll.size++

	return nil
}

func (sll SinglyLinkedList) Find(value int) bool {
	if sll.Empty() {
		return false
	}

	current := sll.head

	for current != nil {
		if current.value == value {
			return true
		}

		current = current.next
	}

	return false
}

func (sll *SinglyLinkedList) Erase(value int) error {
	if sll.Empty() {
		return ErrEmptyList
	}

	previous, current := sll.head, sll.head
	for current != nil {
		if current.value == value {
			break
		}

		previous = current
		current = current.next
	}

	// it reached the end of the list without finding the element
	if current == nil {
		return ErrValueNotFound
	}

	// list had a single element, it is empty now
	if previous.next == nil && previous.value == value {
		sll.head = nil
		sll.tail = nil
		sll.size--

		return nil
	}

	// element found in the first position
	if previous == current {
		sll.head = current.next
		current.next = nil
		sll.size--

		return nil
	}

	previous.next = current.next
	current.next = nil
	sll.size--

	return nil
}

func (sll SinglyLinkedList) Empty() bool {
	return sll.head == nil
}

func (sll SinglyLinkedList) Size() int {
	return sll.size
}

func (sll *SinglyLinkedList) Reverse() error {
	if sll.Empty() {
		return ErrEmptyList
	}

	previous, current := sll.head, sll.head.next
	previous.next = nil
	sll.tail = previous

	for current != nil {
		nextCurrent := current.next
		nextPrevious := current

		// revert the node pointer
		current.next = previous

		current = nextCurrent
		previous = nextPrevious
	}

	sll.head = previous

	return nil
}

func (sll SinglyLinkedList) String() string {
	str := ""
	sep := ""

	if sll.Empty() {
		return str
	}

	current := sll.head
	for current != nil {
		str += sep + strconv.Itoa(current.value)
		current = current.next
		sep = " > "
	}

	return str + " > nil"
}
