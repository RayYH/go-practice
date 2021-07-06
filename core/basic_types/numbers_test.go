package basic_types

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegers(t *testing.T) {
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

	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang
	t.Run("never use uintptr", func(t *testing.T) {
		var anUintPtr uintptr
		assert.Equal(t, uintptr(0), anUintPtr)
	})
}

func TestFloatPointNumbers(t *testing.T) {
	var aFloat32 float32
	var aFloat64 float64
	assert.Equal(t, "0 0", fmt.Sprint(aFloat32, aFloat64))
	assert.Equal(t, "3.1415927, 3.141592653589793", fmt.Sprintf("%v, %v", float32(math.Pi), math.Pi))
}

func TestDifferentRadixes(t *testing.T) {
	t.Run("bin", func(t *testing.T) {
		a := 0b10000001
		assert.Equal(t, 129, a)
		b := -0b00000001
		assert.Equal(t, -1, b)
		c := 0b100000000000000000000000000000000000000000000000000000001
		assert.Greater(t, c, 0)
	})

	t.Run("oct", func(t *testing.T) {
		a := 077
		assert.Equal(t, 63, a)
	})

	t.Run("hex", func(t *testing.T) {
		a := 0xFF
		assert.Equal(t, 255, a)
	})
}

func TestWithExponentPart(t *testing.T) {
	assert.Equal(t, float64(1000), 1e3)
}

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

	var vOfComplex64 complex64
	assert.Equal(t, complex64(0+0i), vOfComplex64)
}

func ExampleDisplayIntegerLiterals() {
	DisplayIntegerLiterals()
	// Output: 42 42 384 384 384 195951310 195951310 113774485586118 1701411834604692310 17014118346046930
}

func ExampleDisplayFloatPointLiterals() {
	DisplayFloatPointLiterals()
	// Output: 0.000 72.400 72.400 2.718 1.000 0.000 1000000.000 0.250 12345.000 15.000 15.000 0.250 2048.000 1.938 0.500 0.125 348.000
}

func ExampleDisplayImaginaryLiterals() {
	DisplayImaginaryLiterals()
	// Output: (0.00000+0.00000i) starts at byte position 0
	// (0.00000+123.00000i) starts at byte position 1
	// (0.00000+83.00000i) starts at byte position 2
	// (0.00000+2748.00000i) starts at byte position 3
	// (0.00000+0.00000i) starts at byte position 4
	// (0.00000+2.71828i) starts at byte position 5
	// (0.00000+1.00000i) starts at byte position 6
	// (0.00000+0.00000i) starts at byte position 7
	// (0.00000+1000000.00000i) starts at byte position 8
	// (0.00000+0.25000i) starts at byte position 9
	// (0.00000+12345.00000i) starts at byte position 10
	// (0.00000+0.25000i) starts at byte position 11
	// (100.00000+10.00000i) starts at byte position 12
}
