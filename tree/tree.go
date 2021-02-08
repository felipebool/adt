package tree

import (
	"github.com/felipebool/adt/tree/avl"
	"github.com/felipebool/adt/tree/binary"
)

type Tree interface {
	Add(key int)
	Remove(key int)
	Search(key int) bool
	Empty() bool
	InOrder() []int
	PreOrder() []int
	PostOrder() []int
}

func NewBinaryTree() Tree {
	return &binary.Tree{}
}

func NewAVLTree() Tree {
	return &avl.Tree{}
}
