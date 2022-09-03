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
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		if got := shape.Area(); got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("area of rectangle", func(t *testing.T) {
		rectangle := &Rectangle{Height: 10.0, Width: 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := &Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
