package basic_types

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegers(t *testing.T) {
	// 所有整型类型的变量的零值都是 0
	t.Run("vars' zero values are all 0", func(t *testing.T) {
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

	// 尽量避免直接使用 uintptr，关于 uint 和 uintptr 的区别可以查看下面的链接
	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
	t.Run("never use uintptr", func(t *testing.T) {
		var anUintPtr uintptr
		assert.Equal(t, uintptr(0), anUintPtr)
	})
}

// float 有两种精度 float32 和 float64，尽量使用 float64，因为 math 中的包函数中可接受的参数类型大多是 float64
func TestFloatPointNumbers(t *testing.T) {
	var aFloat32 float32
	var aFloat64 float64
	assert.Equal(t, "0 0", fmt.Sprint(aFloat32, aFloat64))
	// Go 中 float32 精确到小数点后 7 位 (有效数字 6 位)，float64 精确到小数点后 15 位 (有效数字 14 位)
	assert.Equal(t, "3.1415927, 3.141592653589793", fmt.Sprintf("%v, %v", float32(math.Pi), math.Pi))
}

func TestDifferentRadixes(t *testing.T) {
	// 二进制的高位 1 不代表符号位，负数必须通过取反运算符 - 来声明
	t.Run("bin", func(t *testing.T) {
		a := 0b10000001
		assert.Equal(t, 129, a)
		b := -0b00000001
		assert.Equal(t, -1, b)
		c := 0b100000000000000000000000000000000000000000000000000000001
		assert.Greater(t, c, 0)
	})

	// 前缀 0 表示 8 进制数
	t.Run("oct", func(t *testing.T) {
		a := 077
		assert.Equal(t, 63, a)
	})

	// 前缀 0x 表示 16 进制数
	t.Run("hex", func(t *testing.T) {
		a := 0xFF
		assert.Equal(t, 255, a)
	})
}

// e 表示 10 的连乘
func TestWithExponentPart(t *testing.T) {
	assert.Equal(t, 1000, 1e3)
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

	// 复数字面量默认是 complex128 类型，因此与 complex64 类型比较时要先进行转换
	var vOfComplex64 complex64
	assert.Equal(t, complex64(0+0i), vOfComplex64)
}
