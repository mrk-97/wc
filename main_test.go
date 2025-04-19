package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4 word5\n")

	expected := 5

	result := count(b, false, false)

	if result != expected {
		t.Errorf("Expected %d, Got %d instead.\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\nlin2\nline3 word1")

	expected := 3

	result := count(b, true, false)

	if result != expected {
		t.Errorf("Expected %d, Got %d instead.\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("mahi is a good dev")

	expected := 18

	result := count(b, false, true)

	if result != expected {
		t.Errorf("Expected %d, Got %d instead", expected, result)
	}
}
