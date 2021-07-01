package basic_types

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

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerType(t *testing.T) {
	t.Run("integer types", func(t *testing.T) {
		// int 和 uint 在 32 位操作系统上占 32 位，在 64 位操作系统上占 64 位
		// uintptr 的长度被设定为足够存放一个指针
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
		assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0", fmt.Sprint(anInt, anInt8, anInt16, anInt32, anInt32, anInt64,
			anUint, anUint8, anUint16, anUint32, anUint64))
	})

	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
	// The runtime system considers an `unsafe.Pointer` as a reference to an object, which keeps the
	// object alive for GC. It does not consider a `uintptr` as such a reference.
	// (That is, while unsafe.Pointer has a pointer type, uintptr has integer type.)
	t.Run("never use uintptr", func(t *testing.T) {
		var anUintPtr uintptr
		assert.Equal(t, uintptr(0), anUintPtr)
	})
}

// float 有两种精度 float32 和 float64，尽量使用 float64，因为 math 中的包函数中可接受的参数类型大多是 float64
func TestFloatType(t *testing.T) {
	var aFloat32 float32
	var aFloat64 float64
	assert.Equal(t, "0 0", fmt.Sprint(aFloat32, aFloat64))
	// Go 中 float32 精确到小数点后 7 位 (有效数字 6 位)，float64 精确到小数点后 15 位 (有效数字 14 位)
	assert.Equal(t, "3.1415927, 3.141592653589793", fmt.Sprintf("%v, %v", float32(math.Pi), math.Pi))
}

// 字面量浮点数是可以直接用 == 进行比较的 (与 Python 相同)，但是大多数时候我们都会使用精度来判断两个浮点数是否相同
func TestComparingFloatVariables(t *testing.T) {
	isEqual := func(f1, f2, p float64) bool {
		return math.Abs(f1-f2) < p
	}

	returnFloat := func() float64 {
		return 3.00000001
	}

	assert.True(t, returnFloat() == returnFloat())
	assert.True(t, isEqual(3.000001, 3.000002, 0.0001))
	assert.False(t, isEqual(3.000001, 3.000002, 0.000000001))
}

func TestComplexType(t *testing.T) {
	var defaultComplexValue complex128
	assert.Equal(t, "(0+0i)", fmt.Sprint(defaultComplexValue))
	re, im := 1.1, 2.2
	comp := complex(re, im)
	assert.Equal(t, "(1.1+2.2i)", fmt.Sprint(comp))
	assert.Equal(t, 1.1, real(comp))
	assert.Equal(t, 2.2, imag(comp))

	// default is complex128
	var vOfComplex64 complex64
	assert.Equal(t, complex64(0+0i), vOfComplex64)
}

func TestUint8FromInt(t *testing.T) {
	uint8FromInt := func(n int) (uint8, error) {
		if 0 <= n && n <= math.MaxUint8 {
			return uint8(n), nil
		}
		return 0, fmt.Errorf("%d is out of the uint8 range", n)
	}

	var aInt = 1
	var expected uint8 = 1
	var got uint8
	got, _ = uint8FromInt(aInt)
	assert.Equal(t, got, expected)
}

// byte 类型是 uint8 的别名，适用于传统的 ASCII 编码
// rune 类型是 int32 的别名，适用于 UTF-8 编码
func TestBytesPresentation(t *testing.T) {
	// byte 是 uint8 的别名，因此 byte 只能定义 ASCII 字符
	var chA byte = 65
	var chB byte = '\x41'
	// 多字节字符我们可以用 int 类型变量存储
	// 在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U
	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	// 所存储的字符最终呈现什么，取决于我们如何解析它
	// %c - character
	// %d/%v - integer
	// %X - UTF-8 bytes
	// %U - UTF-8 core point
	assert.Equal(t, "65 65", fmt.Sprint(chA, chB)) // 65, 65
	assert.Equal(t, "65 946 1053236", fmt.Sprint(ch1, ch2, ch3))
	assert.Equal(t, fmt.Sprintf("%d - %d - %d", ch1, ch2, ch3), "65 - 946 - 1053236")
	assert.Equal(t, fmt.Sprintf("%c - %c - %c", ch1, ch2, ch3), "A - β - \U00101234")
	assert.Equal(t, fmt.Sprintf("%X - %X - %X", ch1, ch2, ch3), "41 - 3B2 - 101234")
	assert.Equal(t, fmt.Sprintf("%U - %U - %U", ch1, ch2, ch3), "U+0041 - U+03B2 - U+101234")
}
