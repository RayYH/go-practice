package interfaces

import (
	"bytes"
	"math"
)

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

type Shaper interface{ Area() float64 }

type Square struct{ side float64 }
type Rectangle struct{ length, width float64 }

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

func (rec *Rectangle) Area() float64 {
	return rec.width * rec.length
}

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Element interface{}
type Vector struct {
	elems []Element
}

func (v *Vector) At(i int) Element {
	return v.elems[i]
}

func (v *Vector) Set(i int, e Element) {
	if len(v.elems) == cap(v.elems) {
		v.elems = append(v.elems, make([]Element, len(v.elems)+1)...)
	}
	v.elems[i] = e
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////
