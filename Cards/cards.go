package main

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