package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// we do use %2f for float64
		t.Errorf("wanted %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	//The only new syntax here is creating an "anonymous struct", areaTests.
	//  We are declaring a slice of structs by using []struct with two fields,
	//   the shape and the want. Then we fill the slice with cases.
	areaTest := []struct {
		name string
		shape Shape
		want float64
		hasArea float64
	} {
		{name: "Rectangle",shape: Rectangle{Width: 12,Height: 6},hasArea: 72.0},
		{name: "Circle",shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle",shape: Triangle{Base: 12,Height: 6},hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}