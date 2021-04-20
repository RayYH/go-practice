package types

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestStringDeclarations(t *testing.T) {
	var str string
	str = "Hello World"
	assert.Equal(t, "Hello World", str)
	assert.Equal(t, str[0], uint8('H'))
}

func TestSlicesCanBeConvertedToStrings(t *testing.T) {
	// string literal
	var givenString = "Geeks"
	// Creating and initializing a slices of byte
	mySlice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	// Creating a string from the slices
	myString1 := string(mySlice1)
	// rune
	mySlice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	// creating a string from the slice
	myString2 := string(mySlice2)

	assert.Equal(t, myString1, givenString)
	assert.Equal(t, myString2, givenString)
}

func TestStringsLength(t *testing.T) {
	// 6 + 3 * 2 = 12
	myStr := "Hello 世界"
	assert.Equal(t, 12, len(myStr))
	assert.Equal(t, 8, utf8.RuneCountInString(myStr))
}

func TestMultiLineStrings(t *testing.T) {
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
}

func TestStringsOperations(t *testing.T) {
	s := "a-b-c"
	// accessing elements
	assert.Equal(t, "a", string(s[0]))
	// split string
	assert.Equal(t, []string{"a", "b", "c"}, strings.Split(s, "-"))
	// join array
	assert.Equal(t, "a-b-c", strings.Join([]string{"a", "b", "c"}, "-"))
	// trim spaces
	assert.Equal(t, "a b c", strings.TrimSpace(" a b c "))
	// concatenate strings
	assert.Equal(t, "Hello World", "Hello"+" "+"World")
}

func TestStringConversions(t *testing.T) {
	// string to int
	s := "10"
	i, err := strconv.Atoi(s)
	assert.Equal(t, 10, i)
	assert.Nil(t, err)

	// int to string
	n := 99
	str := strconv.Itoa(n)
	assert.Equal(t, "99", str)

	// a can be any type
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
		s = ""
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
