package main

import "fmt"

func main() {

	fmt.Println("Dead.")


	//return

	//cards := []string{"King of Diamonds"}
	cards := []string{"King of Diamonds", newCard("Queen", "Spades")}

	fmt.Println("Still Dead.")
	fmt.Println(cards)
	return

	for _, card :=  range which_card {

		for _, suit :=  range which_suit {

			cards = append(cards, newCard(card, suit))
		}
	}
	cards = append(cards, "King of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard(card, suit string) string {
	return fmt.Sprintf("%s of %s", card, suit) 
}

var  which_card = []string{"Ace", "Two", "Three", "Four", "Five"}
var  which_suit = []string{"Clubs", "Diamonds", "Hearts", "Spades"}

