package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	tests := []struct {
		input_b  [][]byte
		input_w  string
		expected bool
	}{
		{
			input_b:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			input_w:  "ABCCED",
			expected: true,
		},
		{
			input_b:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			input_w:  "SEE",
			expected: true,
		},
		{
			input_b:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			input_w:  "ABCB",
			expected: false,
		},
	}
	for _, test := range tests {
		actual := exist(test.input_b, test.input_w)
		assert.Equal(t, test.expected, actual)
	}
}
