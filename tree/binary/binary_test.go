package binary_test

import (
	"testing"

	"github.com/felipebool/adt/tree"
)

func TestTree_Add(t *testing.T) {
	cases := map[string]struct {
		input             []int
		expectedInOrder   string
		expectedPreOrder  string
		expectedPostOrder string
	}{
		"empty tree": {
			input:             []int{},
			expectedInOrder:   "()",
			expectedPreOrder:  "()",
			expectedPostOrder: "()",
		},
		"single element": {
			input:             []int{1},
			expectedInOrder:   "(1)",
			expectedPreOrder:  "(1)",
			expectedPostOrder: "(1)",
		},
		"multiple elements": {
			input:             []int{3, 1, 4},
			expectedInOrder:   "(1)(3)(4)",
			expectedPreOrder:  "(3)(1)(4)",
			expectedPostOrder: "(1)(4)(3)",
		},
		"repeated elements": {
			input:             []int{3, 1, 4, 4},
			expectedInOrder:   "(1)(3)(4)",
			expectedPreOrder:  "(3)(1)(4)",
			expectedPostOrder: "(1)(4)(3)",
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			bt := tree.NewBinaryTree()

			for _, element := range tc.input {
				bt.Add(element)
			}

			gotInOrder := bt.InOrder()
			if gotInOrder != tc.expectedInOrder {
				t.Errorf("in order - got: %s; expected: %s", gotInOrder, tc.expectedInOrder)
			}

			gotPreOrder := bt.PreOrder()
			if gotPreOrder != tc.expectedPreOrder {
				t.Errorf("pre order - got: %s; expected: %s", gotPreOrder, tc.expectedPreOrder)
			}

			gotPostOrder := bt.PostOrder()
			if gotPostOrder != tc.expectedPostOrder {
				t.Errorf("post order - got: %s; expected: %s", gotPostOrder, tc.expectedPostOrder)
			}
		})
	}
}

// func TestTree_Remove(t *testing.T) {

// }

func TestTree_Search(t *testing.T) {
	cases := map[string]struct {
		inputNodes []int
		inputValue int
		expected   bool
	}{
		"empty tree": {
			inputNodes: []int{},
			inputValue: 1,
			expected:   false,
		},
		"single element - found": {
			inputNodes: []int{1},
			inputValue: 1,
			expected:   true,
		},
		"single element - not found": {
			inputNodes: []int{1},
			inputValue: 2,
			expected:   false,
		},
		"multiple elements - found": {
			inputNodes: []int{3, 1, 4},
			inputValue: 4,
			expected:   true,
		},
		"multiple elements - not found": {
			inputNodes: []int{3, 1, 4},
			inputValue: 5,
			expected:   false,
		},
		"small search element": {
			inputNodes: []int{3, 1, 4},
			inputValue: 1,
			expected:   true,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			bt := tree.NewBinaryTree()

			for _, element := range tc.inputNodes {
				bt.Add(element)
			}

			got := bt.Search(tc.inputValue)
			if got != tc.expected {
				t.Errorf("got: %v; expected: %v", got, tc.expected)
			}
		})
	}
}
