package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		input1   []int
		input2   []int
		expected []int
	}{
		{
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			input1:   []int{},
			input2:   []int{},
			expected: []int{},
		},
		{
			input1:   []int{},
			input2:   []int{0},
			expected: []int{0},
		},
	}
	for _, test := range tests {
		//Arrange - create the ListNodes
		createListNodes := func(n []int) *ListNode {
			var node *ListNode
			if len(n) > 0 {
				node = &ListNode{
					Val: n[0],
				}
				current := node
				for i := 1; i < len(n); i++ {
					current.Next = &ListNode{
						Val: n[i],
					}
					current = current.Next
				}
			}
			return node
		}
		input1, input2 := createListNodes(test.input1), createListNodes(test.input2)

		//Act
		actual := mergeTwoLists(input1, input2)

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
