package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		input    [][]byte
		expected int
	}{
		{
			input: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			input: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			input: [][]byte{
				{'1', '0', '1'},
				{'1', '1', '1'},
			},
			expected: 1,
		},
		{
			input: [][]byte{
				{'1', '0', '0', '1', '0', '0', '1'},
			},
			expected: 3,
		},
		{
			input: [][]byte{
				{'1'},
				{'0'},
				{'0'},
				{'1'},
				{'0'},
				{'0'},
				{'1'},
			},
			expected: 3,
		},
	}
	for i, test := range tests {
		actual := numIslands(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("Test#%d", i+1))
	}
}
