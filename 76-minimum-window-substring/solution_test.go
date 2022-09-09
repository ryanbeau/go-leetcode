package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		input_s  string
		input_t  string
		expected string
	}{
		{input_s: "ADOBECODEBANC", input_t: "ABC", expected: "BANC"},
		{input_s: "a", input_t: "a", expected: "a"},
		{input_s: "a", input_t: "aa", expected: ""},
	}
	for _, test := range tests {
		actual := minWindow(test.input_s, test.input_t)
		assert.Equal(t, test.expected, actual)
	}
}
