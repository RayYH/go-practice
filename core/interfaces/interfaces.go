package interfaces

import "math"

////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Vertex struct{ X, Y float64 }
type Abser interface{ Abs() float64 }
type MyFloat64 float64

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (mf MyFloat64) Abs() float64 {
	if mf < 0 {
		return float64(-mf)
	}
	return float64(mf)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

/// 接口

type Shaper interface{ Area() float64 }

/// 结构体

type Square struct{ side float64 }
type Rectangle struct{ length, width float64 }

/// 实现接口

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

func (rec *Rectangle) Area() float64 {
	return rec.width * rec.length
}

/// 接口作为参数类型

func GetType(shaper Shaper) string {
	switch shaper.(type) {
	case *Rectangle:
		return "Rec"
	case *Square:
		return "Squ"
	default:
		return "Unknown"
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
