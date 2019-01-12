package main


func main(){

	cards := newDeckFromFile("my")
	// my_cards will work but my will result in an error
	cards.print()

}
