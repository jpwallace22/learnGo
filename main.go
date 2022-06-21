package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println(hand.toString())
	remainingDeck.print()
}
