package singly_test

import (
	"testing"

	"github.com/felipebool/adt/list"
	"github.com/felipebool/adt/list/singly"
)

func TestSinglyLinkedList_FrontBack(t *testing.T) {
	cases := map[string]struct {
		listBuilder        func() list.LinkedList
		expectedFrontValue int
		expectedFrontError error
		expectedBackValue  int
		expectedBackError  error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expectedFrontValue: -1,
			expectedFrontError: singly.ErrEmptyList,
			expectedBackValue:  -1,
			expectedBackError:  singly.ErrEmptyList,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expectedFrontValue: 1,
			expectedFrontError: nil,
			expectedBackValue:  1,
			expectedBackError:  nil,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expectedFrontValue: 1,
			expectedFrontError: nil,
			expectedBackValue:  3,
			expectedBackError:  nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()

			gotFront, errFront := ll.Front()
			if gotFront != tc.expectedFrontValue {
				t.Errorf("front - got: %d; expected: %d", gotFront, tc.expectedFrontValue)
			}

			if errFront != tc.expectedFrontError {
				t.Errorf("front - got: %v; expected: %v", errFront, tc.expectedFrontError)
			}

			gotBack, errBack := ll.Back()
			if gotBack != tc.expectedBackValue {
				t.Errorf("back - got: %d; expected: %d", gotBack, tc.expectedBackValue)
			}

			if errBack != tc.expectedBackError {
				t.Errorf("back - got: %v; expected: %v", errBack, tc.expectedBackError)
			}
		})
	}
}

func TestSinglyLinkedList_PopFront(t *testing.T) {
	cases := map[string]struct {
		listBuilder   func() list.LinkedList
		expectedValue int
		expectedError error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expectedValue: -1,
			expectedError: singly.ErrEmptyList,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushFront(1)

				return ll
			},
			expectedValue: 1,
			expectedError: nil,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushFront(3)
				ll.PushFront(2)
				ll.PushFront(1)

				return ll
			},
			expectedValue: 1,
			expectedError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()

			got, err := ll.PopFront()
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			if got != tc.expectedValue {
				t.Errorf("got: %d; expected: %d", got, tc.expectedValue)
			}
		})
	}
}

func TestSinglyLinkedList_PopBack(t *testing.T) {
	cases := map[string]struct {
		listBuilder   func() list.LinkedList
		expectedValue int
		expectedError error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expectedValue: -1,
			expectedError: singly.ErrEmptyList,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expectedValue: 1,
			expectedError: nil,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expectedValue: 3,
			expectedError: nil,
		},
		"big list": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)

				return ll
			},
			expectedValue: 6,
			expectedError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()

			got, err := ll.PopBack()
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			if got != tc.expectedValue {
				t.Errorf("got: %d; expected: %d", got, tc.expectedValue)
			}
		})
	}
}

func TestSinglyLinkedList_PushBefore(t *testing.T) {
	cases := map[string]struct {
		listBuilder    func() list.LinkedList
		inputValue     int
		inputBefore    int
		expectedString string
		expectedError  error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			inputValue:     1,
			inputBefore:    1,
			expectedString: "",
			expectedError:  singly.ErrEmptyList,
		},
		"single element - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			inputValue:     2,
			inputBefore:    1,
			expectedString: "2 > 1 > nil",
			expectedError:  nil,
		},
		"single element - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			inputValue:     1,
			inputBefore:    2,
			expectedString: "1 > nil",
			expectedError:  singly.ErrValueNotFound,
		},
		"multiple elements - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)

				return ll
			},
			inputValue:     3,
			inputBefore:    4,
			expectedString: "1 > 2 > 3 > 4 > 5 > 6 > nil",
			expectedError:  nil,
		},
		"multiple elements - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)

				return ll
			},
			inputValue:     1,
			inputBefore:    3,
			expectedString: "1 > 2 > 4 > 5 > 6 > nil",
			expectedError:  singly.ErrValueNotFound,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			err := ll.PushBefore(tc.inputValue, tc.inputBefore)
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			got := ll.String()
			if got != tc.expectedString {
				t.Errorf("got: %s; expected: %s", got, tc.expectedString)
			}
		})
	}
}

func TestSinglyLinkedList_PushAfter(t *testing.T) {
	cases := map[string]struct {
		listBuilder    func() list.LinkedList
		inputValue     int
		inputAfter     int
		expectedString string
		expectedError  error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			inputValue:     1,
			inputAfter:     1,
			expectedString: "",
			expectedError:  singly.ErrEmptyList,
		},
		"single element - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			inputValue:     2,
			inputAfter:     1,
			expectedString: "1 > 2 > nil",
			expectedError:  nil,
		},
		"single element - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			inputValue:     1,
			inputAfter:     2,
			expectedString: "1 > nil",
			expectedError:  singly.ErrValueNotFound,
		},
		"multiple elements - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)

				return ll
			},
			inputValue:     3,
			inputAfter:     2,
			expectedString: "1 > 2 > 3 > 4 > 5 > 6 > nil",
			expectedError:  nil,
		},
		"multiple elements - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)

				return ll
			},
			inputValue:     1,
			inputAfter:     3,
			expectedString: "1 > 2 > 4 > 5 > 6 > nil",
			expectedError:  singly.ErrValueNotFound,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			err := ll.PushAfter(tc.inputValue, tc.inputAfter)
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			got := ll.String()
			if got != tc.expectedString {
				t.Errorf("got: %s; expected: %s", got, tc.expectedString)
			}
		})
	}
}

