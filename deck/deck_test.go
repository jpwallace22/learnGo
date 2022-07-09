package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	testDeck(d, t)
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")
	newD := newDeckFromFile(("_deckTesting"))

	testDeck(newD, t)

	os.Remove("_deckTesting")
}

func testDeck(d deck, t *testing.T) {
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card to be Two of Spades, but got \"%v\"", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got \"%v\"", d[len(d)-1])
	}
}
