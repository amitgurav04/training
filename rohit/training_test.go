package rohit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddition(t *testing.T) {
	sum := Addition(3, 2)
	if sum != 5 {
		t.Error("Addition of 2 and 3 is not", sum)
	}
}

func TestSubtraction(t *testing.T) {
	sum := Subtraction(3, 2)
	if sum != 1 {
		t.Error("Addition of 2 and 3 is not", sum)
	}
}

func TestMultiplication(t *testing.T) {
	n1 := 3
	n2 := 2
	assert := assert.New(t)
	expectedProduct := 6
	actualProduct := Multiplication(n1, n2)
	assert.Equal(expectedProduct, actualProduct, "Are tu pagal hai. Output galat aayaa.")
}
