package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type DeckType []string

func (deck DeckType) print() {
	for _, card := range deck {
		fmt.Println(card)
	}
}

func newDeck() DeckType {
	deck := DeckType{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+" of "+suit)
		}
	}

	return deck
}

func deal(deck DeckType, handSize int) (DeckType, DeckType) {
	return deck[:handSize], deck[handSize:]
}

func (d DeckType) toString() string {
	return strings.Join([]string(d), ",")
}

func (d DeckType) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) DeckType {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		println("Error: ", err)
		os.Exit(1)
	}

	return DeckType(strings.Split(string(bs), ","))
}

func (d DeckType) shuffle() DeckType {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i, _ := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
