package oop

type Vertex struct{ X, Y float64 }

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
