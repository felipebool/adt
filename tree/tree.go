package tree

import "github.com/felipebool/adt/tree/binary"

type Binary interface {
	Add(value int)
	Remove(value int) error
	Search(value int) bool
	Empty() bool
	InOrder() string
	PreOrder() string
	PostOrder() string
}

func NewBinaryTree() Binary {
	return &binary.Tree{}
}
