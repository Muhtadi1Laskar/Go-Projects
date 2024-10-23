package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("Cards/new_cards")

	saved_cards := cards.newDeckFromFile("Cards/new_cards")

	saved_cards.Print()
	fmt.Println()

	saved_cards.shuffle()
	saved_cards.Print()
}