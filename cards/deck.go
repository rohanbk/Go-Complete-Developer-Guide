package main

import (
	"fmt"
	"io/ioutil"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Create function emitter for instantiating a new deck of cards
func newDeck() deck {
	cards := deck{}

	suits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, (value + " of " + suit))
		}
	}

	return cards
}

// Create function receiver for type 'deck'
func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// Create function to deal out a hand of 'n' cards
// Return two slices: the player's hand and the remaining deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func saveToFile(d deck) error {
	err := ioutil.WriteFile("/home/rohanbk/go/go_complete_developer_guide/cards/hand.txt", []byte(d.toString()), 0644)
	return err
}

func (d deck) toString() string {
	var deckString string
	for _, card := range d {
		if deckString == "" {
			deckString = card
		} else {
			deckString = deckString + " " + card
		}
	}
	return deckString
}
