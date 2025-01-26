package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input   string
		expected []string
	}{
		{
			input: "  hello world     ",
			expected: []string{"hello","world"},
		},
		{
			input: "PIKACHU DrAcAuFeU bulbizarre",
			expected: []string{"pikachu", "dracaufeu", "bulbizarre"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "    ",
			expected: []string{},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(`TestCleanInput error for case %v with input '%s'
			expected: %v of len %v
			actual: %v of len %v`,i, c.input, c.expected, len(c.expected), actual, len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf(`TestCleanInput error for input '%s on word position %v
				expected: %s
				actual: %s`, c.input, i, expectedWord, word)
			}
		}
	}
}