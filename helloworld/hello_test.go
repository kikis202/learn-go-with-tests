package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to someone", func(t *testing.T) {
		got := Hello("Kristap")
		want := "Hello, Kristap!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("empty string defaults to \"World\"", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
