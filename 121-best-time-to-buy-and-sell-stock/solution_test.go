package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, test := range tests {
		actual := maxProfit(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
