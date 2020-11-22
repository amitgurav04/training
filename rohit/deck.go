package rohit

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func (d Deck) Shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Deck) Print() {
	for _, val := range d {
		fmt.Println(val)
	}
}

func Deal(d Deck, handshake int) (Deck, Deck) {
	return d[:handshake], d[handshake:]
}
func NewCards() string {
	return "Nine Of String"
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println(string(bs))
	cardsString := strings.Split(string(bs),",")
	return Deck(cardsString)
}
