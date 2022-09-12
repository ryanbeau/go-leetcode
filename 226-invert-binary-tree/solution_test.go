package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected *TreeNode
	}{
		{
			input: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			input:    nil,
			expected: nil,
		},
	}
	for _, test := range tests {
		actual := invertTree(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
