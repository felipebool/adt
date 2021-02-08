package avl

type Tree struct {
	root *node
}

func (avl *Tree) Add(value int) {
	newNode := node{
		key:    value,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	if avl.Empty() {
		avl.root = &newNode
		return
	}

	avl.add(avl.root, nil, &newNode)
	avl.checkBalance(&newNode)
}

func (avl *Tree) add(current, parent, newNode *node) {
	// place to new node found
	if current == nil {
		newNode.parent = parent
		if parent.key > newNode.key {
			parent.left = newNode
			return
		}

		parent.right = newNode
		return
	}

	// avoid duplicates
	if current.key == newNode.key {
		return
	}

	if current.key > newNode.key {
		avl.add(current.left, current, newNode)
	}

	if current.key < newNode.key {
		avl.add(current.right, current, newNode)
	}

	avl.checkBalance(newNode)
}

func (avl *Tree) checkBalance(n *node) {
	if n.BalanceFactor() > 1 || n.BalanceFactor() < -1 {
		avl.balance(n)
	}

	// it reached root node
	if n.parent == nil {
		return
	}

	avl.checkBalance(n.parent)
}

func (avl *Tree) balance(n *node) {
	if n == nil || (n.left == nil && n.right == nil) {
		return
	}

	if n.BalanceFactor() > 1 {
		// left/left height, left/right height
		llh, lrh := 0, 0

		if n.left == nil {
			return
		}

		if n.left.left != nil {
			llh = n.left.left.Height()
		}

		if n.left.right != nil {
			lrh = n.left.right.Height()
		}

		if llh > lrh {
			n = avl.rightRotate(n)
		} else {
			n = avl.leftRightRotate(n)
		}
	} else {
		// right/left height, right/right height
		rlh, rrh := 0, 0

		if n.right == nil {
			return
		}

		if n.right.left != nil {
			rlh = n.right.left.Height()
		}

		if n.right.right != nil {
			rlh = n.right.right.Height()
		}

		if rlh > rrh {
			n = avl.leftRotate(n)
		} else {
			n = avl.rightLeftRotate(n)
		}
	}

	if n.parent == nil {
		avl.root = n
	}
}

func (avl *Tree) leftRotate(n *node) *node {
	aux := n.right
	n.right = aux.left
	aux.left = n

	return aux
}

func (avl *Tree) rightRotate(n *node) *node {
	aux := n.left
	n.left = aux.right
	aux.right = n

	return aux
}

func (avl *Tree) rightLeftRotate(n *node) *node {
	n.right = avl.rightRotate(n.right)
	return avl.leftRotate(n)
}

func (avl *Tree) leftRightRotate(n *node) *node {
	n.left = avl.leftRotate(n.left)
	return avl.rightRotate(n)
}

func (avl *Tree) Remove(value int) {
	panic("not implemented") // TODO: Implement
}

func (avl Tree) Search(value int) bool {
	if avl.Empty() {
		return false
	}

	current := avl.root

	for current != nil {
		if current.key == value {
			return true
		}

		if current.key > value {
			current = current.left
			continue
		}

		current = current.right
	}

	return false
}

func (avl Tree) Empty() bool {
	return avl.root == nil
}

func (avl Tree) InOrder() []int {
	if avl.Empty() {
		return []int{}
	}

	return inOrder(avl.root)
}

func inOrder(t *node) []int {
	keys := []int{}

	if t == nil {
		return keys
	}

	keys = append(keys, inOrder(t.left)...)
	keys = append(keys, t.key)
	keys = append(keys, inOrder(t.right)...)

	return keys
}

func (avl Tree) PreOrder() []int {
	if avl.Empty() {
		return []int{}
	}

	return preOrder(avl.root)
}

func preOrder(t *node) []int {
	keys := []int{}

	if t == nil {
		return keys
	}

	keys = append(keys, t.key)
	keys = append(keys, preOrder(t.left)...)
	keys = append(keys, preOrder(t.right)...)

	return keys
}

func (avl Tree) PostOrder() []int {
	if avl.Empty() {
		return []int{}
	}

	return postOrder(avl.root)
}

func postOrder(t *node) []int {
	keys := []int{}

	if t == nil {
		return keys
	}

	keys = append(keys, postOrder(t.left)...)
	keys = append(keys, postOrder(t.right)...)
	keys = append(keys, t.key)

	return keys
}
