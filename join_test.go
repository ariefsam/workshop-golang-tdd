package main

import "testing"

func TestJoin(t *testing.T) {
	var sliceText []string
	sliceText = []string{"Halo", "halo", "Bandung"}
	actual := Join(sliceText)
	expected := "Halo halo Bandung"

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}
