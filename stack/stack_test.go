package stack_test

import (
	"testing"

	"github.com/felipebool/adt/stack"
)

func TestStack_Push(t *testing.T) {
	cases := map[string]struct {
		input    []int
		expected string
	}{
		"empty stack": {
			input:    []int{},
			expected: "",
		},
		"single element": {
			input:    []int{1},
			expected: "1 > nil",
		},
		"multiple elements": {
			input:    []int{1, 2, 3},
			expected: "3 > 2 > 1 > nil",
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			s := stack.NewStack()

			for _, element := range tc.input {
				s.Push(element)
			}

			got := s.String()
			if got != tc.expected {
				t.Errorf("got: %s; expected: %s", got, tc.expected)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	cases := map[string]struct {
		stackBuilder  func() stack.Stack
		expectedValue int
		expectedError error
	}{
		"empty list": {
			stackBuilder: func() stack.Stack {
				return stack.NewStack()
			},
			expectedValue: -1,
			expectedError: stack.ErrEmptyStack,
		},
		"single element": {
			stackBuilder: func() stack.Stack {
				s := stack.NewStack()
				s.Push(1)

				return s
			},
			expectedValue: 1,
			expectedError: nil,
		},
		"multiple elements": {
			stackBuilder: func() stack.Stack {
				s := stack.NewStack()
				s.Push(1)
				s.Push(2)
				s.Push(3)

				return s
			},
			expectedValue: 3,
			expectedError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			s := tc.stackBuilder()
			got, err := s.Pop()
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			if got != tc.expectedValue {
				t.Errorf("got: %d; expected: %d", got, tc.expectedValue)
			}
		})
	}
}
