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
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: &Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: &Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: &Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %g want %g", test.shape, got, test.hasArea)
			}
		})
	}
}
