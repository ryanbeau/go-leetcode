package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		input_a  *TreeNode
		input_b  *TreeNode
		expected bool
	}{
		{
			input_a: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			input_b: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			input_a: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2},
			},
			input_b: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		actual := isSameTree(test.input_a, test.input_b)
		assert.Equal(t, test.expected, actual)
	}
}
