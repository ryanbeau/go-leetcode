package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}, {}, {}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input:    [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input:    [][]int{{1, 2, 3}, {}},
			expected: []int{1, 2, 3},
		},
		{
			input:    [][]int{{}, {1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			input:    [][]int{},
			expected: []int{},
		},
		{
			input:    [][]int{{}},
			expected: []int{},
		},
	}
	for _, test := range tests {
		input := []*ListNode{}

		//Arrange - create the ListNodes
		for i := 0; i < len(test.input); i++ {
			if len(test.input[i]) == 0 {
				input = append(input, nil)
				continue
			}

			current := &ListNode{
				Val: test.input[i][0],
			}
			input = append(input, current)

			for j := 1; j < len(test.input[i]); j++ {
				next := &ListNode{
					Val: test.input[i][j],
				}
				current.Next = next
				current = next
			}
		}

		actual := mergeKLists(input)

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
