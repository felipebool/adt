package queue_test

import (
	"testing"

	"github.com/felipebool/adt/queue"
)

func TestQueue_Enqueue(t *testing.T) {
	cases := map[string]struct {
		input    []int
		expected string
	}{
		"empty queue": {
			input:    []int{},
			expected: "",
		},
		"single element": {
			input:    []int{1},
			expected: "1 > nil",
		},
		"multiple elements": {
			input:    []int{1, 2, 3},
			expected: "1 > 2 > 3 > nil",
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			q := queue.NewQueue()
			for _, element := range tc.input {
				q.Enqueue(element)
			}

			got := q.String()
			if got != tc.expected {
				t.Errorf("got: %s; expected: %s", got, tc.expected)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	cases := map[string]struct {
		queueBuilder  func() queue.Queue
		expectedValue int
		expectedError error
	}{
		"empty queue": {
			queueBuilder: func() queue.Queue {
				return queue.NewQueue()
			},
			expectedValue: -1,
			expectedError: queue.ErrEmptyQueue,
		},
		"single element": {
			queueBuilder: func() queue.Queue {
				q := queue.NewQueue()
				q.Enqueue(1)

				return q
			},
			expectedValue: 1,
			expectedError: nil,
		},
		"multiple elements": {
			queueBuilder: func() queue.Queue {
				q := queue.NewQueue()
				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)

				return q
			},
			expectedValue: 1,
			expectedError: nil,
		},
	}

	for label := range cases {
		tc := cases[label]

		t.Run(label, func(t *testing.T) {
			t.Parallel()

			q := tc.queueBuilder()

			got, err := q.Dequeue()
			if err != tc.expectedError {
				t.Errorf("got: %v; expected: %v", err, tc.expectedError)
			}

			if got != tc.expectedValue {
				t.Errorf("got: %d; expected: %d", got, tc.expectedValue)
			}
		})
	}
}
