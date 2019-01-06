package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Create function receiver for type 'deck'
func (d deck) printDeck() {
	fmt.Println("Printing cards...")
	for index, card := range d {
		fmt.Println(index, card)
	}
}
