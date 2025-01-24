package structs

import "testing"

func TestRectangle(t *testing.T) {
	rect := Rectangle{10.0, 10.0}

	t.Run("perimeter", func(t *testing.T) {
		got := Perimeter(rect)
		want := 40.0

		if got != want {
			t.Errorf("got %2.f want %2.f", got, want)
		}
	})

	t.Run("area", func(t *testing.T) {
		got := Area(rect)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
