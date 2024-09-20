package utils_test

import (
	"testing"

	"github.com/yashikota/chronotes/pkg/utils"
)

func TestMd2HTML(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic Markdown",
			input:    "# Hello\nThis is a test",
			expected: "<h1 id=\"hello\">Hello</h1>\n\n<p>This is a test</p>\n",
		},
		{
			name:     "List",
			input:    "- Item 1\n- Item 2",
			expected: "<ul>\n<li>Item 1</li>\n<li>Item 2</li>\n</ul>\n",
		},
		{
			name:     "Link",
			input:    "[example](https://example.com)",
			expected: "<p><a href=\"https://example.com\" target=\"_blank\">example</a></p>\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := utils.Md2HTML([]byte(tc.input))
			if string(result) != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, string(result))
			}
		})
	}
}
