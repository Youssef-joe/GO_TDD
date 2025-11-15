package loop

import (
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeated("a", 5)
	}
}

func assertCorrectMsg(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func TestLoop(t *testing.T) {

	t.Run("5 (a)s", func(t *testing.T) {
		got := Repeated("a", 5)
		expected := "aaaaa"

		assertCorrectMsg(t, got, expected)
	})

	t.Run("10 (b)s", func(t *testing.T) {
		got := Repeated("b", 10)
		expected := "bbbbbbbbbb"
		assertCorrectMsg(t, got, expected)
	})

}
