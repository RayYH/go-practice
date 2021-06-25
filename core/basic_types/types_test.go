package basic_types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestBooleanType(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool)
}

func TestIntegerType(t *testing.T) {
	t.Run("integer types", func(t *testing.T) {
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

func TestStringType(t *testing.T) {
	// see strings/strings_test for more information
	var aStr string
	assert.Equal(t, aStr, "")
}

func TestFloatType(t *testing.T) {
	var aFloat32 float32
	var aFloat64 float64
	assert.Equal(t, "0 0", fmt.Sprint(aFloat32, aFloat64))
}

func TestComparingFloatVariables(t *testing.T) {
	isEqual := func(f1, f2, p float64) bool {
		return math.Abs(f1-f2) < p
	}
	assert.True(t, isEqual(3.000001, 3.000002, 0.0001))
	assert.False(t, isEqual(3.000001, 3.000002, 0.000000001))
}

func TestComplexType(t *testing.T) {
	var defaultComplexValue complex128
	assert.Equal(t, "(0+0i)", fmt.Sprint(defaultComplexValue))
	re := 1.1
	im := 2.2
	comp := complex(re, im)
	assert.Equal(t, "(1.1+2.2i)", fmt.Sprint(comp))
	assert.Equal(t, 1.1, real(comp))
	assert.Equal(t, 2.2, imag(comp))

	// default is complex128
	var vOfComplex64 complex64
	assert.Equal(t, complex64(0+0i), vOfComplex64)
}

// The expression `T(v)` converts the value `v` to the type `T`.
func TestTypeConversions(t *testing.T) {
	var n int16 = 34
	var m int32
	m = int32(n)
	assert.NotEqual(t, reflect.TypeOf(n), reflect.TypeOf(m))
	// int32(n) and m are of the same type and same value
	assert.Equal(t, reflect.TypeOf(int32(n)), reflect.TypeOf(m))

	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	var r uint = 5
	assert.Equal(t, r, z)
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
	if expected != got {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestTypeInterface(t *testing.T) {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i
	assert.Equal(t, "int float64 complex128", fmt.Sprintf("%T %T %T", i, j, k))
}

func ExampleDisplayBytes() {
	DisplayBytes()
	// Output: 65 65
	// 65 - 946 - 1053236
	// A - β - 􁈴
	// 41 - 3B2 - 101234
	// U+0041 - U+03B2 - U+101234
}

func ExampleDisplayArrays() {
	DisplayArrays()
	// Output: '\x00'
	// '\x00'
	// 0 0
	// 0 0
	// 0 0
	// 0 0
	// 0.00
	// 1.11
	// 2.22
	// 1234
}