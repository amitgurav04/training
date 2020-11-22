package rohit

import (
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
	expectedProduct := 6
	actualProduct := Multiplication(n1, n2)
	if actualProduct != 6 {
		t.Errorf("Product of %d and %d is not: %d. Getting product: %d", n1, n2, expectedProduct, actualProduct)
	}
}
