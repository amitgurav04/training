package main

import (
	"fmt"
)

func main() {

	card1, _ := newCard("A", "B")

	fmt.Println(card1)
	fmt.Println(card2)


	card4, _ := newCard("King Of Diamond", "ABC")

	fmt.Println(card4)
}

func newCard(c1, c4 string) (c3, c2 string) {
	c3 = "Manoj"
	c2 = "Rohit"
	//c1 = "Ace Of Spades"
	c1 = "King Of Diamond"

	c4 = "ABC"

	return c1, c4
}
