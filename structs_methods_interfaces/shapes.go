package shapes

import "math"

// Rectangle has a method called Area that returns a float64, so it satifsfies the Shape interface
// Circle also has a method called Area that returns a float64, so that also satisfies the Shape interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Give Rectangle a method called Area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Give Circle a method called Area
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
