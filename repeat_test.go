package main

import "testing"

func TestRepeatString(t *testing.T) {
	var expected, actual string
	expected = "aaa"
	actual = RepeatString("a")

	if expected != actual {
		t.Errorf("Expected " + expected + ", got " + actual)
	}
}
