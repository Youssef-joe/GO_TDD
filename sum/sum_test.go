package sum

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("that's an array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			// here i'm using %v which stands for "default"
			// also to say: %d stands for ints
			// annnd %q stands for strings
			// that's what we've got so far and there's still alot actually
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// here we won't be using got != want as usual
	// cuz Go does not let you use equality operators with slices
	// and keep in mind that this function expects the elements to be comparable.
	//  So, it can't be applied to slices with non-comparable elements like 2D slices.
	if !slices.Equal(got, want) {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
func TestSumAllNails(t *testing.T) {

	helperMsg := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			// that's kinda slang to put ma signature actually
			t.Errorf("i've wanted %v but i've got %v", want, got)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		helperMsg(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		helperMsg(t, got, want)
	})
}
