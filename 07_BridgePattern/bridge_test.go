package bridgepattern

import (
	"fmt"
	"testing"
)

type RedCircle struct{}

func NewRedCircle() *RedCircle {
	return new(RedCircle)
}

func (rc *RedCircle) DrwaCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, y: %d ]\n", radius, x, y)
}

func TestShapeCircle(t *testing.T) {
	s := NewShapeCircle(5, 2, 2, &RedCircle{})
	s.Draw()
}
