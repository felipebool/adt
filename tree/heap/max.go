package heap

type max struct {
	keys []int
	last int
	cap int
}

func (m *max) Insert(key int) error {
	err := ErrFullHeap

	if m.last < m.cap - 1 {
		m.last++
		m.keys[m.last] = key
		m.trickleUp(m.last)
		err = nil
	}	

	return err
}

func (m *max) Remove() (int, error) {
	if m.last == -1 {
		return -1, ErrEmptyHeap
	}

	if m.last == 0 {
		key := m.keys[m.last]
		m.last--

		return key, nil
	}

	key := m.keys[0]
	m.keys[0], m.keys[m.last] = m.keys[m.last], m.keys[0]
	m.last--
	m.trickleDown(0)

	return key, nil
}

func (m *max) Sort() []int {
	sorted := make([]int, 0)

	for m.last >= 0 {
		sorted = append(sorted, m.keys[0])
		m.keys[0] = m.keys[m.last]
		m.last--
		m.trickleDown(0)
	}

	return sorted
}

func (m *max) trickleUp(index int) {
	parent := getParentIndex(index)

	if parent >= 0 && m.keys[index] > m.keys[parent] {
		m.keys[index], m.keys[parent] = m.keys[parent], m.keys[index]
		m.trickleUp(parent)
	}
}

func (m *max) trickleDown(index int) {
	child := m.getGreaterChildIndex(index)

	if child > 0 && m.keys[index] < m.keys[child] {
		m.keys[index], m.keys[child] = m.keys[child], m.keys[index]
		m.trickleDown(child)
		return
	}
}

func (m *max) getGreaterChildIndex(index int) int {
	left, right := getChildrenIndexes(index, m.last)

	if left < 0 {
		if right < 0 {
			return -1
		}

		return right
	}

	if right < 0 {
		return left
	}

	if m.keys[left] > m.keys[right] {
		return left
	}

	return right
}

func NewMaxHeap(capacity int) Heap {
	return &max{
		keys: make([]int, capacity),
		last: -1,
		cap: capacity,
	}
}
