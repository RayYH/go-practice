package types

import (
	"fmt"
	"math"
)

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
//
// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
//
// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers
//
// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts
//
// byte        alias for uint8
// rune        alias for int32, represents a Unicode code point
//
// uint        either 32 or 64 bits
// int         same size as uint
// uintptr     an unsigned integer large enough to store the uninterpreted bits of a pointer value

// 一个安全的从 int 转为 uint 的 func
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}

	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func DisplayBytes() {
	var chA byte = 65
	var chB byte = '\x41'
	fmt.Println(chA, chB)
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	// %c - character
	// %d/%v - integer
	// %X - UTF-8 bytes
	// %U - UTF-8 core point
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 core point
}

func DisplayArrays() {
	var a [2]byte
	var byteValue byte
	for _, byteValue = range a {
		fmt.Printf("%q", byteValue)
	}

	const N = 3
	var b [2 * N]struct {
		x, y int32
	}
	var structValue struct {
		x, y int32
	}
	fmt.Printf("\n")
	for _, structValue = range b {
		fmt.Printf("%d%d", structValue.x, structValue.y)
	}

	var c [3]*float64
	var d = [3]float64{
		0.00, 1.11, 2.22,
	}
	fmt.Printf("\n")
	var floatValuePointer *float64
	var index int
	for index = range d {
		c[index] = &d[index]
	}

	for _, floatValuePointer = range c {
		fmt.Printf("%.2f", *floatValuePointer)
	}

	fmt.Printf("\n")

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
