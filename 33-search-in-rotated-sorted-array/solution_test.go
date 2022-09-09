package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		inputNums   []int
		inputTarget int
		expected    int
	}{
		{
			inputNums:   []int{4, 5, 6, 7, 0, 1, 2},
			inputTarget: 0,
			expected:    4,
		},
		{
			inputNums:   []int{4, 5, 6, 7, 0, 1, 2},
			inputTarget: 5,
			expected:    1,
		},
		{
			inputNums:   []int{4, 5, 6, 7, 0, 1, 2},
			inputTarget: 3,
			expected:    -1,
		},
		{
			inputNums:   []int{1},
			inputTarget: 0,
			expected:    -1,
		},
		{
			inputNums:   []int{8, 1, 2, 3, 4, 5, 6},
			inputTarget: 1,
			expected:    1,
		},
		{
			inputNums:   []int{8, 1, 2, 3, 4, 5, 6},
			inputTarget: 4,
			expected:    4,
		},
		{
			inputNums:   []int{1, 2, 3, 4, 5, 6, 7},
			inputTarget: 3,
			expected:    2,
		},
		{
			inputNums:   []int{1, 2, 3, 4, 5, 6, 7},
			inputTarget: 6,
			expected:    5,
		},
		{
			inputNums:   []int{9, 0, 4, 8},
			inputTarget: 0,
			expected:    1,
		},
	}
	for _, test := range tests {
		actual := search(test.inputNums, test.inputTarget)
		assert.Equal(t, test.expected, actual)
	}
}
