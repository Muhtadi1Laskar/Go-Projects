package main


func main() {
	cards := newDeck()
	dealtCard, resetDeck := deal(cards, 4)
	cards.saveToFile("Cards/new_cards")

	dealtCard.Print()
	resetDeck.Print()
}