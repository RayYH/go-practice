package constants

import "fmt"

// The type of constants go supported:
// boolean, rune, integer, floating-point, complex, string

// unsafe.Sizeof, cap, len can be applied to some expressions
// real and imag applied to a complex constant and complex applied to numeric constants
// The boolean truth values are represented by the predeclared constants true and false
// The default type of an untyped constant is bool, rune, int, float64, complex128 or string

// constant declaration format: const identifier [type] = value
// type can be omitted most of the time, because the compiler can infer its type based on the value of the variable
const Pi = 3.14159

const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

const name, age = "Ray", 24

const (
	a = iota // a == 0 (iota == 0)
	b = iota // b == 1 (iota == 1)
	c = iota // c == 2 (iota == 2)
)

// By definition, multiple uses of iota in the same ConstSpec all have the same value
const (
	w       = iota             // w == 0           (iota == 0)
	x, y, z = iota, iota, iota // x == y == z == 1 (iota == 1)
)

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                 //                        (iota == 2, unused)
	bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
)

// we can omit iota after iota's first occurrence
const (
	d = iota // d == 0 (iota == 0)
	e        // e == 1 (iota == 1)
	f        // f == 2 (iota == 2)
)

// explicitly assignments will not reset the iota value
const (
	g = iota     // g == 0              (iota == 0)
	h            // h == 1              (iota == 1)
	i = "string" // i == "string"       (iota == 2)
	j            // j == i == "string"  (iota == 3)
	k = iota     // k == 4              (iota == 4)
)

const (
	l = 7    // l == 7      (iota == 0)
	m = 8    // m == 8      (iota == 1)
	n        // n == m == 8 (iota == 2)
	o = iota // o == 3      (iota == 3)
	p        // p == 4      (iota == 4)
)

// use type to alias a known type
type Color int

// The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly.
const (
	RED    Color = iota // RED == 0    (iota == 0)
	ORANGE              // ORANGE == 1 (iota == 1)
	YELLOW              // YELLOW == 2 (iota == 2)
	_                   //             (iota == 3)
	_                   //             (iota == 4)
	INDIGO              // INDIGO == 5 (iota == 5)
	VIOLET              // VIOLET == 6 (iota == 6)
)

type ByteSize float64

const (
	_           = iota             // ignore 0
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

// Numeric constants represent exact values of arbitrary precision and do not overflow.
const HigherPrecisionPi = 3.14159265358979323846264338327950288419716939937510582097494459
const LessThanOne = 3.141592653589793 / HigherPrecisionPi

// print value of defined constants
func DisplaySizes() {
	fmt.Println("sizes:", KB, MB, GB, TB, PB, EB, ZB, YB)
}
