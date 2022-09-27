package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		input_p  []int
		input_i  []int
		expected *TreeNode
	}{
		{
			input_p: []int{3, 9, 20, 15, 7},
			input_i: []int{9, 3, 15, 20, 7},
			expected: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			input_p:  []int{-1},
			input_i:  []int{-1},
			expected: &TreeNode{Val: -1},
		},
		{
			input_p: []int{1, 2},
			input_i: []int{2, 1},
			expected: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2}},
		},
		{
			input_p: []int{1, 2},
			input_i: []int{1, 2},
			expected: &TreeNode{Val: 1,
				Right: &TreeNode{Val: 2}},
		},
	}
	for _, test := range tests {
		actual := buildTree(test.input_p, test.input_i)
		assert.Equal(t, test.expected, actual)
	}
}
