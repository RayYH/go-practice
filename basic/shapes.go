package basic

import "math"

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

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {

	perimeter = 2 * (rectangle.Width + rectangle.Height)

	return
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
