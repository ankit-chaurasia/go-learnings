package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// fmt.Println(card)
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards.print()
	// fmt.Println(cards)
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	cards := newDeckFromFile("my_cards")
	cards.print()
	fmt.Println("---- Shuffling the cards -----")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
