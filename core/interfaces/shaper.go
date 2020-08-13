package main

type Shaper interface {
	Area() float64
}

type Square struct {
	side float64
}

func (sq *Square) Area() float64 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float64
}

func (rec *Rectangle) Area() float64 {
	return rec.width * rec.length
}
