package interfaces

import "math"

////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 虽然 Go 中没有类和继承的概念，但是 Go 语言里有非常灵活的接口概念，通过它可以实现很多面向对象的特性。
// 接口定义了一组不包含实现的抽象方法，Go 语言中的接口都很简短，通常它们只会包含 0~3 个方法。
// 在 Go 的底层，接口可以被认为一个是一个包含了值和类型的元组：(value, type)。
// 接口值保存了一个具体底层类型的具体值，接口值调用方法时会执行其底层类型的同名方法。

////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Abser interface{ Abs() float64 }

type Vertex struct{ X, Y float64 }
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
	// .(type) 关键字可以获取变量的类型
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
