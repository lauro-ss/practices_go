package main

import (
	"os"
	"testing"
)

func TestInitDeck(t *testing.T) {
	d := initDeck()

	if len(d) != 40 {
		t.Fatalf("Expected a deck with the size of 40, but recive a sized deck of: %v", len(d))
	}

	if d[0].Name != "Spades" {
		t.Errorf("The deck created it's not correct")
	}
	if d[len(d)-1].Name != "Clubs" {
		t.Errorf("The deck created it's not correct")
	}
}

func TestSaveFileDeck(t *testing.T) {
	fn := "_deckTest.txt"
	defer os.Remove(fn)

	d := initDeck()

	er := d.SaveToFile(fn)
	if er != nil {
		t.Fatalf("%v", er)
	}

	d, er = ReadFromFile(fn)
	if er != nil {
		t.Fatalf("%v", er)
	}

	if len(d) != 40 {
		t.Errorf("Expected a deck with the size of 40, but recive a sized deck of: %v", len(d))
	}
}
