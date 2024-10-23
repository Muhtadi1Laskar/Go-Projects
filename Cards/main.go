package main


func main() {
	cards := newDeck()
	cards.saveToFile("Cards/new_cards")

	saved_cards := cards.newDeckFromFile("Cards/new_cards")

	saved_cards.Print()
}