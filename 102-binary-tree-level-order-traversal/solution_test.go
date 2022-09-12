package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected [][]int
	}{
		{
			input: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			input:    &TreeNode{Val: 1},
			expected: [][]int{{1}},
		},
		{
			input:    nil,
			expected: [][]int{},
		},
	}
	for _, test := range tests {
		actual := levelOrder(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
