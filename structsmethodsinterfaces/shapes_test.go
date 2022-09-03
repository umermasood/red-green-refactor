package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Height: 10.0, Width: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{Height: 10.0, Width: 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
