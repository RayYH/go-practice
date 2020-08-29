package types

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

// 布尔类型 bool
func TestBoolean(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool, "default value of boolean type is false")
}

// 整型
func TestInt(t *testing.T) {
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
}

// 字符串的默认值是一个空字符串
// 数学运算中应尽可能地使用 float64 位，因为 math 包中提供的类型都是 float64
// 浮点型
func TestFloat(t *testing.T) {
	var aFloat32 float32
	var aFloat64 float64
	assert.Equal(t, "0 0", fmt.Sprint(aFloat32, aFloat64))
}

// 可以使用 type(variable) 来进行显式转换一个变量的类型
func TestCasts(t *testing.T) {
	var n int16 = 34
	var m int32
	m = int32(n)
	assert.Equal(t, "34 34", fmt.Sprint(m, n))
	// x, y int
	var x, y = 3, 4
	// f float64
	var f = math.Sqrt(float64(x*x + y*y))
	// z uint
	var z = uint(f)
	var r uint = 5
	assert.Equal(t, r, z)
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

func TestTypeInterface(t *testing.T) {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i
	assert.Equal(t, "int float64 complex128", fmt.Sprintf("%T %T %T", i, j, k))
}

func TestComplex(t *testing.T) {
	var defaultComplexValue complex128
	assert.Equal(t, "(0+0i)", fmt.Sprint(defaultComplexValue))
	re := 1.1
	im := 2.2
	comp := complex(re, im)
	assert.Equal(t, "(1.1+2.2i)", fmt.Sprint(comp))
	assert.Equal(t, 1.1, real(comp))
	assert.Equal(t, 2.2, imag(comp))
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
	// Output: '\x00''\x00'
	// 000000000000
	// 0.001.112.22
	// 1234
}

func TestCreateAStringFromSlices(t *testing.T) {
	var givenString = "Geeks"

	// Creating and initializing a slices of byte
	mySlice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	// Creating a string from the slices
	myString1 := string(mySlice1)

	if myString1 != givenString {
		t.Errorf("%s not equal to %s", myString1, givenString)
	}

	mySlice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	myString2 := string(mySlice2)

	if myString2 != givenString {
		t.Errorf("%s not equal to %s", myString2, givenString)
	}
}

func TestGetStringLength(t *testing.T) {
	// 6 + 3 * 2 = 12
	myStr := "Hello 世界"
	assert.Equal(t, 12, len(myStr))
	assert.Equal(t, 8, utf8.RuneCountInString(myStr))
}

func TestStringUsage(t *testing.T) {
	var rawStr1 = ``
	assert.Equal(t, "", rawStr1)

	var xmlStr1 = `
<xml>
    <node></node>
</xml>
`
	assert.Equal(t, "\n<xml>\n    <node></node>\n</xml>\n", xmlStr1)

	var sql = `
SELECT *
FROM table
`
	assert.Equal(t, "\nSELECT *\nFROM table\n", sql)

	var json = `
{
    "key": "value"
}
`
	assert.Equal(t, "\n{\n    \"key\": \"value\"\n}\n", json)

	var multilineStr = `
    1234567890
    ~!@#$%^&*()_+|
    abcABC
    `
	assert.Equal(t, "\n    1234567890\n    ~!@#$%^&*()_+|\n    abcABC\n    ", multilineStr)

	str1 := "Hello"
	str2 := "World"
	// 字符串拼接使用 +
	assert.Equal(t, "Hello World", str1+" "+str2)
}

func TestStringOperations(t *testing.T) {
	s := "a-b-c"
	assert.Equal(t, "a", string(s[0]))
	// 但是不能通过 s[0] = 'X' 来修改 s 的值
	assert.Equal(t, []string{"a", "b", "c"}, strings.Split(s, "-"))
	a := []string{"a", "b", "c"}
	assert.Equal(t, "a-b-c", strings.Join(a, "-"))
	assert.Equal(t, "a b c", strings.TrimSpace(" a b c "))
}

func TestStringCast(t *testing.T) {
	s := "10"
	i, err := strconv.Atoi(s)
	assert.Equal(t, 10, i)
	assert.Nil(t, err)

	n := 99
	str := strconv.Itoa(n)
	assert.Equal(t, "99", str)

	var a interface{}
	a = "str"
	str2 := a.(string)
	assert.Equal(t, "str", str2)
}

// using +=
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	s := ""
	for i := 0; i < b.N; i++ {
		s += "s"
	}
	b.StopTimer()
}

// using Sprintf
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	s := "s"
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%v", s)
	}
	b.StopTimer()
}

// using strings.Builder
func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()
	var stringsBuilder strings.Builder
	s := "s"
	for i := 0; i < b.N; i++ {
		stringsBuilder.WriteString(s)
	}
	_ = stringsBuilder.String()
	b.StopTimer()
}

// using bytes.Buffer
func BenchmarkBytesBuffer(b *testing.B) {
	b.ResetTimer()
	s := "s"
	var bytesBuffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		bytesBuffer.WriteString(s)
	}
	_ = bytesBuffer.String()
	b.StopTimer()
}
