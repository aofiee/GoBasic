package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 6)
	fmt.Println([]byte("การที่เรา"))
	cards.shuffle()
	cards.print()
	fmt.Println("-----------------------")
	hand.print()
	fmt.Println("-----------------------")
	remainingDeck.print()
	fmt.Println(cards.toString())
	cards.saveToFile("aofiee.txt")
	hand.saveToFile("onhand.txt")
	remainingDeck.saveToFile("remaining.txt")

	onhandDeck := newDeckFromFile("onhand.txt")
	onhandDeck.print()

}
