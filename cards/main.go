package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)

	err := hand.saveToFile()

	if err != nil {
		panic(err)
	}

	readHand, err := readFromFile()

	if err != nil {
		panic(err)
	}

	fmt.Println("Reading from saved file...")
	fmt.Println(readHand)

}
