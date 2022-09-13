package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		input_r  *TreeNode
		input_k  int
		expected int
	}{
		{
			input_r: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			input_k:  1,
			expected: 1,
		},
		{
			input_r: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			input_k:  3,
			expected: 3,
		},
	}
	for i, test := range tests {
		actual := kthSmallest(test.input_r, test.input_k)
		assert.Equal(t, test.expected, actual, "test#:%d", i+1)
	}
}
