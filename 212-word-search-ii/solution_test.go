package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		input_b  [][]byte
		input_w  []string
		expected []string
	}{
		{
			input_b:  [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			input_w:  []string{"oath", "pea", "eat", "rain"},
			expected: []string{"eat", "oath"},
		},
		{
			input_b:  [][]byte{{'a', 'b'}, {'c', 'd'}},
			input_w:  []string{"abcd"},
			expected: []string{},
		},
		{
			input_b:  [][]byte{{'o', 'a', 'b', 'n'}, {'o', 't', 'a', 'e'}, {'a', 'h', 'k', 'r'}, {'a', 'f', 'l', 'v'}},
			input_w:  []string{"oa", "oaa"},
			expected: []string{"oa", "oaa"},
		},
	}
	for _, test := range tests {
		actual := findWords(test.input_b, test.input_w)
		assert.ElementsMatch(t, test.expected, actual)
	}
}
