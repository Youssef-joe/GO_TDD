package main

import "testing"

func isError(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}
}

func TestHello(t *testing.T) {
	t.Run("just saying hello", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		isError(t, got, want)
	})

	t.Run("saying hello to diff langs", func(t *testing.T) {
		got := Hello("Joe", "German")
		want := "Hallo, Joe"

		isError(t, got, want)
	})

	t.Run("saying hello for spanish people", func(t *testing.T) {
		got := Hello("joe", "Spanish")
		want := "Hola, joe"

		isError(t, got, want)
	})

}
