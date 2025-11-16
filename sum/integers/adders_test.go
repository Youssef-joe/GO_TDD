package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
}

func assertCorrectMessage(t testing.TB, expected, sum int) {
	if expected != sum {
		t.Errorf("expected %d but got %d", expected, sum)
	}
}

func TestAdder(t *testing.T) {
	t.Run("testing sum", func(t *testing.T) {
		sum := Add(2, 4)
		expected := 6

		assertCorrectMessage(t, expected, sum)
	})
	t.Run("multiply two numbers", func(t *testing.T) {
		sum := Multiply(2, 4)
		expected := 8

		assertCorrectMessage(t, expected, sum)
	})
}
