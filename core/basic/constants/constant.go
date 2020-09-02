package constants

import "fmt"

// The type of constants go supported:
// boolean, rune, integer, floating-point, complex, string

// unsafe.Sizeof, cap, len can be applied to some expressions
// real and imag applied to a complex constant and complex applied to numeric constants
// The boolean truth values are represented by the predeclared constants true and false
// The default type of an untyped constant is bool, rune, int, float64, complex128 or string

// Constant declaration format: const identifier [type] = value
// type can be omitted most of the time, because the compiler can infer its type based on the value of the variable
const Pi = 3.14159

// 因式分解字 - 使用小括号将多个常量放在同一组里
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// 并行赋值
const name, age = "Ray", 24

// iota 只能用于常量 (const 关键字) 的声明而不能用于其他地方
// iota - 从 0 开始，每次换到 新的一行 其值都会自动 + 1
// 这同时表明位于其前面的一行都必须赋值
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// 下面的例子中有三个 iota 位于同一行，因此，w = 0, x = y = z = 1
const (
	w       = iota
	x, y, z = iota, iota, iota
)

// we can omit iota after iota's first occurrence
const (
	d = iota // 0
	e        // 1
	f        // 2
)

// explicitly assignments will not reset the iota value
const (
	g = iota     // g = iota = 0
	h            // h = iota = 1
	i = "string" // i = "string", iota = 2
	j            // here, iota was 3, but j will be assigned last custom const value: "string"
	k = iota     // k = iota = 4
)

const (
	l = 7    // 第一次 iota 前面的常量必须赋值，但此时 iota 已经被初始化为 0
	m = 8    // 此时 iota = 1
	n        // n = m = 8 但是 iota = 2
	o = iota // 3
	p        // 4
)

// use type to alias a known type
type Color int

// The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly.
const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	_                   // skip some values
	_                   // ..
	INDIGO              // 5
	VIOLET              // 6
)

type ByteSize float64

// Go 语言并不显式支持 enum 关键字，不过你可以使用 iota 来达到类似的效果
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

func DisplaySizes() {
	fmt.Println("sizes:", KB, MB, GB, TB, PB, EB, ZB, YB)
}
