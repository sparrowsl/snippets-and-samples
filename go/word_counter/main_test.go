package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4

	result := count(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}
