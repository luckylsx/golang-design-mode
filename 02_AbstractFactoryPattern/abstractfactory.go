package abstractfactorypattern

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

type Color interface {
	fill() string
}

type Red struct{}

func (r *Red) fill() string {
	return "Red"
}

type Green struct{}

func (r *Green) fill() string {
	return "Green"
}

type Blue struct{}

func (r *Blue) fill() string {
	return "Blue"
}

//AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	NewFactoryShape(t string) Shape
	NewFactoryColor(t string) Color
}

type ShapeFactory struct{}
type ColorFactory struct{}

func (s *ShapeFactory) NewFactoryShape(t string) Shape {
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
func (s *ShapeFactory) NewFactoryColor(t string) Color {
	return nil
}

func (c *ColorFactory) NewFactoryColor(t string) Color {
	if t == "" {
		return nil
	}
	switch t {
	case "red":
		return new(Red)
	case "green":
		return new(Green)
	case "blue":
		return new(Blue)
	}
	return nil
}

func (c *ColorFactory) NewFactoryShape(t string) Shape {
	return nil
}

//FactoryProducer 超级工厂类，用于获取工厂实例
func FactoryProducer(factory string) AbstractFactory {
	if factory == "" {
		return nil
	}
	switch factory {
	case "color":
		return new(ColorFactory)
	case "shape":
		return new(ShapeFactory)
	default:
		return nil
	}
}
