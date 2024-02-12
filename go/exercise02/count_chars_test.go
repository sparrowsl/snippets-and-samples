package main

import "testing"

func TestCountCharacters(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{"hello", 5},
		{"hell", 4},
	}

	for _, test := range testCases {
		result := countChars(test.input)

		if result != test.output {
			t.Errorf("Got %v, want %v", result, test.output)
		}
	}
}
