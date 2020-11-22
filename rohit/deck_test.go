package rohit

import (
	"reflect"
	"testing"
)

var failureCount int

func TestNewDeck(t *testing.T) {
	expectedLengthOfDeck := 52
	cards := NewDeck()
	if len(cards) != expectedLengthOfDeck {
		failureCount++
		t.Errorf("Length Of Cards is not %d. It's:%d", expectedLengthOfDeck, len(cards))
	}

	firstCard := cards[0]
	expectedCard := "Ace Of Spades"
	if firstCard != expectedCard {
		t.Errorf("Expected Card: %s, Actual Card:%s", expectedCard, firstCard)
	}

	if reflect.TypeOf(cards[0]) != reflect.TypeOf("abc"){
		t.Errorf("Expected String Type. But returning non string")
	}
}
