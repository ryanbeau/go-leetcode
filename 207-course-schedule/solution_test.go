package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	tests := []struct {
		input_n  int
		input_p  [][]int
		expected bool
	}{
		{
			input_n:  2,
			input_p:  [][]int{{1, 0}},
			expected: true,
		},
		{
			input_n:  2,
			input_p:  [][]int{{1, 0}, {0, 1}},
			expected: false,
		},
	}
	for _, test := range tests {
		actual := canFinish(test.input_n, test.input_p)
		assert.Equal(t, test.expected, actual)
	}
}
