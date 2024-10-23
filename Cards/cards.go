package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := []string{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, index int) (deck, deck) {
	return d[:index], d[index:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}