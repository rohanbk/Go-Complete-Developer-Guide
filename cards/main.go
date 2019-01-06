package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Printing hand...")
	hand.printDeck()

	fmt.Println("Printing remaining deck...")
	remainingDeck.printDeck()
}
