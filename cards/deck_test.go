package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_newDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 6 {
		t.Errorf("Expected length: 6, Actual: %v", len(d))
	}

	testDeck := []string{"Ace of Hearts", "Two of Hearts", "Three of Hearts", "Ace of Spades", "Two of Spades", "Three of Spades"}

	for index, card := range d {
		if testDeck[index] != card {
			t.Errorf("%v is not equal to %v: Order of deck generation is not correct", index, card)
		}
	}
}

//Remove any *_testing.txt files before and after tests attempt to execute
func Test_readFromFile(t *testing.T) {
	fp := "_hand_testing.txt"
	fileCleanup(fp)

	d := newDeck()
	d.saveToFile(fp)

	nd := readFromFile(fp)

	//Check the size of each deck
	if len(nd) != len(d) {
		t.Errorf("Deck lengths do not match. %v is not equal to %v", len(nd), len(d))
	}

	//Check the contents of the file that was saved and read from disk
	for index, card := range nd {
		if d[index] != card {
			t.Errorf("%v is not equal to %v", d[index], card)
		}
	}

	fileCleanup(fp)

}

func fileCleanup(filePath string) {
	err := os.Remove(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}
}
