package factorypattern

type Shape interface {
	Draw() string
}

type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return "rectangle"
}

type Square struct{}

func (s *Square) Draw() string {
	return "square"
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "circle"
}

func NewShapeFactory(t string) Shape {
	if t == "" {
		return nil
	}
	switch t {
	case "rectangle":
		return new(Rectangle)
	case "square":
		return new(Square)
	case "circle":
		return new(Circle)
	}
	return nil
}
