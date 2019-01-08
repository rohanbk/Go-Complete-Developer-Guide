package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)

	fmt.Println("Saving hand to file...")
	hand.saveToFile()

	readHand := readFromFile()

	fmt.Println("Reading from saved file...")
	fmt.Println(readHand)

}
