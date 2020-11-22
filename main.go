package main

func main() {

	cards := newDeck()
	hand, _ := deal(cards, 3)
	// hand.saveToFile("mycards")
	// readFileFromDisc("mycards").print()
	hand.shuffle()
	hand.print()
}
