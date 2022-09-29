package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		input_c  []int
		input_t  int
		expected [][]int
	}{
		{
			input_c:  []int{2, 3, 6, 7},
			input_t:  7,
			expected: [][]int{{2, 2, 3}, {7}},
		},
		{
			input_c:  []int{2, 3, 5},
			input_t:  8,
			expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			input_c:  []int{2},
			input_t:  1,
			expected: [][]int{},
		},
		{
			input_c:  []int{4, 2, 8},
			input_t:  8,
			expected: [][]int{{4, 4}, {4, 2, 2}, {2, 2, 2, 2}, {8}},
		},
		{
			input_c:  []int{7, 3, 2},
			input_t:  18,
			expected: [][]int{{7, 7, 2, 2}, {7, 3, 3, 3, 2}, {7, 3, 2, 2, 2, 2}, {3, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 2, 2, 2}, {3, 3, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 2, 2, 2}},
		},
	}
	for _, test := range tests {
		actual := combinationSum(test.input_c, test.input_t)
		assert.Equal(t, test.expected, actual)
	}
}
