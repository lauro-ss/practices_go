package main

import "testing"

func TestInitDeck(t *testing.T) {
	d := initDeck()

	if len(d) != 40 {
		t.Errorf("Expected a deck with the size of 40, but recive a sized deck of: %v", len(d))
	}
}
