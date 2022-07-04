package main

func main() {
	cards := newDeckFromFile("the_cards")
	cards.shuffle()
	cards.print()
}
