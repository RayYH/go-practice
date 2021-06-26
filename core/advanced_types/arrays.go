package advanced_types

import "fmt"

func DisplayArrays() {
	// byte array
	var a [2]byte
	var byteValue byte
	for _, byteValue = range a {
		fmt.Printf("%q\n", byteValue)
	}

	// struct array
	const N = 2
	var b [2 * N]struct {
		x, y int32
	}
	var structValue struct {
		x, y int32
	}
	for _, structValue = range b {
		fmt.Printf("%d %d\n", structValue.x, structValue.y)
	}

	// float64 array
	var c [3]*float64
	var d = [3]float64{
		0.00, 1.11, 2.22,
	}
	var floatValuePointer *float64
	var index int
	for index = range d {
		c[index] = &d[index]
	}
	for _, floatValuePointer = range c {
		fmt.Printf("%.2f\n", *floatValuePointer)
	}

	// nested array
	var e = [2][2]int{
		{1, 2},
		{3, 4},
	}
	var arrayValue [2]int
	var intValue int
	for _, arrayValue = range e {
		for _, intValue = range arrayValue {
			fmt.Print(intValue)
		}
	}
}
