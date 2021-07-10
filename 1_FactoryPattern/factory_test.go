package factorypattern

import "testing"

func TestRectangle(t *testing.T) {
	rectangle := NewShapeFactory("rectangle")
	if rectangle == nil {
		t.Fatal("factory new failed")
	}
	res := rectangle.Draw()
	if res != "rectangle" {
		t.Errorf("error, wanted: %v, got: %v", "rectangle", res)
	}
}

func TestSquare(t *testing.T) {
	square := NewShapeFactory("square")
	if square == nil {
		t.Fatal("factory new failed")
	}
	res := square.Draw()
	if res != "square" {
		t.Errorf("error, wanted: %v, got: %v", "square", res)
	}
}

func TestCircle(t *testing.T) {
	circle := NewShapeFactory("circle")
	if circle == nil {
		t.Fatal("factory new failed")
	}
	res := circle.Draw()
	if res != "circle" {
		t.Errorf("error, wanted: %v, got: %v", "circle", res)
	}
}
