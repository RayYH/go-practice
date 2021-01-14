package oop

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
