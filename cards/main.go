package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)

	fmt.Println("Saving hand to file...")
	hand.saveToFile("hand.txt")

	readHand := readFromFile("hand.txt")

	fmt.Println("Reading from saved file...")
	fmt.Println(readHand)

	fmt.Println("Shuffle hand...")
	readHand.shuffle()
	readHand.print()

}
