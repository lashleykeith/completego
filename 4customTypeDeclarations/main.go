package main


func main(){
	// we don't need to write []string{"Ace of Diamonds", newCard()}
	// we made a type so now we can use cards := deck{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string{
	return "Five of Diamonds"
}