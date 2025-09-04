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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test           should  be 4  ",
			expected: []string{"test", "should", "be", "4"},
		},
		{
			input:    "THIS TEST SHOULD PASS",
			expected: []string{"this", "test", "should", "pass"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Slice lengths don't match.\nexpected:%d\nactual:%d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Resulting string does not match.\nexpected:%s\nactual%s", expectedWord, word)
			}
		}
	}
}
