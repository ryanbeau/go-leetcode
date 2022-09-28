package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodecSerialize(t *testing.T) {
	tests := []struct {
		input    *TreeNode
		expected string
	}{
		{
			input: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			expected: "1,2,,,3,4,,,5,,",
		},
		{
			input: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: "3,9,,,20,15,,,7,,",
		},
		{
			input:    nil,
			expected: "",
		},
	}
	for _, test := range tests {
		codec := Constructor()
		actual := codec.serialize(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

func TestCodecDeserialize(t *testing.T) {
	tests := []struct {
		input    string
		expected *TreeNode
	}{
		{
			input: "1,2,,,3,4,,,5,,",
			expected: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
		{
			input: "-10,9,,,20,15,,,7,,",
			expected: &TreeNode{Val: -10,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			input:    "",
			expected: nil,
		},
	}
	for _, test := range tests {
		codec := Constructor()
		actual := codec.deserialize(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
