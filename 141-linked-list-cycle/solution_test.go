package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		input_h  []int
		input_p  int
		expected bool
	}{
		{
			input_h:  []int{3, 2, 0, -4},
			input_p:  1,
			expected: true,
		},
		{
			input_h:  []int{1, 2},
			input_p:  0,
			expected: true,
		},
		{
			input_h:  []int{1},
			input_p:  -1,
			expected: false,
		},
	}
	for _, test := range tests {
		var input *ListNode

		//Arrange - create the ListNodes
		if len(test.input_h) > 0 {
			input = &ListNode{
				Val: test.input_h[0],
			}
			current := input
			for i := 1; i < len(test.input_h); i++ {
				current.Next = &ListNode{
					Val: test.input_h[i],
				}
				current = current.Next
			}
			last := current
			if test.input_p >= 0 {
				c := 0
				for i := 0; i < c; i++ {
					current = current.Next
				}
				last.Next = current
			}
		}

		actual := hasCycle(input)

		assert.Equal(t, test.expected, actual)
	}
}
