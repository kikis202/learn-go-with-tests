package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone", func(t *testing.T) {
		got := Hello("Kristap", "")
		want := "Hello, Kristap!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to \"World\"", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Kristap", "Spanish")
		want := "Halo, Kristap!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Kristap", "French")
		want := "Bonjour, Kristap!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
