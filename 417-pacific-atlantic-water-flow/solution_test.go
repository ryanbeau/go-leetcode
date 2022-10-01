package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input:    [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			expected: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			input:    [][]int{{1}},
			expected: [][]int{{0, 0}},
		},
	}
	for _, test := range tests {
		actual := pacificAtlantic(test.input)
		assert.ElementsMatch(t, test.expected, actual, fmt.Sprintf("expected: %v\nactual:   %v", test.expected, actual))
	}
}
