package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractBetweenBraces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		found    bool
	}{
		{
			input:    "This is a {{sample}} string.",
			expected: "sample",
			found:    true,
		},
		{
			input:    "No braces here!",
			expected: "",
			found:    false,
		},
		{
			input:    "Multiple {{braces}} here {{and}} there.",
			expected: "braces",
			found:    true,
		},
		{
			input:    "Unclosed {{braces",
			expected: "",
			found:    false,
		},
		{
			input:    "Only {{one}} brace {{pair}}",
			expected: "one",
			found:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ExtractBetweenBraces(tt.input)
			assert.Equal(t, tt.expected, result, fmt.Sprintf("%v should be %v", result, tt.expected))
		})
	}
}
