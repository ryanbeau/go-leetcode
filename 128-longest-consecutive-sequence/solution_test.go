package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{100, 4, 200, 1, 3, 2}, expected: 4},
		{input: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, expected: 9},
		{input: []int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3}, expected: 5},
	}
	for _, test := range tests {
		actual := longestConsecutive(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
