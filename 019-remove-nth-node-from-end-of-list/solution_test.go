package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input_v  []int
		input_n  int
		expected []int
	}{
		{
			input_v:  []int{1, 2, 3, 4, 5},
			input_n:  2,
			expected: []int{1, 2, 3, 5},
		},
		{
			input_v:  []int{1},
			input_n:  1,
			expected: []int{},
		},
		{
			input_v:  []int{1, 2},
			input_n:  1,
			expected: []int{1},
		},
		{
			input_v:  []int{1, 2},
			input_n:  2,
			expected: []int{2},
		},
	}
	for _, test := range tests {
		var input *ListNode

		//Arrange - create the ListNodes
		if len(test.input_v) > 0 {
			input = &ListNode{
				Val: test.input_v[0],
			}
			current := input
			for i := 1; i < len(test.input_v); i++ {
				current.Next = &ListNode{
					Val: test.input_v[i],
				}
				current = current.Next
			}
		}

		actual := removeNthFromEnd(input, test.input_n)

		//Assert - make an int array of the vals
		actualVals := []int{}
		current := actual
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}

		assert.Equal(t, test.expected, actualVals)
	}
}
