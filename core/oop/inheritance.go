package oop

import "math"

type Camera struct{}
type Phone struct{}
type CameraPhone struct {
	Camera
	Phone
}

type Point struct{ x, y float64 }
type NamedPoint struct {
	Point
	name string
}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

func (p *Phone) Call() string {
	return "Ring Ring"
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
