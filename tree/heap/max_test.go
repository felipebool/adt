package heap_test

import (
	"reflect"
	"testing"

	"github.com/felipebool/adt/tree/heap"
)

func TestMaxInsert(t *testing.T) {
	cases := map[string]struct {
		heapBuilder   func() heap.Heap
		input         int
		expectedError error
	}{
		"no capacity - empty heap": {
			heapBuilder: func() heap.Heap {
				return heap.NewMaxHeap(0)
			},
			input:         1,
			expectedError: heap.ErrFullHeap,
		},
		"no capacity - full heap": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)

				return h
			},
			input:         1,
			expectedError: heap.ErrFullHeap,
		},
		"with capacity - empty heap": {
			heapBuilder: func() heap.Heap {
				return heap.NewMaxHeap(3)
			},
			input:         1,
			expectedError: nil,
		},
		"with capacity - non empty heap": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(1)
				h.Insert(2)

				return h
			},
			input:         1,
			expectedError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			h := tc.heapBuilder()
			if err := h.Insert(tc.input); err != tc.expectedError {
				t.Errorf("got: %+v; expected: %+v", err, tc.expectedError)
			}
		})
	}
}

func TestMaxRemove(t *testing.T) {
	cases := map[string]struct {
		heapBuilder         func() heap.Heap
		expectedRemoveValue int
		expectedRemoveError error
	}{
		"empty heap": {
			heapBuilder: func() heap.Heap {
				return heap.NewMaxHeap(0)
			},
			expectedRemoveValue: -1,
			expectedRemoveError: heap.ErrEmptyHeap,
		},
		"single element": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(1)
				h.Insert(1)

				return h
			},
			expectedRemoveValue: 1,
			expectedRemoveError: nil,
		},
		"multiple elements": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(2)
				h.Insert(1)
				h.Insert(3)

				return h
			},
			expectedRemoveValue: 3,
			expectedRemoveError: nil,
		},
		"multiple elements - sorted": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)

				return h
			},
			expectedRemoveValue: 3,
			expectedRemoveError: nil,
		},
		"multiple elements - reverse sorted": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(3)
				h.Insert(2)
				h.Insert(1)

				return h
			},
			expectedRemoveValue: 3,
			expectedRemoveError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			h := tc.heapBuilder()

			gotValue, gotError := h.Remove()
			if gotValue != tc.expectedRemoveValue {
				t.Errorf("got: %d; expected: %d", gotValue, tc.expectedRemoveValue)
			}

			if gotError != tc.expectedRemoveError {
				t.Errorf("got: %+v; expected: %+v", gotError, tc.expectedRemoveError)
			}
		})
	}
}

func TestMaxSort(t *testing.T) {
	cases := map[string]struct {
		heapBuilder         func() heap.Heap
		expected []int
	}{
		"empty heap - no capacity": {
			heapBuilder: func() heap.Heap {
				return heap.NewMaxHeap(0)
			},
			expected: []int{},
		},
		"empty heap - with capacity": {
			heapBuilder: func() heap.Heap {
				return heap.NewMaxHeap(3)
			},
			expected: []int{},
		},
		"single element - full": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(1)
				h.Insert(1)

				return h
			},
			expected: []int{1},
		},
		"single element - non full": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(1)

				return h
			},
			expected: []int{1},
		},
		"multiple elements - full": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(3)
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)

				return h
			},
			expected: []int{3, 2, 1},
		},
		"multiple elements - non full": {
			heapBuilder: func() heap.Heap {
				h := heap.NewMaxHeap(5)
				h.Insert(1)
				h.Insert(2)
				h.Insert(3)

				return h
			},
			expected: []int{3, 2, 1},
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			h := tc.heapBuilder()

			sorted := h.Sort()
			if !reflect.DeepEqual(tc.expected, sorted) {
				t.Errorf("got: %+v; expected: %+v", sorted, tc.expected)
			}
		})
	}
}
