package oop

import "math"

type Vertex struct{ X, Y float64 }

type Abser interface{ Abs() float64 }

type MyFloat64 float64

// There are two reasons to use a pointer receiver.
// 1. The method can modify the value that its receiver points to.
// 2. Avoid copying the value on each method call.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (mf MyFloat64) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}
