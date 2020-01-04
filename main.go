package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("cards")
	// cards.print()
	// cards := newDeck()
	// cards.saveToFile("cards")
	// hand, remainingCards := deal(cards, 5)

	// fmt.Println("-----Hand-----")
	// hand.print()
	// fmt.Println("-----Remaining Cards-----")
	// remainingCards.print()
}
