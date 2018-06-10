
package main 

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)	

}

//You have to give it a type string 
func newCard() string{
	return "Five of Diamonds"
}
/*
if you want to make card an int
func newCard() int {
	return 12
}

*/