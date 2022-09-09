package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			input:    []int{1, 2},
			expected: []int{2, 1},
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

		actual := reverseList(input)

		//Assert - make an int array of the vals
		var actualVals []int
		current := actual
		for current != nil {
			actualVals = append(actualVals, current.Val)
			current = current.Next
		}

		assert.Equal(t, test.expected, actualVals)
	}
}
