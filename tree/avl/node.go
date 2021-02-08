package avl

type node struct {
	key    int
	left   *node
	right  *node
	parent *node
}

func (n *node) Height() int {
	return height(n, 1)
}

func height(n *node, initial int) int {
	if n == nil {
		return 0
	}

	left := height(n.left, initial + 1)
	right := height(n.right, initial + 1)

	if left >= right {
		return left
	}

	return right
}

func (n *node) BalanceFactor() int {
	return n.left.Height() - n.left.Height()
}
