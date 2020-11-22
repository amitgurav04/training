package main

import (
	"fmt"
	"github.com/amitgurav04/training/rohit"
)

func main() {

	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	modifyArr(arr)
	//arr = append(arr, 5)
	fmt.Print(arr)
	cards := rohit.Deck{}
	if len(cards) != 0 {
		fmt.Println("CardSet:")
		cards.Print()
	}

	cardSuit := []string{
		"Spades",
		"Heart",
		"Diamond",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
	}

	for _, suit := range cardSuit {
		for _, val := range cardValues {
			cards = append(cards, val+" Of "+suit)
		}
	}
	fmt.Println("\n==============CardSet Now:======================")
	cards.Print()
	d1, d2 := rohit.Deal(cards, 5)

	fmt.Println("\n===================First Set====================")
	d1.Print()

	fmt.Println("\n===================Second Set===================")
	d2.Print()

	fmt.Println("=======Deck Of Cards========:\n", cards.ToString())
	fmt.Println("============Save to File small_deck.txt=============")
	err := cards.SaveToFile("small_deck.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File created successfully with deck")
	}

	fmt.Println("============Read from File small_deck.txt and convert it into Deck =============")
	filename := "small_deck.txt"
	d4 := rohit.NewDeckFromFile(filename)
	fmt.Println("============Print Deck 4 =============")
	d4.Print()

	fmt.Println("============Shuffle =============")
	d4.Shuffle()

	fmt.Println("============Deck 4 after shuffle=============")
	d4.Print()

	friendsCount := 11
	fmt.Printf("Hello %d Friends\n", friendsCount)

	deck := rohit.NewDeck()

	fmt.Printf("\n=======================New Deck ========================\n %+v\n", deck)
	deck.Print()
}

func modifyArr(a []int) {
	a[2] = 33
	return
}
