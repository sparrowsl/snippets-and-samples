package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4

	result := count(b, false, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 \nword3 word4\n")
	expected := 2

	result := count(b, true, false)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 24

	result := count(b, false, true)

	if result != expected {
		t.Errorf("Expected %d, got %d instead.\n", expected, result)
	}
}
