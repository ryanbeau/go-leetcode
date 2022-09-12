package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 6,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 5},
			},
		},
		Right: &TreeNode{Val: 8,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 9},
		},
	}

	tests := []struct {
		input_r  *TreeNode
		input_p  *TreeNode
		input_q  *TreeNode
		expected *TreeNode
	}{
		{
			input_r:  root,
			input_p:  root.Left,
			input_q:  root.Right,
			expected: root,
		},
		{
			input_r:  root,
			input_p:  root.Left,
			input_q:  root.Left.Right,
			expected: root.Left,
		},
	}
	for _, test := range tests {
		actual := lowestCommonAncestor(root, test.input_p, test.input_q)
		assert.Equal(t, test.expected, actual)
	}
}
