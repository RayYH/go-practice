package basic

import (
	"fmt"
	"math"
	"testing"
)

// 布尔类型 bool
func ExampleDisplayBoolean() {
	var aBool bool
	fmt.Println("bool:", aBool)
	// Output: bool: false
}

// 整型
func ExampleDisplayInt() {
	var anInt int
	var anInt8 int8
	var anInt16 int16
	var anInt32 int32
	var anInt64 int64
	var anUint uint
	var anUint8 uint8
	var anUint16 uint16
	var anUint32 uint32
	var anUint64 uint64
	fmt.Println("Default int values:",
		anInt, anInt8, anInt16, anInt32, anInt32, anInt64,
		anUint, anUint8, anUint16, anUint32, anUint64)
	// Output: Default int values: 0 0 0 0 0 0 0 0 0 0 0
}

// 字符串的默认值是一个空字符串
// 数学运算中应尽可能地使用 float64 位，因为 math 包中提供的类型都是 float64
// 浮点型
func ExampleDisplayFloat() {
	var aFloat32 float32
	var aFloat64 float64
	fmt.Println("Default float values:", aFloat32, aFloat64)
	// Output: Default float values: 0 0
}

// 可以使用 type(variable) 来进行显式转换一个变量的类型
func ExampleDisplayCasts() {
	var n int16 = 34
	var m int32
	m = int32(n)
	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
	// x, y int
	var x, y = 3, 4
	// f float64
	var f = math.Sqrt(float64(x*x + y*y))
	// z uint
	var z = uint(f)
	fmt.Println(x, y, z)
	// Output: 32 bit int is: 34
	// 16 bit int is: 34
	// 3 4 5
}

func TestUint8FromInt(t *testing.T) {
	var aInt = 1
	var expected uint8 = 1
	var got uint8
	got, _ = Uint8FromInt(aInt)
	if expected != got {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func ExampleTypeInterface() {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i
	fmt.Printf("%T %T %T", i, j, k)
	// Output: int float64 complex128
}

func ExampleDisplayComplex() {
	var defaultComplexValue complex128
	fmt.Println("default:", defaultComplexValue)
	re := 1.1
	im := 2.2
	comp := complex(re, im)
	fmt.Println(comp)
	fmt.Println(real(comp), imag(comp))
	// Output: default: (0+0i)
	// (1.1+2.2i)
	// 1.1 2.2
}

func ExampleDisplayByte() {
	var chA byte = 65
	var chB byte = '\x41'
	fmt.Println(chA, chB)
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	// %c - 字符
	// %d/%v - 整型
	// %X - 十六进制标识
	// %U - Unicode格式
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	// Output: 65 65
	// 65 - 946 - 1053236
	// A - β - 􁈴
	// 41 - 3B2 - 101234
	// U+0041 - U+03B2 - U+101234
}

func ExampleDisplayArrays() {
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

	// Output: '\x00''\x00'
	// 000000000000
	// 0.001.112.22
	// 1234
}
