package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			input: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{Val: 4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 4},
				},
			},
			expected: false,
		},
		{
			input: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
	}
	for i, test := range tests {
		actual := isValidBST(test.input)
		assert.Equal(t, test.expected, actual, "test#:%d", i+1)
	}
}
