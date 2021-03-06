package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 6)

	hand.print()
	remainingCards.print()

	fmt.Println("---------------")

	cards.shuffle()
	cards.print()
}
