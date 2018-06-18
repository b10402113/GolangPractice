package main

func main() {
	cards := newDeck()
	// cards.print()
	handcards, cards := deal(cards, 5)
	handcards.print()
	cards.print()
}
