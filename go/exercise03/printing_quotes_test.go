package main

import "testing"

func TestPrintQuote(t *testing.T) {
	testCases := []struct {
		author, quote, output string
	}{
		{"Don Moen", "Less is More", "Don Moen says, \"Less is More\""},
	}

	for _, test := range testCases {
		result := printQuote(test.author, test.quote)

		if test.output != result {
			t.Errorf("Got %v, want %v", result, test.output)
		}
	}
}
