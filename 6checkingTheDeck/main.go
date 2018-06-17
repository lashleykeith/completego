package main


func main(){
	// we don't need to write []string{"Ace of Diamonds", newCard()}
	// we made a type so now we can use cards := deck{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	/*
	We can print cards because we gave it the deck receiver
	cards is referenced to d in deck.go like this or self in other langauges
	  
	 */
	cards.print()
}

func newCard() string{
	return "Five of Diamonds"
}