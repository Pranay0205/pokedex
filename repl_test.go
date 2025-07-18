package main

import "testing"


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
				input:    "foo    bar",
				expected: []string{"foo", "bar"},
			},
			{
				input:    "singleword",
				expected: []string{"singleword"},
			},
			{
				input:    "   ",
				expected: []string{},
			},
			{
				input:    "a   b   c",
				expected: []string{"a", "b", "c"},
			},
			{
				input:    "",
				expected: []string{},
			},
			{
				input:    "   spaced   out   words   ",
				expected: []string{"spaced", "out", "words"},
			},
		}


	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("cleanInput(%q) = %v; want %v", c.input, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] = %q; want %q", c.input, i, word, expectedWord)
			}
		}
	}
}

