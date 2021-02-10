package heap

import (
	"errors"
	"math"
)

var ErrEmptyHeap = errors.New("empty heap")
var ErrFullHeap = errors.New("full heap")

type Heap interface {
	Insert(key int) error
	Remove() (int, error)
	Sort() []int
}

func getChildrenIndexes(index, last int) (int, int) {
	if index < 0 {
		return -1, -1
	}

	if index == last {
		return -1, -1
	}

	left := 2 * index + 1
	right := 2 * index + 2

	if left > last {
		left = -1
	}

	if right > last {
		right = -1
	}

	return left, right
}

func getParentIndex(index int) int {
	if index < 1 {
		return -1
	}

	return int(math.Floor((float64(index) - 1)/2))
}
