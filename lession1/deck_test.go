package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if deck == nil {
		t.Errorf("Error deck length got %v", len(deck))
	}
}

func TestSaveToDeckAndNewDeckFormFile(t *testing.T) {
	os.Remove("deck_testing.txt")

	deck := newDeck()
	deck.saveToFile("deck_testing.txt")

	loadDeck := newDeckFromFile("deck_testing.txt")
	if len(loadDeck) == 0 {
		t.Errorf("Load from file is Error %v", len(loadDeck))
	}

	os.Remove("deck_testing.txt")
}
