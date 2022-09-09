package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"lint", "code", "love", "you"},
			expected: "4)lint4)code4)love3)you",
		},
		{
			input:    []string{"testing4)this", ""},
			expected: "13)testing4)this0)",
		},
		{
			input:    []string{},
			expected: "",
		},
	}

	for _, test := range tests {
		actual := encode(test.input)

		assert.Equal(t, test.expected, actual)
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "4)lint4)code4)love3)you",
			expected: []string{"lint", "code", "love", "you"},
		},
		{
			input:    "13)testing4)this0)",
			expected: []string{"testing4)this", ""},
		},
		{
			input:    "",
			expected: []string{},
		},
	}
	for _, test := range tests {
		actual := decode(test.input)
		assert.Equal(t, test.expected, actual)
	}
}
