package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	testConfigFile := "../configs/config-test.yaml"

	configContent := "app:\n" +
		"  name: \"Link Stack\"\n" +
		"datasource:\n" +
		"  uri: \"{{MONGODB_URI}}\"\n"

	// 설정 파일 작성
	err := os.WriteFile(testConfigFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Error writing test config file: %v", err)
	}
	defer func() {
		// 파일 삭제
		if err := os.Remove(testConfigFile); err != nil {
			t.Fatalf("Error removing test config file: %v", err)
		}
	}()

	env := LoadEnvironment(nil)
	// Test LoadConfig
	expectedConfig := map[string]interface{}{
		"app": map[string]interface{}{
			"name": "Link Stack",
		},
		"datasource": map[string]interface{}{
			"uri": env["MONGODB_URI"],
		},
	}

	config := LoadConfig(testConfigFile)
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)
}

func TestReplaceEnvVariables(t *testing.T) {
	env := map[string]string{
		"DB_URI":   "http://db.example.com",
		"APP_NAME": "TestApp",
	}

	tests := []struct {
		name     string
		input    map[string]interface{}
		expected map[string]interface{}
	}{
		{
			name: "Replace placeholders with environment variables",
			input: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "{{APP_NAME}}",
				},
				"datasource": map[string]interface{}{
					"uri": "{{DB_URI}}",
				},
			},
			expected: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "TestApp",
				},
				"datasource": map[string]interface{}{
					"uri": "http://db.example.com",
				},
			},
		},
		{
			name: "No placeholders in input",
			input: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "StaticName",
				},
				"datasource": map[string]interface{}{
					"uri": "StaticURI",
				},
			},
			expected: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "StaticName",
				},
				"datasource": map[string]interface{}{
					"uri": "StaticURI",
				},
			},
		},
		{
			name: "Placeholder with missing environment variable",
			input: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "{{NON_EXISTENT}}",
				},
			},
			expected: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "",
				},
			},
		},
		{
			name: "Nested placeholders",
			input: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "{{APP_NAME}}",
					"details": map[string]interface{}{
						"uri": "{{DB_URI}}",
					},
				},
			},
			expected: map[string]interface{}{
				"app": map[string]interface{}{
					"name": "TestApp",
					"details": map[string]interface{}{
						"uri": "http://db.example.com",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceEnvVariables(tt.input, env)
			assert.Equal(t, tt.expected, tt.input)
		})
	}
}

func TestGetEnvValue(t *testing.T) {
	env := map[string]string{
		"DB_URI":   "http://db.example.com",
		"APP_NAME": "TestApp",
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid placeholder with environment variable",
			input:    "{{DB_URI}}",
			expected: "http://db.example.com",
		},
		{
			name:     "Valid placeholder but no environment variable",
			input:    "{{NON_EXISTENT}}",
			expected: "",
		},
		{
			name:     "Plain string without placeholder",
			input:    "Just a string",
			expected: "Just a string",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Placeholder with spaces",
			input:    "{{ APP_NAME }}",
			expected: "TestApp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getEnvValue(env, tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestExtractBetweenBraces(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "This is a {{sample}} string.",
			expected: "sample",
		},
		{
			input:    "This is a {{ sample2 }} string.",
			expected: "sample2",
		},
		{
			input:    "No braces here!",
			expected: "No braces here!",
		},
		{
			input:    "Multiple {{braces}} here {{and}} there.",
			expected: "braces",
		},
		{
			input:    "Unclosed {{braces",
			expected: "Unclosed {{braces",
		},
		{
			input:    "Only {{one}} brace {{pair}}",
			expected: "one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ExtractBetweenBraces(tt.input)
			assert.Equal(t, tt.expected, result, fmt.Sprintf("%v should be %v", result, tt.expected))
		})
	}
}
