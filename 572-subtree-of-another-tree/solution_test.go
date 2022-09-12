package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		input_a  *TreeNode
		input_b  *TreeNode
		expected bool
	}{
		{
			input_a: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 5},
			},
			input_b: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			input_a: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 4,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2,
						Left: &TreeNode{Val: 0},
					},
				},
				Right: &TreeNode{Val: 5},
			},
			input_b: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			input_b: &TreeNode{Val: 1,
				Left: nil,
				Right: &TreeNode{Val: 1,
					Left: nil,
					Right: &TreeNode{Val: 1,
						Left: nil,
						Right: &TreeNode{Val: 1,
							Left: nil,
							Right: &TreeNode{Val: 1,
								Left: nil,
								Right: &TreeNode{Val: 1,
									Left: &TreeNode{Val: 2},
								},
							},
						},
					},
				},
			},
			input_a: &TreeNode{Val: 1,
				Left: nil,
				Right: &TreeNode{Val: 1,
					Left: nil,
					Right: &TreeNode{Val: 1,
						Left: nil,
						Right: &TreeNode{Val: 1,
							Left: nil,
							Right: &TreeNode{Val: 1,
								Left: nil,
								Right: &TreeNode{Val: 1,
									Left: nil,
									Right: &TreeNode{Val: 1,
										Left: nil,
										Right: &TreeNode{Val: 1,
											Left: nil,
											Right: &TreeNode{Val: 1,
												Left: nil,
												Right: &TreeNode{Val: 1,
													Left: nil,
													Right: &TreeNode{Val: 1,
														Left: &TreeNode{Val: 2},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: true,
		},
	}
	for _, test := range tests {
		actual := isSubtree(test.input_a, test.input_b)
		assert.Equal(t, test.expected, actual)
	}
}
