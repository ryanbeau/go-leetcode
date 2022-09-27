package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected int
	}{
		{
			input: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 6,
		},
		{
			input: &TreeNode{Val: -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 42,
		},
		{
			input:    &TreeNode{Val: -10},
			expected: -10,
		},
		{
			input: &TreeNode{Val: -10,
				Right: &TreeNode{Val: 10}},
			expected: 10,
		},
	}
	for _, test := range tests {
		actual := maxPathSum(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