func TestSinglyLinkedList_Find(t *testing.T) {
	cases := map[string]struct {
		listBuilder func() list.LinkedList
		input       int
		expected    bool
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			input:    1,
			expected: false,
		},
		"single element - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			input:    1,
			expected: true,
		},
		"single element - not found ": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			input:    2,
			expected: false,
		},
		"multiple elements - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			input:    2,
			expected: true,
		},
		"multiple elements - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			input:    4,
			expected: false,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()

			got := ll.Find(tc.input)
			if got != tc.expected {
				t.Errorf("got: %v; expected: %v", got, tc.expected)
			}
		})
	}
}

func TestSinglyLinkedList_Erase(t *testing.T) {
	cases := map[string]struct {
		listBuilder func() list.LinkedList
		input       int
		expected    error
	}{
		"empty": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			input:    1,
			expected: singly.ErrEmptyList,
		},
		"single element - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			input:    1,
			expected: nil,
		},
		"single element - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			input:    2,
			expected: singly.ErrValueNotFound,
		},
		"multiple elements - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			input:    1,
			expected: nil,
		},
		"multiple elements - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			input:    4,
			expected: singly.ErrValueNotFound,
		},
		"big list - found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)
				ll.PushBack(7)
				ll.PushBack(8)
				ll.PushBack(9)

				return ll
			},
			input:    5,
			expected: nil,
		},
		"big list - not found": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)
				ll.PushBack(4)
				ll.PushBack(5)
				ll.PushBack(6)
				ll.PushBack(7)
				ll.PushBack(8)
				ll.PushBack(9)

				return ll
			},
			input:    10,
			expected: singly.ErrValueNotFound,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			got := ll.Erase(tc.input)

			if got != tc.expected {
				t.Errorf("got: %v; expected: %v", got, tc.expected)
			}
		})
	}
}

func TestSinglyLinkedList_Empty(t *testing.T) {
	cases := map[string]struct {
		listBuilder func() list.LinkedList
		expected    bool
	}{
		"empty": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expected: true,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expected: false,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expected: false,
		},
		"make it empty": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				if err := ll.Erase(1); err != nil {
					t.Error(err)
				}

				if err := ll.Erase(2); err != nil {
					t.Error(err)
				}

				if err := ll.Erase(3); err != nil {
					t.Error(err)
				}

				return ll
			},
			expected: true,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			got := ll.Empty()

			if got != tc.expected {
				t.Errorf("got: %v; expected: %v", got, tc.expected)
			}
		})
	}
}

func TestSinglyLinkedList_Size(t *testing.T) {
	cases := map[string]struct {
		listBuilder func() list.LinkedList
		expected    int
	}{
		"empty": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expected: 0,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expected: 1,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expected: 3,
		},
		"make it empty": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				if err := ll.Erase(1); err != nil {
					t.Error(err)
				}

				if err := ll.Erase(2); err != nil {
					t.Error(err)
				}

				if err := ll.Erase(3); err != nil {
					t.Error(err)
				}

				return ll
			},
			expected: 0,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			got := ll.Size()

			if got != tc.expected {
				t.Errorf("got: %d; expected: %d", got, tc.expected)
			}
		})
	}
}

func TestSinglyLinkedList_Reverse(t *testing.T) {
	cases := map[string]struct {
		listBuilder    func() list.LinkedList
		expectedString string
		expectedError  error
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expectedString: "",
			expectedError:  singly.ErrEmptyList,
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expectedString: "1 > nil",
			expectedError:  nil,
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expectedString: "3 > 2 > 1 > nil",
			expectedError:  nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()
			err := ll.Reverse()
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			got := ll.String()
			if got != tc.expectedString {
				t.Errorf("got: %s; expected: %s", got, tc.expectedString)
			}
		})
	}
}

func TestSinglyLinkedList_String(t *testing.T) {
	cases := map[string]struct {
		listBuilder func() list.LinkedList
		expected    string
	}{
		"empty list": {
			listBuilder: func() list.LinkedList {
				return list.NewSinglyLinkedList()
			},
			expected: "",
		},
		"single element": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)

				return ll
			},
			expected: "1 > nil",
		},
		"multiple elements": {
			listBuilder: func() list.LinkedList {
				ll := list.NewSinglyLinkedList()
				ll.PushBack(1)
				ll.PushBack(2)
				ll.PushBack(3)

				return ll
			},
			expected: "1 > 2 > 3 > nil",
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			ll := tc.listBuilder()

			got := ll.String()
			if got != tc.expected {
				t.Errorf("got: %s; expected: %s", got, tc.expected)
			}
		})
	}
}
