package binary

/**
value > left.value
value < right.value
*/
type node struct {
	key    int
	left   *node
	right  *node
	parent *node
}

type Tree struct {
	root *node
}

func (b *Tree) Add(key int) {
	newNode := node{key: key, left: nil, right: nil, parent: nil}

	if b.Empty() {
		b.root = &newNode
		return
	}

	b.add(b.root, nil, &newNode)
}

func (b *Tree) add(current, parent, newNode *node) {
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
		b.add(current.left, current, newNode)
	}

	if current.key < newNode.key {
		b.add(current.right, current, newNode)
	}
}

func (b *Tree) Remove(key int) {
	if b.Empty() {
		return
	}

	// key found in the root node
	if b.root.key == key {
		replacement := inOrderPredecessor(b.root)
		replacement.parent = nil
		replacement.left = b.root.left
		replacement.right = b.root.right
		b.root = replacement

		return
	}

	b.remove(b.root, key)
}

func (b *Tree) remove(t *node, key int) {
	// key not found
	if t == nil {
		return
	}

	if t.key > key {
		b.remove(t.left, key)
		return
	}

	if t.key < key {
		b.remove(t.right, key)
		return
	}

	// key found in a leaf node
	if t.left == nil && t.right == nil {
		if t.parent.key > key {
			t.parent.left = nil
			return
		}
		t.parent.right = nil
		return
	}

	// key found in a non leaf
	replacement := inOrderPredecessor(t)
	replacement.parent = t.parent
	replacement.left = t.left
	replacement.right = t.right

	if t.parent.key > key {
		t.parent.left = replacement
		return
	}

	t.parent.right = replacement
}

func inOrderPredecessor(n *node) *node {
	if n == nil {
		return nil
	}

	var current, previous *node = n.left, nil

	for current != nil {
		previous = current
		current = current.right
	}

	predecessor := previous.parent.right
	previous.parent.right = nil

	return predecessor
}

func (b *Tree) Search(value int) bool {
	if b.Empty() {
		return false
	}

	current := b.root

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

func (b *Tree) Empty() bool {
	return b.root == nil
}

func (b *Tree) InOrder() []int {
	if b.Empty() {
		return []int{}
	}

	return inOrder(b.root)
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

func (b *Tree) PreOrder() []int {
	if b.Empty() {
		return []int{}
	}

	return preOrder(b.root)
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

func (b *Tree) PostOrder() []int {
	if b.Empty() {
		return []int{}
	}

	return postOrder(b.root)
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
