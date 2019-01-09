package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Create function emitter for instantiating a new deck of cards
func newDeck() deck {
	cards := deck{}

	suits := []string{"Hearts", "Spades"}
	values := []string{"Ace", "Two", "Three"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, (value + " of " + suit))
		}
	}

	return cards
}

// Create function receiver for type 'deck' to print all cards in the deck
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

func (d deck) saveToFile(filePath string) {
	err := ioutil.WriteFile(filePath, []byte(d.toString()), 0644)

	if err != nil {
		panic(err)
	}
}

func readFromFile(filePath string) deck {
	bs, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) toString() string {
	deckString := strings.Join(d, ",")
	return deckString
}

// Create function receiver for type 'deck' to randomize the order of the cards in the deck
func (d deck) shuffle() {
	deckLength := len(d)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		r := r.Intn(deckLength)
		d[i], d[r] = d[r], d[i]
	}
}
