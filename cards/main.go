package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.printDeck()
}

func newCard() string {
	return "5 of diamonds"
}
