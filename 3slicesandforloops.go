package main 

import "fmt"

/*
//1.This appends Six of Spades to the slice
func main(){
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)
}

func newCard() string{
	return "Five of Diamonds"
}

*/

//2 iterate over the list of slices using a forloop
func main(){
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// everytime we go through the list of cards we are putting away the last cards.
	// This is print a list like so 0 Ace of Diamonds, 1 Five of Diamonds
	for i, card := range cards{
		fmt.Println(cards)
	}
}

func newCard() string{
	return "Five of Diamonds"
}