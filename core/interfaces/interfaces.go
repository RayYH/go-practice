package interfaces

import (
	"bytes"
	"math"
)

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样
// 比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////

type List []int

func (l List) Len() int {
	return len(l)
}

// Append 定义在指针类型上
func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

// CountInto 定义在值类型上
func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

// LongEnough 定义在值类型上
func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Vector 中存储的所有元素都是 Element 类型，要得到它们的原始类型需要用到类型断言

type Element interface{}
type Vector struct {
	elems []Element
}

func (v *Vector) At(i int) Element {
	return v.elems[i]
}

func (v *Vector) Set(i int, e Element) {
	v.elems[i] = e
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
