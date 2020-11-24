package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		// if the length is wrong, notify the test handler
		// FORMATING DIRECTIVES - note the syntax with %d to print the length
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		//if the first card is not Ace of Spades, notify the error handler
		t.Errorf("Expected the first card to be Ace of Spades, instead got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		// if the last card is not King of Clubs, notify the error handler
		t.Errorf("Expected the last card to be King of Clubs, instead got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// remove any potential remnants of a failed test from directory
	os.Remove("_decktesting")

	// create neww deck
	deck := newDeck()
	// attempt to save
	deck.saveToFile("_decktesting")

	// attempt to load deck
	loadedDeck := newDeckFromFile("_decktesting")

	// assert that the loaded deck has correct length
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, instead got %v", len(loadedDeck))
	}

	// clean up the test file after successful test
	os.Remove("_decktesting")
}
