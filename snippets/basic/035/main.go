package main

import (
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func myPrint(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			_, _ = os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) { // type switch
		case Stringer:
			_, _ = os.Stdout.WriteString(a.String())
		case int:
			_, _ = os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			_, _ = os.Stdout.WriteString(a)
		// more types
		default:
			_, _ = os.Stdout.WriteString("???")
		}
	}
}

func main() {
	myPrint(Day(1), "was", Celsius(18.36)) // Tuesday was 18.4 °C
}
