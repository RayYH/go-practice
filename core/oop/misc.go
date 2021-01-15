package oop

import "fmt"

func Classifier(items ...interface{}) {
	for i, x := range items {
		// A type switch is a construct that permits several type assertions in series.
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
