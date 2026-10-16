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
			input:    "	hello 	world	",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "one two three four five words",
			expected: []string{"one", "two", "three", "four", "five", "words"},
		},
		{
			input:    "Some UpperCase WORDS",
			expected: []string{"some", "uppercase", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length %d, received length %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word: %s, received word: %s", expectedWord, word)
			}
		}
	}
}
