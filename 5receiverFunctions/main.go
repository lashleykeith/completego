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




/*
 Cards
 newDeck -> creates and returns a list of playing cards.
 print	 -> Log out the contents of a deck of cards.
 shuffle -> Shuffles all the cards in a deck.
 deal    -> Create a 'hand' of cards
 saveToFile -> Save a list of cards to a file on the local machine
 newDeckFromFile -> Load a list of cards from the local machine.


*/