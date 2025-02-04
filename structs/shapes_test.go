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
		got := rect.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
