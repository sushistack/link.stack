package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLink4Markdown(t *testing.T) {
	tests := []struct {
		url         string
		anchorText  string
		followRate  int
		expected    string
		expectError bool
	}{
		{
			url:         "https://example.com",
			anchorText:  "Example",
			followRate:  0,
			expected:    "[Example](https://example.com){:rel=\"nofollow\"}",
			expectError: false,
		},
		{
			url:         "https://example.com",
			anchorText:  "Example",
			followRate:  100,
			expected:    "[Example](https://example.com)",
			expectError: false,
		},
		{
			url:         "https://example.com",
			anchorText:  "",
			followRate:  100,
			expected:    "[여기](https://example.com)",
			expectError: false,
		},
		{
			url:         "",
			anchorText:  "Example",
			followRate:  0,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.url, func(t *testing.T) {
			link, err := createLink4Markdown(tt.url, tt.anchorText, tt.followRate)
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.expected, link, fmt.Sprintf("%v should be %v", link, tt.expected))
		})
	}
}
