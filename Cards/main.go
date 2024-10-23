package main


func main() {
	cards := newDeck()
	dealtCard, resetDeck := deal(cards, 4)

	dealtCard.Print()
	resetDeck.Print()
}