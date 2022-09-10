package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 6, 2, 5, 3, 4},
		},
	}
	for _, test := range tests {
		var input *ListNode

		//Arrange - create the ListNodes
		if len(test.input) > 0 {
			input = &ListNode{
				Val: test.input[0],
			}
			current := input
			for i := 1; i < len(test.input); i++ {
				current.Next = &ListNode{
					Val: test.input[i],
				}
				current = current.Next
			}
		}

		reorderList(input)

		//Assert - make an int array of the vals
		var actualVals []int
		current := input
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}

		assert.Equal(t, test.expected, actualVals)
	}
}
