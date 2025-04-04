package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   Nyellow Werld    ",
			expected: []string{"nyellow", "werld"},
		},
		{
			input:    "brokie lemonadass",
			expected: []string{"brokie", "lemonadass"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d words", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("cleanInput(%q) returned %q at position %d, expected %q", c.input, word, i, expectedWord)
			}
		}
	}
}
