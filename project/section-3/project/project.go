package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	remainingCards.print()
	hand.print()
}
