package interfaces

import "fmt"

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

func Classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func ClassifierCaller() {
	Classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}
