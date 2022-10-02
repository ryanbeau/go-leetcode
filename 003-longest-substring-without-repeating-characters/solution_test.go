package solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "abcabcbb",
			expected: 3,
		},
		{
			input:    "bbbbb",
			expected: 1,
		},
		{
			input:    "pwwkew",
			expected: 3,
		},
	}
	for _, test := range tests {
		actual := lengthOfLongestSubstring(test.input)
		assert.Equal(t, test.expected, actual, fmt.Sprintf("lengthOfLongestSubstring(\"%s\")", test.input))
	}
}
