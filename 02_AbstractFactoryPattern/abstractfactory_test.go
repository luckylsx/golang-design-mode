package abstractfactorypattern

import "testing"

func TestShape(t *testing.T) {
	factory := FactoryProducer("shape")
	if factory == nil {
		t.Fatal("error")
	}
	shape := factory.NewFactoryShape("square")
	if shape == nil {
		t.Fatal("error")
	}
	if res := shape.Draw(); res != "square" {
		t.Errorf("error, wanted: %v, got: %v", "square", res)
	}
}

func TestColor(t *testing.T) {
	factory := FactoryProducer("color")
	if factory == nil {
		t.Fatal("error")
	}
	color := factory.NewFactoryColor("red")
	if color == nil {
		t.Fatal("error")
	}
	if res := color.fill(); res != "Red" {
		t.Errorf("error, wanted: %v, got: %v", "Red", res)
	}
}
