package main

import (
	"fmt"
)

func main() {

	cards := newDeck()
	cards.shuffle()

	_, remainingCards := deal(cards, 5)

	fmt.Println("Hand: ")
	cards.print()

	remainingCards.saveToFile("./cards")
}
